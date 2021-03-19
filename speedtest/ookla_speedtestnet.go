package speedtest

import (
	"errors"
	"github.com/showwin/speedtest-go/speedtest"
	"time"
)

func newOoklaSpeedtestNet() (*ooklaSpeeedtestNet, error) {
	user, err := speedtest.FetchUserInfo()
	if err != nil {
		return nil, err
	}

	serverList, err := speedtest.FetchServerList(user)
	if err != nil {
		return nil, err
	}

	targets, err := serverList.FindServer(nil)
	if err != nil {
		return nil, err
	}

	return &ooklaSpeeedtestNet{
		servers: targets,
	}, nil
}

// ooklaSpeeedtestNet implements SpeedTester interface by performing speed testing via speedtest.net
type ooklaSpeeedtestNet struct {
	servers speedtest.Servers
}

// Test performs speed testing via speedtest.net and returns upload, download speeds
func (o *ooklaSpeeedtestNet) Test() (upload, download float64, err error) {
	upload, err = o.testSpeed(o.upload)
	if err != nil {
		return 0, 0, err
	}

	download, err = o.testSpeed(o.download)
	if err != nil {
		return 0, 0, err
	}

	return upload, download, nil
}

func (o *ooklaSpeeedtestNet) testSpeed(networkAction func(server *speedtest.Server) chan networkActionResult) (float64, error) {
	timeout := time.After(testTimeout)
	totalExecuted := 0
	var totalSpeed float64
	currentServer := 0

loop:
	for {
		select {
		case <-timeout:
			if totalExecuted == 0 {
				return 0, errors.New("timeout")
			}

			break loop
		case res := <-networkAction(o.servers[currentServer]):
			if res.err != nil {
				return 0, res.err
			}

			totalSpeed += res.speed
			totalExecuted++
		}

		currentServer = (currentServer + 1) % len(o.servers)
	}

	return totalSpeed / float64(totalExecuted), nil
}

func (o *ooklaSpeeedtestNet) upload(server *speedtest.Server) chan networkActionResult {
	resChan := make(chan networkActionResult)

	go func() {
		err := server.UploadTest(false)
		if err != nil {
			resChan <- networkActionResult{0, err}
			return
		}

		resChan <- networkActionResult{server.ULSpeed, nil}
	}()

	return resChan
}

func (o *ooklaSpeeedtestNet) download(server *speedtest.Server) chan networkActionResult {
	resChan := make(chan networkActionResult)

	go func() {
		err := server.DownloadTest(false)
		if err != nil {
			resChan <- networkActionResult{0, err}
			return
		}

		resChan <- networkActionResult{server.DLSpeed, nil}
	}()

	return resChan
}
