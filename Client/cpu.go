package main

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

type CPUStats struct {
	Cores         int
	UsagePerCore  []float64
	AverageUsage  float64
	ClockSpeedMhz []float64
}

// GetCPUStats fetches CPU-related information
func GetCPUStats() (CPUStats, error) {
	cores, err := cpu.Counts(true)
	if err != nil {
		return CPUStats{}, err
	}

	usagePerCore, err := cpu.Percent(500*time.Millisecond, true)
	if err != nil {
		return CPUStats{}, err
	}

	averageUsage, err := cpu.Percent(500*time.Millisecond, false)
	if err != nil {
		return CPUStats{}, err
	}

	info, err := cpu.Info()
	if err != nil {
		return CPUStats{}, err
	}

	clockSpeeds := []float64{}
	for _, core := range info {
		clockSpeeds = append(clockSpeeds, core.Mhz)
	}

	return CPUStats{
		Cores:         cores,
		UsagePerCore:  usagePerCore,
		AverageUsage:  averageUsage[0],
		ClockSpeedMhz: clockSpeeds,
	}, nil
}
