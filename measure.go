package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	linuxkit = "/usr/local/bin/hyperkit"
	desktop = "com.docker.hyperkit"
	bigsur = "com.apple.Virtualization.VirtualMachine"
)

func main() {
	binary := linuxkit
	out, err := exec.Command("/bin/sh", "-c", `ps -o pid,args -x | grep ` + binary + ` | grep -v grep | cut -f 1 -d " "`).CombinedOutput()
	if err != nil {
		panic(err)
	}
	pid := strings.TrimSpace(string(out))
	fmt.Printf("Monitoring %s pid %s\n", binary, pid)
	fmt.Printf("Waiting for 60s for it to settle\n")
	time.Sleep(60 * time.Second)
	total := float64(0.0)
	for i := 0; i < 60; i++ {
		x := cpu(pid)
		total += x
		fmt.Printf("%d / 60 was %v\n", i, x)
		time.Sleep(time.Second)
	}
	fmt.Printf("%f\n", total/60.0)
}

func cpu(pid string) float64 {
	out, err := exec.Command("/bin/sh", "-c", "ps -o %cpu -p "+pid+" | tail -1").CombinedOutput()
	if err != nil {
		panic(err)
	}
	percent, err := strconv.ParseFloat(strings.TrimSpace(string(out)), 64)
	if err != nil {
		panic(err)
	}
	return percent
}
