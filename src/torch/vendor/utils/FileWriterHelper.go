package utils

import (
	"fmt"
	"os"

	"bitbucket.org/ansenwork/ilog"
)

// Handler 文件操作符
type Handler struct {
	file *os.File
}

// SetPath 设置操作文件
func (handler *Handler) SetPath(path string) {

	file, err := OpenOrCreate(path)
	if nil != err {
		ilog.Panicf("fail to open or crate file[%v]: %v", path, err.Error())
	}
	(*handler).file = file
	ilog.Infof("handler bind to file[%v]", path)
}

// Close 关闭文件
func (handler *Handler) Close() {

	if nil != (*handler).file {
		err := (*handler).file.Close()
		if nil != err {
			ilog.Panicf("fail to close file: %v", err.Error())
		}
	}
}

// Append 追加文件
func (handler *Handler) Append(format string, a ...interface{}){
	
	if nil== (*handler).file{
		ilog.Errorf("fail to write file: init file first")
		return
	}

	if _, err:=(*handler).file.WriteString(fmt.Sprintf(format, a)); nil!= err{
		ilog.Errorf("fail to write file: %v", err.Error())
	}
}