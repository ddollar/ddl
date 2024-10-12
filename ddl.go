package ddl

import "io"

func If[T any](condition bool, thenValue, elseValue T) T {
	if condition {
		return thenValue
	}

	return elseValue
}

type multiWriteCloser struct {
	io.Writer
	cs []io.Closer
}

func MultiWriteCloser(ws ...io.Writer) io.WriteCloser {
	m := &multiWriteCloser{Writer: io.MultiWriter(ws...)}

	for _, w := range ws {
		if c, ok := w.(io.Closer); ok {
			m.cs = append(m.cs, c)
		}
	}

	return m
}

func (m *multiWriteCloser) Close() error {
	var first error

	for _, c := range m.cs {
		if err := c.Close(); err != nil && first == nil {
			first = err
		}
	}

	return first
}
