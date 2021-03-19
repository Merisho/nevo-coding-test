package main

import (
    "fmt"
    "github.com/merisho/nevo-coding-test/speedtest"
    "log"
)

func main() {
    upload, download, err := speedtest.Test(speedtest.SpeedtestNet)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Speedtest.net upload speed: %f Mbps\n", upload)
    fmt.Printf("Speedtest.net download speed: %f Mbps\n", download)

    upload, download, err = speedtest.Test(speedtest.FastCom)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Fast.com upload speed: %f Mbps\n", upload)
    fmt.Printf("Fast.com download speed: %f Mbps\n", download)
}
