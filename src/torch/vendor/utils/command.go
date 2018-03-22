package utils

import (
	"fmt"
	"os/exec"
)

// OpenURI 浏览器打开对应链接
func OpenURI(uri string) error {

	fmt.Printf("open URI[%s]", uri)
	return exec.Command("cmd", "/C", "start", uri).Start()
}
