package speedtest

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestOoklaSpeeedtestNet_upload(t *testing.T) {
    st, err := newOoklaSpeedtestNet()
    assert.NoError(t, err)

    upload, err := st.testSpeed(st.upload)
    assert.NoError(t, err)
    assert.Greater(t, upload, float64(0))

    fmt.Println("Speedtest.net upload speed:", upload)
}

func TestOoklaSpeeedtestNet_download(t *testing.T) {
    st, err := newOoklaSpeedtestNet()
    assert.NoError(t, err)

    download, err := st.testSpeed(st.download)
    assert.NoError(t, err)
    assert.Greater(t, download, float64(0))

    fmt.Println("Speedtest.net download speed:", download)
}
