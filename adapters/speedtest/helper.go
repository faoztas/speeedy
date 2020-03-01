package speedtest

import (
	"fmt"
	"math"

	"github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"

	"github.com/speeedy/db"
	"github.com/speeedy/models"
	"github.com/speeedy/utils"
)

func (r *Result) Response() string {
	if r == nil {
		return utils.TestFailMsg
	}
	downloadSpeed := float64(r.Download.Bandwidth) * 8 / 1000000
	uploadSpeed := float64(r.Upload.Bandwidth) * 8 / 1000000

	return fmt.Sprintf(utils.TestF, downloadSpeed, uploadSpeed,
		int(math.Round(r.Ping.Jitter)), int(math.Round(r.Ping.Latency)))
}

func (r *Result) ResponseWithDetails() string {
	if r == nil {
		return utils.TestFailMsg
	}
	var isVpn string
	downloadSpeed := float64(r.Download.Bandwidth) * 8 / 1000000
	uploadSpeed := float64(r.Upload.Bandwidth) * 8 / 1000000

	switch r.Interface.IsVpn {
	case true:
		isVpn = utils.Enable
	case false:
		isVpn = utils.Disable
	}

	return fmt.Sprintf(utils.TestDetailsF, downloadSpeed, uploadSpeed, r.Ping.Jitter, r.Ping.Latency, r.Isp,
		r.Server.Name, r.Server.Location, r.Server.Country, r.Interface.ExternalIP, r.Interface.InternalIP, isVpn,
		r.Result.URL)
}

func (r *Result) ResponseWithImage() *telebot.Photo {
	if r == nil {
		return nil
	}
	return &telebot.Photo{File: telebot.FromURL(r.Result.URL + ".png")}
}

func (r *Result) ToDB(senderID int) {
	var user models.User
	if err := db.DB.Where("sender_id = ?", senderID).Find(&user).Error; err != nil {
		logrus.Error(err)
		return
	}
	result := models.Result{
		User:              &user,
		UserID:            user.ID,
		Jitter:            r.Ping.Jitter,
		Latency:           r.Ping.Latency,
		DownloadBandwidth: r.Download.Bandwidth,
		DownloadBytes:     r.Download.Bytes,
		DownloadElapsed:   r.Download.Elapsed,
		UploadBandwidth:   r.Upload.Bandwidth,
		UploadBytes:       r.Upload.Bytes,
		UploadElapsed:     r.Upload.Elapsed,
		Isp:               r.Isp,
		UserInternalIP:    r.Interface.InternalIP,
		UserName:          r.Interface.Name,
		UserMacAddr:       r.Interface.MacAddr,
		UserIsVpn:         r.Interface.IsVpn,
		UserExternalIP:    r.Interface.ExternalIP,
		ServerID:          r.Server.ID,
		ServerName:        r.Server.Name,
		ServerLocation:    r.Server.Location,
		ServerCountry:     r.Server.Country,
		ServerHost:        r.Server.Host,
		ServerPort:        r.Server.Port,
		ServerIP:          r.Server.IP,
		ResultID:          r.Result.ID,
		ResultURL:         r.Result.URL,
	}

	if err := db.DB.Create(&result).Error; err != nil {
		logrus.Error(err)
		return
	}
}
