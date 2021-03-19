# Internet speed test
This package tests Internet speeds via fast.com or speedtest.net (you can choose between two).

**Note: the test coverage is less than 80% since in this specific case higher code coverage is not reasonable. I would need to simulate errors on requests to speed testing providers, considering size of the package (and task's timeframe) it might not make sense doing so**

## How to use
The repository contains `main.go` file which performs speed testing using both providers.
If you would like to run this from your code, there is the only method: `speedtest.Test(provider speedtest.Provider)`, for usage example see `main.go`.

## How to test
`go test github.com/merisho/nevo-coding-test/speedtest`

## How it works
In both speed testing providers, the code sends requests until the timeout is reached. Timeout is configured to be 10 seconds.
When the timeout is reached, average speed (Mbps) is calculated by all requests that managed to finish in the given time.

### Speedtest.net
Speed testing via Ookla's speedtest.net is implemented with `github.com/showwin/speedtest-go` package.

### Fast.com
Speed testing via Netflix's fast.com is implemented with `gopkg.in/ddo/go-fast.v0`. However, this package lacks upload speed testing. So I have written it myself.
There is an obvious issue that I acknowledge: timeout for testing download speed is hardcoded in the `go-fast` package and there is no way to change it. It is set to be 10 seconds, so I put the same value for the timeout in my code.
There is a clear way to write a download speed testing (very similar to upload speed testing), however I decided to use the above mentioned package for this to not complicate matter further. 
