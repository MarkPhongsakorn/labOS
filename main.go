package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	cpu1  string
	cpu2  string
	ready []string
	io1   []string
	io2   []string
	io3   []string
	io4   []string
)

func initialized() {
	cpu1 = ""
	cpu2 = ""
	ready = make([]string, 10)
	io1 = make([]string, 10)
	io2 = make([]string, 10)
	io3 = make([]string, 10)
	io4 = make([]string, 10)
}

func showProcess() {
	fmt.Printf("----------------------\n")
	fmt.Printf("CPU1  -> %s\n", cpu1)
	fmt.Printf("CPU2  -> %s\n", cpu2)
	fmt.Printf("READY -> ")
	for i := range ready {
		fmt.Printf("%s ", ready[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O 1 >")
	for i := range io1 {
		fmt.Printf("%s ", io1[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O 2 >")
	for i := range io2 {
		fmt.Printf("%s ", io3[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O 3 >")
	for i := range io3 {
		fmt.Printf("%s ", io3[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O 4 >")
	for i := range io4 {
		fmt.Printf("%s ", io4[i])
	}
	fmt.Printf("\n")
	fmt.Printf("\n\nCommand >")
}

func getCommand() string {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.Trim(data, "\n")
	return data
}

func insertQ(q []string, d string) {
	for i := range q {
		if q[i] == "" {
			q[i] = d
			break
		}
	}
}

func command_new(p string) {
	if cpu1 == "" {
		cpu1 = p
	} else if cpu2 == "" {
		cpu2 = p
	} else {
		insertQ(ready, p)
	}
}

func command_t(cpuName string) {
	if cpuName == "cpu1" {
		cpu1 = deleteQ(ready)
	} else if cpuName == "cpu2" {
		cpu2 = deleteQ(ready)
	}
}

func deleteQ(q []string) string {
	result := q[0]
	for i := range q {
		if i == 0 {
			continue
		}
		q[i-1] = q[i]
	}
	q[9] = ""
	return result
}

func command_ex(cpuName string) {
	p := deleteQ(ready)
	if p == "" {
		return
	}
	if cpuName == "cpu1" {
		insertQ(ready, cpu1)
		cpu1 = p
	} else if cpuName == "cpu2" {
		insertQ(ready, cpu2)
		cpu2 = p
	}
}

func command_io(ioN string, cpuN string) {
	switch ioN {
	case "1":
		io_cpu(io1, cpuN)
	case "2":
		io_cpu(io2, cpuN)
	case "3":
		io_cpu(io3, cpuN)
	case "4":
		io_cpu(io4, cpuN)
	default:
		return
	}
}

func io_cpu(io []string, check_cpu string) {
	if check_cpu == "cpu1" {
		insertQ(io, cpu1)
		cpu1 = ""
	} else if check_cpu == "cpu2" {
		insertQ(io, cpu2)
		cpu2 = ""
	}
	command_ex(check_cpu)
}

func command_iox(ioN string) {
	x := ""
	switch ioN {
	case "1":
		x = deleteQ(io1)
	case "2":
		x = deleteQ(io2)
	case "3":
		x = deleteQ(io3)
	case "4":
		x = deleteQ(io4)
	default:
		return
	}
	if x == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = x
	} else if cpu2 == "" {
		cpu2 = x
	} else {
		insertQ(ready, x)
	}
}

func main() {
	initialized()
	for {
		showProcess()
		command := getCommand()
		commandx := strings.Split(command, " ")
		switch commandx[0] {
		case "exit":
			return
		case "new":
			for i := range commandx {
				if i == 0 {
					continue
				}
				command_new(commandx[i])
			}
		case "terminate":
			for i := range commandx {
				if i == 0 {
					continue
				}
				command_t(commandx[i])
			}
		case "expire":
			for i := range commandx {
				if i == 0 {
					continue
				}
				command_ex(commandx[i])
			}
		case "io":
			command_io(commandx[1], commandx[2])
		case "iox":
			command_iox(commandx[1])
		default:
			fmt.Printf("\nCommand Error!!\n\n")
		}
	}
}
