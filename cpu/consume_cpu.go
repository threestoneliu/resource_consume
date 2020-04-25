package main

import (
	"flag"
	"time"

	"bitbucket.org/bertimus9/systemstat"
	"math"
	"runtime"
	"fmt"
	"strings"
	"strconv"
)

const sleep = time.Duration(10) * time.Millisecond

var (
	cpu_gradient = flag.String("cpu-gradient", "0", "cpu usage gradient")
	interval     = flag.Int("cpu-interval", 60, "cpu load period, unit seconds")

	cpu_base  systemstat.ProcCPUSample
	tar_pct   float64
	cpu_index = 0
	grads_i   []int
)

func doSomething() {
	for i := 1; i < 10000000; i++ {
		x := float64(0)
		x += math.Sqrt(0)
	}
}

func consume_cpu() {
	for {
		cpu := systemstat.GetProcCPUAverage(cpu_base, systemstat.GetProcCPUSample(), systemstat.GetUptime().Uptime)
		if cpu.TotalPct < tar_pct {
			doSomething()
		} else {
			time.Sleep(sleep)
		}
	}
}

func change() {
	if cpu_index >= len(grads_i) {
		cpu_index = 0
	}
	tar_pct = float64(grads_i[cpu_index]) / float64(10)
	cpu_base = systemstat.GetProcCPUSample()
	cpu_index += 1
}

func main() {
	flag.Parse()
	if *interval <= 0 {
		fmt.Println("cpu-period must be greater than 0")
		return
	}
	grads := strings.Split(*cpu_gradient, ",")
	crest := 0
	for _, grad := range grads {
		v, err := strconv.Atoi(grad)
		if err != nil {
			continue
		}
		grads_i = append(grads_i, v)
		if crest < v {
			crest = v
		}
	}
	if len(grads_i) == 0 {
		fmt.Println("cpu-gradient has no valid data ")
		return
	}
	id := time.Duration(*interval) * time.Second

	threads := int(math.Ceil(float64(crest) / float64(1000)))
	runtime.GOMAXPROCS(threads)
	fmt.Println("grad_i: ", grads_i)
	change()

	for i := 0; i < threads; i++ {
		go consume_cpu()
	}

	pt := time.Tick(id)
	for {
		select {
		case <-pt:
			change()
		}
	}

}
