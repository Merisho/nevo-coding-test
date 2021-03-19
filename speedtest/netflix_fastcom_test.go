package speedtest

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNetflixFastCom(t *testing.T) {
	speedtest, err := newNetflixFastCom()
	assert.NoError(t, err)

	upload, download, err := speedtest.Test()
	assert.NoError(t, err)
	assert.Greater(t, upload, float64(0))
	assert.Greater(t, download, float64(0))

	fmt.Println("Fast.com upload and download speeds:", upload, download)
}

func TestNetflixFastComInvalidURL(t *testing.T) {
	speedtest, err := newNetflixFastCom()
	assert.NoError(t, err)

	speedtest.servers = []string{"invalid-url.com"}

	upload, download, err := speedtest.Test()
	assert.EqualError(t, err, "invalid url invalid-url.com")
	assert.Equal(t, upload, float64(0))
	assert.Equal(t, download, float64(0))
}
