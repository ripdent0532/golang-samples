package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	//env := os.Environ()
	//procAttr :=  &os.ProcAttr{
	//	Env: env,
	//	Files: []*os.File{
	//		os.Stdin,
	//		os.Stdout,
	//		os.Stderr,
	//	},
	//}
	//pid, err := os.StartProcess("/bin/ls", []string{"-l"}, procAttr)
	//if err != nil {
	//	fmt.Printf("Error %v starting process!", err)
	//	os.Exit(1)
	//}
	//fmt.Printf("The process id is %v\n", pid)
	cmd := exec.Command("vim")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error %v executing command!", err)
		os.Exit(1)
	}
	fmt.Printf("The command is %v", cmd)

}