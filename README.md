# speeedy [![Go Report Card](https://goreportcard.com/badge/github.com/faoztas/speeedy)](https://goreportcard.com/report/github.com/faoztas/speeedy) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/04d432e424ce4d6aabf9c6237d9d304c)](https://app.codacy.com/manual/faoztas/speeedy?utm_source=github.com&utm_medium=referral&utm_content=faoztas/speeedy&utm_campaign=Badge_Grade_Dashboard) [![GolangCI](https://golangci.com/badges/github.com/faoztas/speeedy.svg)](https://golangci.com/r/github.com/faoztas/speeedy) [![Go 1.13+](https://img.shields.io/badge/go-1.13-9cf.svg)](https://golang.org/dl/) [![platforms](https://img.shields.io/badge/platforms-linux|windows|macos-active.svg)]() [![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

Speeedy is a SpeedTest Telegram bot. Speedtest only runs on the machine it is installed on. Speedtest can also be run routinely with Job.

I developed this bot because I was curious about the performance of the my ISP.

### Prerequisites

* Go 1.13+
* Vgo (Go Module)
* Telegram

### Installing speeedy and Run

* Create a telegram bot with BotFather and get token of created bot.
* Set required environment variables

```bash
$ cd $GOPATH/src
$ mkdir github.com
$ cd github.com
$ git clone git@github.com:faoztas/speeedy.git
$ cd speeedy                # config.toml template is filled. Moves to the project directory.
$ export GO111MODULE=on     # manually active module mode
$ go mod init               # manually e.g., $ go mod init github.com/my/repo
$ go mod tidy               # This command allows you to fetch all the dependencies that you need for testing in your module.
$ go build -o speeedy       # Or $ go run .
$ ./speeedy
```
