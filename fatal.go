package loggy

import (
	systemlog "log"
)

// func Fatal(v ...any)
// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...any) {
	systemlog.Fatal(v...)
}

func FatalStack(v ...any) {
	logLock.Lock()
	defer logLock.Unlock()
	stack()
	systemlog.Fatal(v...)
}

// func Fatalf(format string, v ...any)
// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...any) {
	logLock.Lock()
	defer logLock.Unlock()
	systemlog.Fatalf(format, v...)
}

func FatalfStack(format string, v ...any) {
	logLock.Lock()
	defer logLock.Unlock()
	stack()
	systemlog.Fatalf(format, v...)
}

// func Fatalln(v ...any)
// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(v ...any) {
	logLock.Lock()
	defer logLock.Unlock()
	systemlog.Fatalln(v...)
}

func FatallnStack(v ...any) {
	logLock.Lock()
	defer logLock.Unlock()
	stack()
	systemlog.Fatalln(v...)
}

// func Debug(v ...any)
// Debug calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Print.
func FFatal(v ...any) {
	var x []interface{}
	x = append(x, "FATAL:")
	x = append(x, FileLine(2)+":")
	x = append(x, v)
	logLock.Lock()
	defer logLock.Unlock()
	systemlog.Fatal(x...)
}

// func Debugf(format string, v ...any)
// Debugf calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Printf.
func FFatalf(format string, v ...any) {
	var x []any
	format = "FATAL:%s:" + format
	x = append(x, FileLine(2))
	x = append(x, v...)
	logLock.Lock()
	defer logLock.Unlock()
	systemlog.Fatalf(format, x...)
}

// func Debugln(v ...any)
// Debugln calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
func FFatalln(v ...any) {
	var x []interface{}
	x = append(x, "FATAL:")
	x = append(x, FileLine(2)+":")
	x = append(x, v)
	logLock.Lock()
	defer logLock.Unlock()
	systemlog.Fatalln(v...)
}

// func Debug(v ...any)
// Debug calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Print.
func FFatalStack(v ...any) {
	var x []interface{}
	x = append(x, "FATAL:")
	x = append(x, FileLine(2)+":")
	x = append(x, v)
	logLock.Lock()
	defer logLock.Unlock()
	stack()
	systemlog.Fatal(x...)
}

// func Debugf(format string, v ...any)
// Debugf calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Printf.
func FFatalfStack(format string, v ...any) {
	var x []any
	format = "FATAL:%s:" + format
	x = append(x, FileLine(2))
	x = append(x, v...)
	logLock.Lock()
	defer logLock.Unlock()
	stack()
	systemlog.Fatalf(format, x...)
}

// func Debugln(v ...any)
// Debugln calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
func FFatallnStack(v ...any) {
	var x []interface{}
	x = append(x, "FATAL:")
	x = append(x, FileLine(2)+":")
	x = append(x, v)
	logLock.Lock()
	defer logLock.Unlock()
	stack()
	systemlog.Fatalln(v...)
}


// func Debug(v ...any)
// Debug calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Print.
func FFFatal(v ...any) {
	var x []interface{}
	x = append(x, "FATAL:")
	x = append(x, FileLineFunc(2)+":")
	x = append(x, v)
	logLock.Lock()
	defer logLock.Unlock()
	systemlog.Fatal(x...)
}

// func Debugf(format string, v ...any)
// Debugf calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Printf.
func FFFatalf(format string, v ...any) {
	var x []any
	format = "FATAL:%s:" + format
	x = append(x, FileLineFunc(2))
	x = append(x, v...)
	logLock.Lock()
	defer logLock.Unlock()
	systemlog.Fatalf(format, x...)
}

// func Debugln(v ...any)
// Debugln calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
func FFFatalln(v ...any) {
	var x []interface{}
	x = append(x, "FATAL:")
	x = append(x, FileLineFunc(2)+":")
	x = append(x, v)
	logLock.Lock()
	defer logLock.Unlock()
	systemlog.Fatalln(v...)
}


// func Debug(v ...any)
// Debug calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Print.
func FFFatalStack(v ...any) {
	var x []interface{}
	x = append(x, "FATAL:")
	x = append(x, FileLineFunc(2)+":")
	x = append(x, v)
	logLock.Lock()
	defer logLock.Unlock()
	stack()
	systemlog.Fatal(x...)
}

// func Debugf(format string, v ...any)
// Debugf calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Printf.
func FFFatalfStack(format string, v ...any) {
	var x []any
	format = "FATAL:%s:" + format
	x = append(x, FileLineFunc(2))
	x = append(x, v...)
	logLock.Lock()
	defer logLock.Unlock()
	stack()
	systemlog.Fatalf(format, x...)
}

// func Debugln(v ...any)
// Debugln calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
func FFFatallnStack(v ...any) {
	var x []interface{}
	x = append(x, "FATAL:")
	x = append(x, FileLineFunc(2)+":")
	x = append(x, v)
	logLock.Lock()
	defer logLock.Unlock()
	stack()
	systemlog.Fatalln(v...)
}
