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
    //assert.Greater(t, upload, float64(0))

    fmt.Println(upload)
}
