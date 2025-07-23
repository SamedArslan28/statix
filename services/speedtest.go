package services

import (
	"fmt"
	"github.com/showwin/speedtest-go/speedtest"
)

type Speed struct {
	Download float64
	Upload   float64
	Ping     float64
	Server   string
	Distance float64
}

func RunSpeedTest() string {
	_, err := speedtest.FetchUserInfo()
	if err != nil {
		return "❌ Failed to fetch user info: " + err.Error()
	}

	serverList, err := speedtest.FetchServers()
	if err != nil {
		return "❌ Failed to fetch server list: " + err.Error()
	}

	servers, err := serverList.FindServer([]int{})
	if err != nil || len(servers) == 0 {
		return "❌ Failed to find test server"
	}

	server := servers[0]

	if err := server.DownloadTest(); err != nil {
		return "❌ Download test failed: " + err.Error()
	}
	if err := server.UploadTest(); err != nil {
		return "❌ Upload test failed: " + err.Error()
	}

	return fmt.Sprintf(
		"📡 Server: %s (%.2f km)\n📥 Download: %.2f Mbps\n📤 Upload: %.2f Mbps\n⏱️ Ping: %.2f ms",
		server.Host,
		server.Distance,
		server.DLSpeed*8/1_000_000,
		server.ULSpeed*8/1_000_000,
		server.Latency.Seconds()*1000,
	)
}
