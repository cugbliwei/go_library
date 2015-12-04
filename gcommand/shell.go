package gcommand

import (
	"bytes"
	"log"
	"os/exec"
)

//在go中用来执行shell命令
func Shell(s string) string {
	cmd := exec.Command("/bin/sh", "-c", s)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Println("shell Command error:", err)
		return "shell Command error"
	}
	return out.String()
}
