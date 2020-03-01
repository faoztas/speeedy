package bot

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"

	s "github.com/speeedy/adapters/speedtest"
	"github.com/speeedy/db"
	"github.com/speeedy/models"
	"github.com/speeedy/utils"
)

type Handler struct {
	Bot *telebot.Bot
}

func (h *Handler) Test(m *telebot.Message) {
	logrus.WithField(senderID, m.Sender.ID).WithField(username, m.Sender.Username).Info(m.Text)

	if db.DB.Where("sender_id = ?", m.Sender.ID).Find(&models.User{}).RowsAffected == 0 {
		utils.User(*m)
	}

	h.Bot.Send(m.Sender, utils.TestInProgressMsg)
	if m.Payload == null {
		r := s.SpeedTest(s.UnitOfMeasure, s.UnitMibps, s.SetFormat, s.JsonPrettyFormat)
		msg := r.Response()
		r.ToDB(m.Sender.ID)
		h.Bot.Send(m.Sender, msg)
	} else if m.Payload == details {
		r := s.SpeedTest(s.UnitOfMeasure, s.UnitMibps, s.SetFormat, s.JsonPrettyFormat)
		msg := r.ResponseWithDetails()
		r.ToDB(m.Sender.ID)
		h.Bot.Send(m.Sender, msg)
	} else if m.Payload == image {
		r := s.SpeedTest(s.UnitOfMeasure, s.UnitMibps, s.SetFormat, s.JsonPrettyFormat)
		msg := r.ResponseWithImage()
		r.ToDB(m.Sender.ID)
		if msg != nil {
			h.Bot.SendAlbum(m.Sender, telebot.Album{msg})
		} else {
			h.Bot.Send(m.Sender, utils.TestFailMsg)
		}
	} else if strings.Contains(m.Payload, server) {
		serverID := strings.Split(m.Payload, server)[1]
		_, err := strconv.Atoi(serverID)
		if err != nil {
			logrus.Error(err)
			h.Bot.Send(m.Sender, utils.TestFailWithServerMsg)
		}
		r := s.SpeedTest(s.SetServer, serverID, s.UnitOfMeasure, s.UnitMibps, s.SetFormat, s.JsonPrettyFormat)
		msg := r.ResponseWithDetails()
		r.ToDB(m.Sender.ID)
		h.Bot.Send(m.Sender, msg)
	} else {
		h.Bot.Send(m.Sender, utils.TestUnexpectedMsg)
	}
}

func (h *Handler) Servers(m *telebot.Message) {
	logrus.WithField(senderID, m.Sender.ID).WithField(username, m.Sender.Username).Info(m.Text)
	servers := s.GetServers()
	var msg string
	for i := range servers.Servers {
		msg += fmt.Sprintf(utils.ServerF, i+1, servers.Servers[i].Name, servers.Servers[i].ID,
			servers.Servers[i].Location, servers.Servers[i].Country)
	}

	h.Bot.Send(m.Sender, msg)
}

func (h *Handler) Start(m *telebot.Message) {
	logrus.WithField(senderID, m.Sender.ID).WithField(username, m.Sender.Username).Info(m.Text)
	utils.User(*m)

	h.Bot.Send(m.Sender, utils.StartMsg)
}

func (h *Handler) Help(m *telebot.Message) {
	logrus.WithField(senderID, m.Sender.ID).WithField(username, m.Sender.Username).Info(m.Text)

	h.Bot.Send(m.Sender, utils.HelpMsgTest)
	h.Bot.Send(m.Sender, utils.HelpMsgServer)
	h.Bot.Send(m.Sender, utils.HelpMsgLastBestWorst)
}

func (h *Handler) Text(m *telebot.Message) {
	logrus.WithField(senderID, m.Sender.ID).WithField(username, m.Sender.Username).Info(m.Text)
	h.Bot.Send(m.Sender, utils.NotFoundCmdMsg)
}

func (h *Handler) Last(m *telebot.Message) {
	logrus.WithField(senderID, m.Sender.ID).WithField(username, m.Sender.Username).Info(m.Text)
	var limit int
	var err error

	if m.Payload == null {
		limit = 1
	} else {
		limit, err = strconv.Atoi(m.Payload)
		if err != nil {
			h.Bot.Send(m.Sender, utils.LastFailMsg)
			return
		}
	}

	var result []models.Result
	if err := db.DB.Order("created_at desc").Limit(limit).Find(&result).Error; err != nil {
		h.Bot.Send(m.Sender, utils.DBFailMsg)
		return
	}
	for i := range result {
		h.Bot.Send(m.Sender, result[i].ToMsg())
	}
}

func (h *Handler) Best(m *telebot.Message) {
	logrus.WithField(senderID, m.Sender.ID).WithField(username, m.Sender.Username).Info(m.Text)

	var result models.Result
	if err := db.DB.Order("download_bandwidth desc").First(&result).Error; err != nil {
		h.Bot.Send(m.Sender, utils.DBFailMsg)
		return
	}
	h.Bot.Send(m.Sender, result.ToMsg())
}

func (h *Handler) Worst(m *telebot.Message) {
	logrus.WithField(senderID, m.Sender.ID).WithField(username, m.Sender.Username).Info(m.Text)

	var result models.Result
	if err := db.DB.Order("download_bandwidth asc").First(&result).Error; err != nil {
		h.Bot.Send(m.Sender, utils.DBFailMsg)
		return
	}
	h.Bot.Send(m.Sender, result.ToMsg())
}
