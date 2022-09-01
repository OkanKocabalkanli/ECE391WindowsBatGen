package main

import (
	"fmt"
	"os"
	"strings"
)

func createBat(qemuPath string, qcowPath string, name string, isKernal bool, gdb bool) {
	noCMDCommand := `if not DEFINED IS_MINIMIZED set IS_MINIMIZED=1 && start "" /min "%~dpnx0" %* && exit`
	kernelFlag := ""
	gdbFlag := ""
	if isKernal {
		kernelFlag = "-kernel"
	}
	if gdb{
		gdbFlag = "-gdb tcp:127.0.0.1:1234"
	}
	command := fmt.Sprintf(
		"%sC:\ncd %s\nqemu-system-i386w.exe -hda %s %s -m 512 %s -name %s",
		noCMDCommand,
		qemuPath,
		qcowPath,
		kernelFlag,
		gdbFlag,
		name,
	)

	fmt.Println(command)
	err := os.WriteFile(name+".bat", []byte(command), 0660)

	fmt.Println(err)
}

func main() {
	var wfhDir string
	fmt.Println("Enter Path of your ece391_share folder")
	fmt.Scanln(&wfhDir)
	wfhDir = strings.Replace(wfhDir,`\`,"\\",-1) + "\\"
	fmt.Println(wfhDir)
	var qemuPath string
	qemuPath = wfhDir + "qemu_win"
	fmt.Println(qemuPath)
	develPath := wfhDir + "ece391_share\\work\\vm\\devel.qcow"
	testPath := wfhDir + "ece391_share\\work\\vm\\test.qcow"
	createBat(qemuPath, develPath, "devel", false, false)
	createBat(qemuPath, testPath, "testdebug", true, true)
	createBat(qemuPath, testPath, "test_nodebug", true, false)

}
