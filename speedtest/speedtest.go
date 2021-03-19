package speedtest

type Provider string

const (
	SpeedtestNet Provider = "speedtest.net"
	FastCom      Provider = "fast.com"
)

type SpeedTester interface {
	Test() (upload, download float64, err error)
}

type networkActionResult struct {
	speed float64
	err   error
}

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
