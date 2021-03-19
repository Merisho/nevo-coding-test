package speedtest

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOoklaSpeedtestNet(t *testing.T) {
	speedtest, err := newOoklaSpeedtestNet()
	assert.NoError(t, err)

	upload, download, err := speedtest.Test()
	assert.NoError(t, err)
	assert.Greater(t, upload, float64(0))
	assert.Greater(t, download, float64(0))

	fmt.Println("Speedtest.net upload and download speeds:", upload, download)
}
