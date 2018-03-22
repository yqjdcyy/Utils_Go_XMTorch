package utils

import (
	"testing"
)

func TestHandler_Append(t *testing.T) {

	var handler = new(Handler)
	var path = "D:\\download\\test.txt"

	handler.SetPath(path)
	handler.Append("hello, %s!", "torch")
	handler.Close()

	Remove(path)
}
