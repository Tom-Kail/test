package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func CmdOut(command string) (output string, err error) {
	cmd := exec.Command("cmd.exe", "/c", "phantomjs.exe", " html.js ", "http://www.hnkcsj.com/sdnews.jsp?classid=2&parentid=5")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Start()
	if err != nil {
		return "", err
	}

	err = cmd.Wait()
	if err != nil {
		return "", err
	}

	return out.String(), err
}

func main() {
	output, err := CmdOut("./phantomjs.exe" + " html.js " + "http://www.hnkcsj.com/sdnews.jsp?classid=2&parentid=5")
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
