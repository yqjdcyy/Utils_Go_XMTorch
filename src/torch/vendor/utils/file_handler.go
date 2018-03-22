package utils

import (
	"fmt"
	"os"

	"bitbucket.org/ansenwork/ilog"
)

// OpenOrCreate 打开&创建
func OpenOrCreate(path string) (file *os.File, err error) {

	file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if nil != err {
		ilog.Errorf("fail to Create|Open|Append file[%s]: %v", path, err.Error())
		return nil, err
	}

	ilog.Infof("Open|Create file[%s]", path)
	return file, nil
}

// Remove 移除文件
func Remove(path string) error {

	err := os.Remove(path)
	if nil != err {
		ilog.Errorf("fail to remove file[%v]: %s", path, err.Error())
		return err
	}

	ilog.Infof("remove file[%s]", path)
	return nil
}

// Open 打开文件
func Open(path string) (file *os.File, err error) {

	// check
	if 0 == len(path) {
		err = fmt.Errorf("Open() without path")
		ilog.Error(err.Error())
		return
	}

	// open
	file, err = os.Open(path)
	if nil != err {
		ilog.Errorf("fail to open file[%s]: %s", path, err.Error())
	}
	return
}
