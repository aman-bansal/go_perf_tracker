package sample

import (
	"fmt"
	"github.com/aman-bansal/go_perf_tracker/perf_tracker"
	"time"
)

func Fun1() {
	defer perf_tracker.CalculateFunctionPerf(time.Now(), "Fun1")
	time.Sleep(time.Duration(10) * time.Second)
	fmt.Println("in func 1")
}

func Fun2() {
	defer perf_tracker.CalculateFunctionPerf(time.Now(), "Fun2")
	time.Sleep(time.Duration(5) * time.Second)
	fmt.Println("in func 2")
}

func Fun3() {
	time.Sleep(time.Duration(15) * time.Second)
	fmt.Println("in func 3")
}
