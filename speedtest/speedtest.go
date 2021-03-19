package speedtest

import "time"

// Provider describes the speed testing provider
type Provider string

const (
	// SpeedtestNet is the Ookla's speedtest.net speed testing provider
	SpeedtestNet Provider = "speedtest.net"

	// FastCom is the Netflix's fast.com speed testing provider
	FastCom      Provider = "fast.com"

	// testTimeout the timeout used in speed testing. The code will send requests until the timeout is reached.
	testTimeout = 10 * time.Second
)

// SpeedTester describes an interface speed testing provider must implement
type SpeedTester interface {
	Test() (upload, download float64, err error)
}

// networkActionResult can be returned after upload or download operation is completed
// It describes networking speed in Mbps and error if any
type networkActionResult struct {
	speed float64
	err   error
}

// Test performs the speed testing, returns upload and download speeds in Mbps
func Test(p Provider) (upload, download float64, err error) {
	if p == SpeedtestNet {
		speedtest, err := newOoklaSpeedtestNet()
		if err != nil {
			return 0, 0, err
		}

		return speedtest.Test()
	}

	fastCom, err := newNetflixFastCom()
	if err != nil {
		return 0, 0, err
	}

	return fastCom.Test()
}
