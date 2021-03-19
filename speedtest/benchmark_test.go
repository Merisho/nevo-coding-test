package speedtest

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func BenchmarkNetflixFastCom_Test(b *testing.B) {
    fastCom, err := newNetflixFastCom()
    assert.NoError(b, err)

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        upload, download, err := fastCom.Test()
        assert.NoError(b, err)

        fmt.Println("Fast.com upload and download speeds", upload, download)
    }
}

func BenchmarkOoklaSpeeedtestNet_Test(b *testing.B) {
    speedtest, err := newOoklaSpeedtestNet()
    assert.NoError(b, err)

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        upload, download, err := speedtest.Test()
        assert.NoError(b, err)

        fmt.Println("Speedtest.net upload and download speeds", upload, download)
    }
}
