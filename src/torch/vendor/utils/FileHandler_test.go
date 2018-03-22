package utils

import (
	"testing"
)

func TestOperate(t *testing.T) {

	var path= "D:\\download\\test2.txt"

	file, err:= OpenOrCreate(path)
	if nil!= err{
		t.Errorf(err.Error())
		return
	}
	file.Close()
	err= Remove(path)
	if nil!= err{
		t.Errorf(err.Error())
	}
}