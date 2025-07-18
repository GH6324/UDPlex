# UDPlex 认证协议

## 协议头

| 字段名   | 长度   | 描述                     |
|----------|--------|--------------------------|
| Version  | 1字节  | 协议版本，当前为1        |
| MsgType  | 1字节  | 消息类型                 |
| Reserved | 2字节  | 保留字段                 |
| Length   | 4字节  | 数据部分长度（大端序）   |
| Data     | Length | 数据部分，内容依MsgType  |

---

## 消息类型

| 类型值 | 消息类型      | 描述         |
|--------|---------------|--------------|
| 1      | AuthChallenge | 认证挑战     |
| 2      | AuthResponse  | 认证响应     |
| 4      | Heartbeat     | 心跳包       |
| 5      | Data          | 数据包       |
| 6      | Disconnect    | 断开连接     |

---

## 认证流程

1. **客户端发送 AuthChallenge 消息**，包含一个随机数 challenge 与时间戳，使用 HMAC-SHA256 加密。
    - `challenge` (32字节)：随机数
    - `timestamp` (8字节)：时间戳，单位为毫秒
    - `forwardID` (8字节)：前向连接 ID，标识连接
    - `poolID` (8字节)：连接池 ID，标识连接池
    - `mac` (32字节)：HMAC-SHA256 加密结果
2. **服务器收到 AuthChallenge 后**，认证成功则返回响应，认证失败则直接丢弃数据，无任何响应，避免被探测。
3. **服务器认证时需判断：**
    - `timestamp` 是否在合理范围内（如5分钟内）
    - `mac` 是否正确
4. **服务器认证成功后**，生成一个随机数 response，并使用 HMAC-SHA256 加密 challenge 和 response。
    - `response` (32字节)：服务器生成的随机数
    - `timestamp` (8字节)：时间戳，单位为毫秒
    - `poolID` (8字节)：连接池 ID，标识连接池
    - `mac` (32字节)：HMAC-SHA256 加密结果
5. **客户端收到 AuthResponse 后**，验证 mac 是否正确，若正确则认证成功。

---

## 心跳包

心跳包用于保持连接活跃，客户端和服务器可定期发送心跳包以确认连接状态。

1. 客户端发送 Heartbeat 消息，不包含任何数据。
2. 服务器收到 Heartbeat 后，直接返回相同的 Heartbeat 消息。
3. 若客户端长时间未发送心跳包，服务器可主动断开连接。
4. 客户端连续发送心跳包超过一定次数（如3次）未收到服务器响应，超时 n 秒（如5s）后开始发送 AuthChallenge，直到收到服务器响应或超过最大重试次数（如5次）后开始重连。

---

## 数据包

数据包用于传输实际数据内容，客户端和服务器可互相发送数据包。

- **未加密的数据包格式：**
    - `connID` (8字节)：随机数，用于标识连接
    - `data` (Length字节)：实际数据内容

- **加密的数据包格式（AES-128-GCM）：**
    - `Nonce` (12字节)：AES-128-GCM 初始化向量
    - `Timestamp` (8字节)：时间戳，单位为毫秒（与 data 一起加密，用于防止重放攻击）
    - `connID` (8字节)：随机数，用于标识连接
    - `data` (Length-28字节)：实际数据内容

2. 收到数据后需判断 timestamp 是否在合理范围内（如30s内），超时则丢弃数据包。