package speedtest

import (
    "bytes"
    "errors"
    "gopkg.in/ddo/go-fast.v0"
    "log"
    "net/http"
    "strconv"
    "strings"
    "time"
)

const (
    uploadSize = 24 * 1024 * 1024
    uploadSizeBits = uploadSize * 8
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
    upload, err = n.testUploadSpeed(n.servers)
    if err != nil {
        return 0, 0, err
    }

    download, err = n.testDownloadSpeed(n.servers)
    if err != nil {
        return 0, 0, err
    }

    return upload, download, nil
}

func (n *netflixFastCom) testUploadSpeed(urls []string) (float64, error) {
    var client http.Client
    var total float64
    for _, u := range urls {
        url := n.makeUploadURL(u)
        if url == "" {
            continue
        }

        req, err := n.createUploadRequest(url)
        if err != nil {
            log.Fatal(err)
        }

        start := time.Now()
        res, err := client.Do(req)
        secs := time.Now().Sub(start).Seconds()
        if err != nil {
            return 0, nil
        }

        if res.StatusCode != 200 {
            return 0, errors.New("fast.com response code is not 200")
        }

        total += uploadSizeBits / secs / 1000000
    }

    return total / float64(len(urls)), nil
}

func (n *netflixFastCom) testDownloadSpeed(urls []string) (float64, error) {
    kbps := make(chan float64)

    var downloaded float64
    var count float64
    go func() {
        for k := range kbps {
            count++
            downloaded += k / 1000
        }
    }()

    err := n.fastCom.Measure(urls, kbps)
    if err != nil {
        return 0, err
    }

    return downloaded / count, nil
}

func (n *netflixFastCom) createUploadRequest(url string) (*http.Request, error) {
    data := make([]byte, uploadSize)

    req, err := http.NewRequest("POST", url, bytes.NewReader(data))
    if err != nil {
        return nil, err
    }

    req.Header.Add("Connection", "keep-alive")
    req.Header.Add("Content-Length", strconv.Itoa(uploadSize))
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
