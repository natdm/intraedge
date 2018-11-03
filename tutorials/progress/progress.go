package progress

import "time"

// ProgressWriter is a writer that reports progress being written to, and stores it in a
// byte array.
type ProgressWriter struct {
	done, total uint64
	bs          []byte
	pCh         chan int
}

// New takes the length of the file being written and returns a ProgressWriter
// and a channel to read the current percentage from
func New(len uint64) (pw *ProgressWriter, percent <-chan int) {
	pw = &ProgressWriter{
		done:  0,
		total: len,
		bs:    make([]byte, len),
		pCh:   make(chan int),
	}

	return pw, pw.pCh
}

func (pw *ProgressWriter) eof() bool {
	isDone := pw.done == pw.total
	if isDone {
		close(pw.pCh)
	}
	return isDone
}

func (pw *ProgressWriter) Write(bs []byte) (int, error) {
	for i, b := range bs {
		if pw.eof() {
			return i, nil
		}
		pw.bs[pw.done] = b
		pw.done++
		percent := int((float64(pw.done) / float64(pw.total)) * 100)
		select {
		case pw.pCh <- percent:
		case <-time.After(time.Millisecond):
		}
	}

	if pw.eof() {
		return len(bs), nil
	}
	return len(bs), nil
}

func (pw *ProgressWriter) Bytes() []byte {
	return pw.bs
}
