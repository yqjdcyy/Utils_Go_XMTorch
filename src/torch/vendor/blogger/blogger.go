package blogger

import (
	blog "github.com/YoungPioneers/blog4go"
	imlog "bitbucket.org/ansenwork/ilog"
	"os"
)

// Load 加载配置文件
func Load(path string) (logger imlog.Log, err error) {
	err = blog.NewWriterFromConfigAsFile(path)
	if err != nil {
		return
	}
	logger = new(blogger)
	return
}

// Close 关闭
func Close() {
	blog.Close()
	os.Exit(1)
}

type blogger struct{}

// Debug write the debug msg
func (logger *blogger) Debug(format string) {
	blog.Debug(format)
}

// Debugf write the debug msg
func (logger *blogger) Debugf(format string, args ...interface{}) {
	blog.Debugf(format, args...)
}

// Trace write the trace msg
func (logger *blogger) Trace(format string) {
	blog.Trace(format)
}

// Tracef write the trace msg
func (logger *blogger) Tracef(format string, args ...interface{}) {
	blog.Tracef(format, args...)
}

// Info write the info msg
func (logger *blogger) Info(format string) {
	blog.Info(format)
}

// Infof write the info msg
func (logger *blogger) Infof(format string, args ...interface{}) {
	blog.Infof(format, args...)
}

// Warn write the warn msg
func (logger *blogger) Warn(format string) {
	blog.Warn(format)
}

// Warnf write the warn msg
func (logger *blogger) Warnf(format string, args ...interface{}) {
	blog.Warnf(format, args...)
}

// Error write the error msg
func (logger *blogger) Error(format string) {
	blog.Error(format)
}

// Errorf write the error msg
func (logger *blogger) Errorf(format string, args ...interface{}) {
	blog.Errorf(format, args...)
}

// Panic write the panic msg
func (logger *blogger) Panic(format string) {
	blog.Critical(format)
}

// Panicf write the panic msg
func (logger *blogger) Panicf(format string, args ...interface{}) {
	blog.Criticalf(format, args...)
}
