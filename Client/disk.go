package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/disk"
)

type DiskStats struct {
	Device     string
	MountPoint string
	TotalSpace uint64
	UsedSpace  uint64
	FreeSpace  uint64
	Usage      float64
}

func getDiskStats() ([]DiskStats, error) {

	partitions, err := disk.Partitions(false)
	if err != nil {
		return []DiskStats{}, err
	}

	stats := []DiskStats{}

	for _, partitions := range partitions {
		diskUsage, err := disk.Usage(partitions.Mountpoint)
		if err != nil {
			return []DiskStats{}, err
		}
		stats = append(stats, DiskStats{
			Device:     partitions.Device,
			MountPoint: partitions.Mountpoint,
			TotalSpace: diskUsage.Total,
			UsedSpace:  diskUsage.Used,
			FreeSpace:  diskUsage.Free,
			Usage:      diskUsage.UsedPercent,
		})
	}

	return stats, nil
}

func printDisk() {
	// Get disk usage
	diskUsage, _ := disk.Usage("/")
	fmt.Println("\nDisk Information:")
	fmt.Printf("  Total Disk Space: %v GB\n", diskUsage.Total/1024/1024/1024)
	fmt.Printf("  Used Disk Space: %v GB\n", diskUsage.Used/1024/1024/1024)
	fmt.Printf("  Free Disk Space: %v GB\n", diskUsage.Free/1024/1024/1024)
	fmt.Printf("  Disk Usage: %.2f%%\n", diskUsage.UsedPercent)

	// Get disk partitions
	partitions, _ := disk.Partitions(false)
	for _, partition := range partitions {
		fmt.Printf("  Partition: %s\n", partition.Device)
		fmt.Printf("    Mountpoint: %s\n", partition.Mountpoint)
		fmt.Printf("    Fstype: %s\n", partition.Fstype)
	}
}
