package speedtest

import (
	"encoding/json"
	"io/ioutil"
	"os/exec"

	"github.com/sirupsen/logrus"

	"github.com/speeedy/config"
)

func SpeedTest(args ...string) *Result {
	commands := make([]string, len(args))

	for index, value := range args {
		commands[index] = value
	}

	cmd := exec.Command(config.Env.SpeedTest.Path, commands...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logrus.Error(err)
	}

	if err := cmd.Start(); err != nil {
		logrus.Error(err)
	}

	bOut, err := ioutil.ReadAll(stdout)
	if err != nil {
		logrus.Error(err)
	}

	var result Result

	err = json.Unmarshal(bOut, &result)
	if err != nil {
		logrus.Error(err)
	}

	if err := cmd.Wait(); err != nil {
		logrus.Error(err)
	}

	return &result
}

func GetServers() *ServerList {
	cmd := exec.Command(config.Env.SpeedTest.Path, Servers, SetFormat, JsonPrettyFormat)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logrus.Error(err)
	}

	if err := cmd.Start(); err != nil {
		logrus.Error(err)
	}

	bOut, err := ioutil.ReadAll(stdout)
	if err != nil {
		logrus.Error(err)
	}

	var result ServerList

	err = json.Unmarshal(bOut, &result)
	if err != nil {
		logrus.Error(err)
	}

	if err := cmd.Wait(); err != nil {
		logrus.Error(err)
	}

	return &result
}
