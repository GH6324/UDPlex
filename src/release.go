//go:build !dev
// +build !dev

package main

func (r *Router) incrementBufferRef() {
	// No-op in production
}

func (r *Router) decrementBufferRef() {
	// No-op in production
}

func (r *Router) logBufferRef() {
	// No-op in production
}

// initPprof Do not start pprof in production environment
func initPprof() {
	// No-op in production
}
