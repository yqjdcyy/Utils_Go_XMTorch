package utils

import (
	"bufio"
	"fmt"
	"testing"
)

func TestOperate(t *testing.T) {

	var path = "D:\\download\\test2.txt"

	// create
	file, err := OpenOrCreate(path)
	if nil != err {
		t.Errorf(err.Error())
		return
	}

	// write
	file.WriteString(fmt.Sprintf("1\n2\n3\n4"))
	file.Close()

	// oepn
	file, err = Open(path)
	if nil != err {
		t.Errorf(err.Error())
		return
	}

	// read
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t.Log(scanner.Text())
	}
	file.Close()

	// remove
	err = Remove(path)
	if nil != err {
		t.Errorf(err.Error())
	}
}
