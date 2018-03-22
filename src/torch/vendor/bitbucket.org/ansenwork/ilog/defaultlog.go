package ilog

import (
	"fmt"
	"log"
)

var (
	prefixDebug = " [debug] "
	prefixTrace = " [trace] "
	prefixInfo  = " [info] "
	prefixWarn  = " [warn] "
	prefixError = " [error] "
	prefixPanic = " [panic] "
)

type defaultLogger struct {
}

// Debug write the debug msg
func (logger *defaultLogger) Debug(format string) {
	log.Println(prefixDebug, format)
}

// Debugf write the debug msg
func (logger *defaultLogger) Debugf(format string, args ...interface{}) {
	log.Printf("%s%s", prefixDebug, fmt.Sprintf(format, args...))
}

// Trace write the trace msg
func (logger *defaultLogger) Trace(format string) {
	log.Println(prefixTrace, format)
}

// Tracef write the trace msg
func (logger *defaultLogger) Tracef(format string, args ...interface{}) {
	log.Printf("%s%s", prefixTrace, fmt.Sprintf(format, args...))
}

// Info write the info msg
func (logger *defaultLogger) Info(format string) {
	log.Println(prefixInfo, format)
}

// Infof write the info msg
func (logger *defaultLogger) Infof(format string, args ...interface{}) {
	log.Printf("%s%s", prefixInfo, fmt.Sprintf(format, args...))
}

// Warn write the warn msg
func (logger *defaultLogger) Warn(format string) {
	log.Println(prefixWarn, format)
}

// Warnf write the warn msg
func (logger *defaultLogger) Warnf(format string, args ...interface{}) {
	log.Printf("%s%s", prefixWarn, fmt.Sprintf(format, args...))
}

// Error write the error msg
func (logger *defaultLogger) Error(format string) {
	log.Println(prefixError, format)
}

// Errorf write the error msg
func (logger *defaultLogger) Errorf(format string, args ...interface{}) {
	log.Printf("%s%s", prefixError, fmt.Sprintf(format, args...))
}

// Panic write the panic msg
func (logger *defaultLogger) Panic(format string) {
	log.Fatal(prefixPanic, format)
}

// Panicf write the panic msg
func (logger *defaultLogger) Panicf(format string, args ...interface{}) {
	log.Fatalf("%s%s", prefixPanic, fmt.Sprintf(format, args...))
}
