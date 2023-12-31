package loggy

import (
	"fmt"
	systemlog "log"
)

func init() {
	systemlog.Print("DEBUG Logging enabled")
}

// func Debug(v ...any)
// Debug calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Print.
func Debug(v ...any) {
	var x []interface{}
	x = append(x, "DBG:")
	x = append(x, FileLineFunc(2)+":")
	x = append(x, v)
	logLock.Lock()
	defer logLock.Unlock()
	systemlog.Print(x...)
}

func GDebug(v ...any) {
	var x []interface{}
	gid := fmt.Sprintf("[%04d]", getGID())
	x = append(x, "DBG:")
	x = append(x, gid)
	x = append(x, FileLineFunc(2)+":")
	x = append(x, v)
	logLock.Lock()
	defer logLock.Unlock()
	systemlog.Print(x...)
}

// func Debugf(format string, v ...any)
// Debugf calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, v ...any) {
	var x []any
	format = "DBG:%s:" + format
	x = append(x, FileLineFunc(2))
	x = append(x, v...)
	logLock.Lock()
	defer logLock.Unlock()
	systemlog.Printf(format, x...)
}

func GDebugf(format string, v ...any) {
	var x []any
	gid := fmt.Sprintf("[%04d]", getGID())
	format = "DBG:%s%s:" + format
	x = append(x, gid)
	x = append(x, FileLineFunc(2))
	x = append(x, v...)
	logLock.Lock()
	defer logLock.Unlock()
	systemlog.Printf(format, x...)
}

// func Debugln(v ...any)
// Debugln calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
func Debugln(v ...any) {
	var x []interface{}
	x = append(x, "DBG:")
	x = append(x, FileLineFunc(2)+":")
	x = append(x, v)
	logLock.Lock()
	defer logLock.Unlock()
	systemlog.Println(v...)
}

func GDebugln(v ...any) {
	var x []interface{}
	gid := fmt.Sprintf("[%04d]", getGID())
	x = append(x, "DBG:")
	x = append(x, gid)
	x = append(x, FileLineFunc(2)+":")
	x = append(x, v)
	logLock.Lock()
	defer logLock.Unlock()
	systemlog.Println(v...)
}
