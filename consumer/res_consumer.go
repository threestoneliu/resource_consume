package main

import (
	"flag"
	"runtime"
	"os/exec"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var (
	cpu_gradient = flag.String("cpu-gradient", "0", "cpu usage gradient")
	interval     = flag.Int("cpu-interval", 60, "cpu load period, unit seconds")
	memory       = flag.Int("memory", 0, "memory number. unit mb")
	cpuexec      = flag.String("cpuexec", "gdcpu", "cpu binary path")
	mem          [][]byte
)

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(1)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	for i := 1; i < *memory; i++ {
		buf := make([]byte, 1<<20)
		for i := range buf {
			buf[i] = 0
		}
		mem = append(mem, buf)
	}

	cmd := exec.Command(*cpuexec, fmt.Sprintf("-cpu-gradient=%s", *cpu_gradient), fmt.Sprintf("-cpu-interval=%d", *interval))
	if err := cmd.Start(); err != nil {
		fmt.Printf("error,%v", err)
		return
	}
	<-c
	cmd.Process.Kill()
	cmd.Wait()
}
