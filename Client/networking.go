package main

import (
	"time"

	"github.com/shirou/gopsutil/v3/net"
)

type NetStats struct {
	InterfaceName string
	IPAddresses   []string
	MACAddress    string
	UploadSpeed   float64
	DownloadSpeed float64
}

// GetNetworkStats fetches network details and speeds
func GetNetworkStats() ([]NetStats, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	time.Sleep(1 * time.Second) // Pause to measure speeds
	finalStats, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}

	stats := []NetStats{}
	for _, iface := range interfaces {
		ips := []string{}
		for _, addr := range iface.Addrs {
			ips = append(ips, addr.Addr)
		}

		for _, counter := range finalStats {
			if counter.Name == iface.Name {
				stats = append(stats, NetStats{
					InterfaceName: iface.Name,
					IPAddresses:   ips,
					MACAddress:    iface.HardwareAddr,
					UploadSpeed:   float64(counter.BytesSent*8) / (1 << 30),
					DownloadSpeed: float64(counter.BytesRecv*8) / (1 << 30),
				})
			}
		}
	}

	return stats, nil
}
