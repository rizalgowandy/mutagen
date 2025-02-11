package stream

import (
	"io"
)

// multiCloser is the io.Closer implementation underlying MultiCloser.
type multiCloser struct {
	// closers are the underlying closers.
	closers []io.Closer
}

// Close implements io.Closer.Close.
func (c *multiCloser) Close() error {
	var firstErr error
	for _, closer := range c.closers {
		if err := closer.Close(); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}

// MultiCloser creates a single closer that closes multiple underlying closers.
// The closers are closed in the order specified, and thus higher layers should
// be specified before lower. All closers will be closed, but only the first
// error encountered will be returned.
func MultiCloser(closers ...io.Closer) io.Closer {
	return &multiCloser{closers}
}
