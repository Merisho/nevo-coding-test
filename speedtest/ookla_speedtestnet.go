package speedtest

import "github.com/showwin/speedtest-go/speedtest"

func newOoklaSpeedtestNet() (*ooklaSpeeedtestNet, error) {
    user, err := speedtest.FetchUserInfo()
    if err != nil {
        return nil, err
    }

    serverList, err := speedtest.FetchServerList(user)
    if err != nil {
        return nil, err
    }

    targets, err := serverList.FindServer(nil)
    if err != nil {
        return nil, err
    }

    return &ooklaSpeeedtestNet{
        servers: targets,
    }, nil
}

type ooklaSpeeedtestNet struct {
    servers speedtest.Servers
}

func (o *ooklaSpeeedtestNet) Test() (upload, download float64, err error) {
    var downloadSum float64
    var uploadSum float64

    for _, s := range o.servers {
        err = s.UploadTest(false)
        if err != nil {
            return 0, 0, err
        }

        err = s.DownloadTest(false)
        if err != nil {
            return 0, 0, err
        }

        downloadSum += s.DLSpeed
        uploadSum += s.ULSpeed
    }

    l := float64(len(o.servers))
    return uploadSum / l, downloadSum / l, nil
}
