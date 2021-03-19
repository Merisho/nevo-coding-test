package speedtest

import (
    "bytes"
    "errors"
    "fmt"
    "gopkg.in/ddo/go-fast.v0"
    "log"
    "net/http"
    "strconv"
    "strings"
    "time"
)

const (
    uploadSize = 24 * 1024 * 1024
)

func newNetflixFastCom() (*netflixFastCom, error) {
    fastCom := fast.New()

    err := fastCom.Init()
    if err != nil {
        return nil, err
    }

    urls, err := fastCom.GetUrls()
    if err != nil {
        return nil, err
    }

    return &netflixFastCom{
        fastCom: fastCom,
        servers: urls,
    }, nil
}

type netflixFastCom struct {
    fastCom *fast.Fast
    servers []string
}

func (n *netflixFastCom) Test() (upload, download float64, err error) {
    upload, err = n.testUploadSpeed()
    if err != nil {
        return 0, 0, err
    }

    download, err = n.testDownloadSpeed()
    if err != nil {
        return 0, 0, err
    }

    return upload, download, nil
}

func (n *netflixFastCom) testUploadSpeed() (float64, error) {
    var totalSpeed float64
    totalExecuted := 0
    timeout := time.After(10 * time.Second)
    currentServer := 0

    loop:
        for {
            server := n.servers[currentServer]
            select {
            case <- timeout:
                if totalExecuted == 0 {
                    return 0, errors.New("timeout")
                }

                break loop
            case res := <- n.upload(server):
                if res.err != nil {
                    return 0, res.err
                }

                totalSpeed += res.speed
                totalExecuted++
            }

            currentServer = (currentServer + 1) % len(n.servers)
        }

    return totalSpeed / float64(totalExecuted), nil
}

func (n *netflixFastCom) upload(url string) chan networkActionResult {
    resChan := make(chan networkActionResult)
    var httpClient http.Client

    go func() {
        u := n.makeUploadURL(url)
        if u == "" {
            resChan <- networkActionResult{0, fmt.Errorf("invalid url %s", url)}
            return
        }

        data := make([]byte, uploadSize)
        req, err := n.createUploadRequest(u, data)
        if err != nil {
            log.Fatal(err)
        }

        start := time.Now()
        res, err := httpClient.Do(req)
        duration := time.Now().Sub(start).Seconds()
        if err != nil {
            resChan <- networkActionResult{0, err}
            return
        }

        if res.StatusCode != 200 {
            resChan <- networkActionResult{0, errors.New("fast.com response code is not 200")}
            return
        }

        dataMegabits := float64(len(data) * 8) / 1000000
        resChan <- networkActionResult{dataMegabits / duration, nil}
        return
    }()

    return resChan
}

func (n *netflixFastCom) createUploadRequest(url string, data []byte) (*http.Request, error) {
    req, err := http.NewRequest("POST", url, bytes.NewReader(data))
    if err != nil {
        return nil, err
    }

    req.Header.Add("Connection", "keep-alive")
    req.Header.Add("Content-Length", strconv.Itoa(len(data)))
    req.Header.Add("Content-type", "application/octet-stream")

    return req, nil
}

func (n *netflixFastCom) makeUploadURL(u string) string {
    p := strings.Split(u, "?")
    if len(p) < 2 {
        return ""
    }

    return p[0] + "/range/0-26214400?" + p[1]
}

func (n *netflixFastCom) testDownloadSpeed() (float64, error) {
    kbpsChan := make(chan float64)

    var downloaded float64
    var count float64
    go func() {
        for kbps := range kbpsChan {
            count++
            mbps := kbps / 1000
            downloaded += mbps
        }
    }()

    err := n.fastCom.Measure(n.servers, kbpsChan)
    if err != nil {
        return 0, err
    }

    return downloaded / count, nil
}
