package services

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
)

func GetPlatformInfo() string {
	kernel, err := host.KernelVersion()
	if err != nil {
		return "❌ Failed to get kernel version: " + err.Error()
	}

	platform, family, version, err := host.PlatformInformation()
	if err != nil {
		return "❌ Failed to get platform information: " + err.Error()
	}

	result := fmt.Sprintf(
		"🧠 Kernel Version: %s\n💻 Platform: %s\n🏠 Family: %s\n🔢 Version: %s",
		kernel, platform, family, version,
	)
	return result
}
