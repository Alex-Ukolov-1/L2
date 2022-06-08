package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mitchellh/go-ps"
)

func main() {

	scaner := bufio.NewScanner(os.Stdin)
	for scaner.Scan() {
		commandHandler(scaner.Text())
	}

}
func commandHandler(stringCommand string) {
	command := strings.Split(stringCommand, " ")[0]
	switch command {
	case `\quit`:
		fmt.Println("exit")
		os.Exit(0)
	case "cd":
		cd := strings.Replace(stringCommand, "cd ", "", 1)
		os.Chdir(cd)
	case "pwd":
		dir, _ := os.Getwd()
		fmt.Println(dir)
	case "echo":
		str := strings.Replace(stringCommand, "echo ", "", 1)
		fmt.Println(str)
	case "kill":
		strProc := strings.Replace(stringCommand, "kill ", "", 1)
		pid, err := strconv.Atoi(strProc)
		if err != nil {
			fmt.Println(err)
		}
		proc, err := os.FindProcess(pid)
		if err != nil {
			fmt.Println(err)
		}
		proc.Kill()
	case "ps":
		sliceProc, _ := ps.Processes()
		for _, proc := range sliceProc {
			fmt.Printf("Name p: %v Pid: %v\n", proc.Executable(), proc.Pid())
		}
	default:
		fmt.Println("command not recognized")
	}
}
