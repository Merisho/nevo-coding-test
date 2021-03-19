package speedtest

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestNetflixFastCom_upload(t *testing.T) {
    st, err := newNetflixFastCom()
    assert.NoError(t, err)

    upload, err := st.testUploadSpeed()
    assert.NoError(t, err)
    assert.Greater(t, upload, float64(0))

    fmt.Println("Fast.com upload speed:", upload)
}

func TestNetflixFastCom_download(t *testing.T) {
    st, err := newNetflixFastCom()
    assert.NoError(t, err)

    download, err := st.testDownloadSpeed()
    assert.NoError(t, err)
    assert.Greater(t, download, float64(0))

    fmt.Println("Fast.com download speed:", download)
}
