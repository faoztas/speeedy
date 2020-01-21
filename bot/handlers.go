package bot

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"

	s "speedy/adapters/speedtest"
)

type Handler struct {
	Bot  *telebot.Bot
}

func (h *Handler) Test(m *telebot.Message){
	var result interface{}
	logrus.WithField("sender_id", m.Sender.ID).WithField("username", m.Sender.Username).Info(m.Payload)
	if m.Payload == "" {
		result = s.SpeedTest()
		h.Bot.Send(m.Sender, string(result.([]byte)))
	} else if m.Payload == "details"{
		result = s.SpeedTest(s.SetServer, "23316", s.UnitOfMeasure, s.UnitMibps, s.SetFormat, s.JsonPrettyFormat)
		h.Bot.Send(m.Sender, result.(s.Result))
	} else if m.Payload == "image" {
		result = s.SpeedTest(s.SetServer, "23316", s.UnitOfMeasure, s.UnitMibps, s.SetFormat, s.JsonPrettyFormat)
		photo := & telebot.Photo{File: telebot.FromURL(result.(s.Result).Result.URL +".png")}
		h.Bot.SendAlbum(m.Sender, telebot.Album{photo})
	}
}

func (h *Handler) ServerLists(m *telebot.Message){
	logrus.WithField("sender_id", m.Sender.ID).WithField("username", m.Sender.Username).Info(m.Payload)
	/*if m.Payload != "" {
		result := s.SpeedTest(s.SetServer, "23316", s.UnitOfMeasure, s.UnitMibps, s.SetFormat, s.JsonPrettyFormat)
	}
	result := s.SpeedTest(s.SetServer, "23316", s.UnitOfMeasure, s.UnitMibps, s.SetFormat, s.JsonPrettyFormat)
	photo := & telebot.Photo{File: telebot.FromURL(result+".png")}
	h.Bot.SendAlbum(m.Sender, telebot.Album{photo})*/
}

func (h *Handler) Start(m *telebot.Message) {
	if m.Payload == "" {
		logrus.WithField("if", m.Sender.Username).Info(m.Text)
	} else {
		logrus.WithField("else", m.Sender.Username).Info(m.Text)

	}
}

func (h *Handler) Help(m *telebot.Message) {
	logrus.WithField("help", m.Sender.Username).Info(m.Text)
}

func (h *Handler) Text(m *telebot.Message) {
	logrus.WithField("sender_id", m.Sender.ID).WithField("username", m.Sender.Username).Info(m.Text)


}
