package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/host"
)

type SystemStats struct {
	Hostname        string
	OS              string
	Platform        string
	PlatformVersion string
	KernelVersion   string
	Architecture    string
	Uptime          string
	BootTime        string
	LoggedInUsers   []host.UserStat `json:",omitempty"` // Omit from JSON if empty
}

// GetSystemStats fetches system and user information
func GetSystemStats() (SystemStats, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return SystemStats{}, err
	}

	info, err := host.Info()
	if err != nil {
		return SystemStats{}, err
	}

	uptimeDuration := time.Duration(info.Uptime) * time.Second
	bootTime := time.Unix(int64(info.BootTime), 0).Format("2006-01-02 15:04:05")

	// Platform-specific handling for `host.Users()`
	var loggedInUsers []host.UserStat
	if runtime.GOOS != "windows" { // Skip on Windows
		users, err := host.Users()
		if err != nil {
			fmt.Printf("Error fetching logged-in users: %v\n", err)
		} else {
			loggedInUsers = users
		}
	} else {
		fmt.Println("Skipping logged-in users on Windows (not supported).")
	}

	return SystemStats{
		Hostname:        hostname,
		OS:              runtime.GOOS,
		Platform:        info.Platform,
		PlatformVersion: info.PlatformVersion,
		KernelVersion:   info.KernelVersion,
		Architecture:    runtime.GOARCH,
		Uptime:          uptimeDuration.String(),
		BootTime:        bootTime,
		LoggedInUsers:   loggedInUsers, // Empty for Windows
	}, nil
}
