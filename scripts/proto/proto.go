// +build ignore
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("protoc", "-I", "nc-go", "nc-go/ncgo.proto", "--go_out=nc-go", "--go_out=nc-go")
	cmd.Dir = "./../.."
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatalf(fmt.Sprint(err) + ": " + stderr.String())
	}
}
