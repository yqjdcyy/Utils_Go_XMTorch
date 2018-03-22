package utils

import (
	"fmt"
	"testing"
)

func TestOpenURI(t *testing.T) {

	uri := "D:/work/git/yao/go/Utils_Go_XMTorch/src/torch/template/polarArea.html"

	if err := OpenURI(uri); nil != err {
		fmt.Errorf("fail to open uri[%s]: %s", uri, err.Error())
		t.Errorf("fail to open uri[%s]: %s", uri, err.Error())
	}
	fmt.Printf("open uri[%s]", uri)
}
