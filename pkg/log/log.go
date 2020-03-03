package log

import (
	"fmt"
	"io"
)

type Log interface {
	WriteLog(error)
	WriteFLog(string, error)
}

type StdOutLog struct {
	Out io.Writer
}

func (lg *StdOutLog) WriteLog(err error) {
	fmt.Fprintf(lg.Out, "%w", err)
}
func (lg *StdOutLog) WriteFLog(format string, err error) {
	fmt.Fprintf(lg.Out, format, err)
}
