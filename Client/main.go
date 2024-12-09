package main

import (
	"Watchtower-Client/processes"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	for {
		UploadData()
	}
}

func UploadData() {
	// Change the IP to the server
	apiEndpoint := "http://127.0.0.1:27018"

	// CPU Stats
	cpuStats, err := GetCPUStats()
	if err != nil {
		fmt.Println("Error fetching CPU stats:", err)
		return
	}

	// Memory Stats
	memStats, err := GetMemoryStats()
	if err != nil {
		fmt.Println("Error fetching Memory stats:", err)
		return
	}

	// Network Stats
	netStats, err := GetNetworkStats()
	if err != nil {
		fmt.Println("Error fetching Network stats:", err)
		return
	}

	systemStats, err := GetSystemStats()
	if err != nil {
		fmt.Println("Error fetching System stats:", err)
		return
	}

	diskStats, err := getDiskStats()
	if err != nil {
		fmt.Println("Error fetching Disk stats:", err)
		return
	}

	processCount := 5
	topProcesses, err := processes.GetTopProcesses(processCount)
	if err != nil {
		fmt.Printf("Error fetching top processes: %v\n", err)
		return
	}

	// Prepare payload
	data := map[string]interface{}{
		"system":    systemStats,
		"cpu":       cpuStats,
		"memory":    memStats,
		"network":   netStats,
		"processes": topProcesses,
		"disk":      diskStats,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Send to API
	resp, err := http.Post(apiEndpoint, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error sending data to API:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Data sent successfully. Status Code:", resp.StatusCode)
}
