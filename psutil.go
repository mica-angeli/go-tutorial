package main

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
)

func main() {
	v, _ := mem.VirtualMemory()
	fmt.Println(v)

	procs, _ := process.Processes()
	for _, p := range procs {
		p.CPUPercent()
		// fmt.Println(percent)
	}
	// fmt.Println(procs)
}
