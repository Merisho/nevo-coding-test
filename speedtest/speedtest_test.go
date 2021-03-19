package speedtest

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOoklaSpeedtest(t *testing.T) {
    upload, download, err := Test(SpeedtestNet)
	assert.NoError(t, err)
	assert.Greater(t, upload, float64(0))
	assert.Greater(t, download, float64(0))

	fmt.Println(upload, download)
}

func TestNetflixSpeedTest(t *testing.T) {
	upload, download, err := Test(FastCom)
	assert.NoError(t, err)
	assert.Greater(t, upload, float64(0))
	assert.Greater(t, download, float64(0))

	fmt.Println(upload, download)
}
