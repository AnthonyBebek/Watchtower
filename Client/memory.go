package main

import (
	"github.com/shirou/gopsutil/v3/mem"
)

type MemoryStats struct {
	Total           uint64
	Available       uint64
	Used            uint64
	UsedPercent     float64
	SwapTotal       uint64
	SwapUsed        uint64
	SwapUsedPercent float64
}

// GetMemoryStats fetches memory-related information
func GetMemoryStats() (MemoryStats, error) {
	vmem, err := mem.VirtualMemory()
	if err != nil {
		return MemoryStats{}, err
	}

	swap, err := mem.SwapMemory()
	if err != nil {
		return MemoryStats{}, err
	}

	return MemoryStats{
		Total:           vmem.Total,
		Available:       vmem.Available,
		Used:            vmem.Used,
		UsedPercent:     vmem.UsedPercent,
		SwapTotal:       swap.Total,
		SwapUsed:        swap.Used,
		SwapUsedPercent: swap.UsedPercent,
	}, nil
}
