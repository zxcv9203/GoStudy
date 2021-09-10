package io

// Writer is the interface that wraps the basic Write method
// Source: https://cs.opensource.google/go/go/+/refs/tags/go1.17:src/io/io.go;bpv=1;bpt=1
type Writer interface {
	// Write writes len(p) bytes from p to the underlying data stream
	// It returns the number of bytes written from p (0 <= n <= len(p))
	// and any error encountered that caused the write to stop early.
	// Write must return a non-nil error if it returns n < len(p)
	// Write must not modify the slice data, even temporarily.

	// Implementations must not retain p
	Write(p []byte) (int, err error)
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}