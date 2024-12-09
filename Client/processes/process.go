package processes

import (
	"fmt"
	"sort"

	"github.com/shirou/gopsutil/v3/process"
)

// ProcessInfo holds information about a single process
type ProcessInfo struct {
	Name         string
	PID          int32
	CPUPercent   float64
	MemoryPercent float32
}

// GetTopProcesses returns the top N processes by CPU and memory usage
func GetTopProcesses(topN int) ([]ProcessInfo, error) {
	// Get a list of all processes
	processes, err := process.Processes()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch processes: %v", err)
	}

	var processStats []ProcessInfo

	// Iterate over each process and fetch CPU and memory usage
	for _, proc := range processes {
		cpuPercent, err := proc.CPUPercent()
		if err != nil {
			continue // Skip processes where CPU usage cannot be retrieved
		}

		memoryPercent, err := proc.MemoryPercent()
		if err != nil {
			continue // Skip processes where memory usage cannot be retrieved
		}

		name, err := proc.Name()
		if err != nil {
			name = "Unknown"
		}

		// Add process stats to the list
		processStats = append(processStats, ProcessInfo{
			Name:         name,
			PID:          proc.Pid,
			CPUPercent:   cpuPercent,
			MemoryPercent: memoryPercent,
		})
	}

	// Sort processes by CPU usage in descending order
	sort.Slice(processStats, func(i, j int) bool {
		return processStats[i].CPUPercent > processStats[j].CPUPercent
	})

	// Return the top N processes
	if len(processStats) > topN {
		return processStats[:topN], nil
	}
	return processStats, nil
}
