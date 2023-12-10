package utils

import (
	"fmt"
	"os/exec"
)

func OpenLinkMac(link string) {
	runMacCmd("open", "-a", "Arc", link)
}

func runMacCmd(cmd string, args ...string) {
	err := exec.Command(cmd, args...).Start()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
