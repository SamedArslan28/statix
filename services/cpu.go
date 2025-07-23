package services

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"time"
)

func GetCpuInfo() string {
	// Get logical and physical CPU counts
	logicalCount, err := cpu.Counts(true)
	if err != nil {
		return "❌ Failed to get logical CPU count: " + err.Error()
	}

	physicalCount, err := cpu.Counts(false)
	if err != nil {
		return "❌ Failed to get physical CPU count: " + err.Error()
	}

	// Get CPU usage (overall)
	usage, err := cpu.Percent(500*time.Millisecond, false)
	if err != nil {
		return "❌ Failed to get CPU usage: " + err.Error()
	}

	// Get detailed info
	infoStats, err := cpu.Info()
	if err != nil || len(infoStats) == 0 {
		return "❌ Failed to get CPU details: " + err.Error()
	}

	info := infoStats[0]

	return fmt.Sprintf(
		`🧠 CPU Information
------------------------
Model Name : %s
Vendor ID  : %s
Cores      : %d physical / %d logical
Speed      : %.2f MHz
Usage      : %.2f%%`, info.ModelName, info.VendorID, physicalCount, logicalCount, info.Mhz, usage[0])
}
