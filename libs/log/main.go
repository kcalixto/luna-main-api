package logger

import (
	"fmt"
	"runtime"
)

func getcaller() string {
	pc, _, _, ok := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		return details.Name()
	}
	return ""
}

func Info(e string, items ...any) {
	fmt.Printf("["+getcaller()+" INFO]"+e, items...)
}

func Error(e string, items ...any) {
	fmt.Printf("["+getcaller()+" ERROR]"+e, items...)
}
