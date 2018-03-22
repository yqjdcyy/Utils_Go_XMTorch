package ilog

var logger = Log(new(defaultLogger))

const (
	// DEBUG 1
	DEBUG = iota
	// TRACE 2
	TRACE
	// INFO 3
	INFO
	// WARN 4
	WARN
	// ERROR 5
	ERROR
	// PANIC 6
	PANIC
)

// Log the interface use in im
// 5 level: debug, trace, info, warn, error, panic
type Log interface {
	Debug(format string)
	Debugf(format string, args ...interface{})
	Trace(format string)
	Tracef(format string, args ...interface{})
	Info(format string)
	Infof(format string, args ...interface{})
	Warn(format string)
	Warnf(format string, args ...interface{})
	Error(format string)
	Errorf(format string, args ...interface{})
	Panic(format string)
	Panicf(format string, args ...interface{})
}

// SetLogger set the Log impl
// l the Log impl. must not be nil
func SetLogger(l Log) {
	if l != nil {
		logger = l
		return
	}
}

var loggerLevel = DEBUG

// SetLevel 设置水平
func SetLevel(level int) {
	if level < DEBUG || level > PANIC {
		logger.Errorf("level out of index: %d", level)
		return
	}
	loggerLevel = level
}

func isInLoggerLevel(level int) bool {
	if level < loggerLevel {
		return false
	}
	return true
}

// Debug write the debug msg
func Debug(format string) {
	if !isInLoggerLevel(DEBUG) {
		return
	}
	logger.Debug(format)
}

// Debugf write the debug msg
func Debugf(format string, args ...interface{}) {
	if !isInLoggerLevel(DEBUG) {
		return
	}
	logger.Debugf(format, args...)
}

// Trace write the trace msg
func Trace(format string) {
	if !isInLoggerLevel(TRACE) {
		return
	}
	logger.Trace(format)
}

// Tracef write the trace msg
func Tracef(format string, args ...interface{}) {
	if !isInLoggerLevel(TRACE) {
		return
	}
	logger.Tracef(format, args...)
}

// Info write the info msg
func Info(format string) {
	if !isInLoggerLevel(INFO) {
		return
	}
	logger.Info(format)
}

// Infof write the info msg
func Infof(format string, args ...interface{}) {
	if !isInLoggerLevel(INFO) {
		return
	}
	logger.Infof(format, args...)
}

// Warn write the warn msg
func Warn(format string) {
	if !isInLoggerLevel(WARN) {
		return
	}
	logger.Warn(format)
}

// Warnf write the warn msg
func Warnf(format string, args ...interface{}) {
	if !isInLoggerLevel(WARN) {
		return
	}
	logger.Warnf(format, args...)
}

// Error write the error msg
func Error(format string) {
	if !isInLoggerLevel(ERROR) {
		return
	}
	logger.Error(format)
}

// Errorf write the error msg
func Errorf(format string, args ...interface{}) {
	if !isInLoggerLevel(ERROR) {
		return
	}
	logger.Errorf(format, args...)
}

// Panic write the panic msg
func Panic(format string) {
	if !isInLoggerLevel(PANIC) {
		return
	}
	logger.Panic(format)
}

// Panicf write the panic msg
func Panicf(format string, args ...interface{}) {
	if !isInLoggerLevel(PANIC) {
		return
	}
	logger.Panicf(format, args...)
}
