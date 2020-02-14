package log

import (
	"fmt"
	"log"
)

var DebugMode = false

func Printf(format string, a ...interface{}) {
	if DebugMode {
		fmt.Printf(format, a...)
	}
}

func Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}
func Fatalln(v ...interface{}) {
	if v == nil || len(v) == 0 || (len(v) == 1 && v[0] == nil) {
		return
	}
	log.Fatalln(v...)
}
