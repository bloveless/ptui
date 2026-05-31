package main

import (
	"fmt"
	"io"
)

type Logger struct {
	out io.Writer
}

func NewLogger(out io.Writer) Logger {
	return Logger{
		out: out,
	}
}

func (l Logger) Logf(msg string, args ...any) {
	_, err := fmt.Fprintf(l.out, msg+"\n", args...)
	if err != nil {
		fmt.Printf("unable to write to log file: %v", err)
	}
}
