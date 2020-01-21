package speedtest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/sirupsen/logrus"
)

func SpeedTest(args ...string) interface{}{
	commands := make([]string, len(args))

	for index, value := range args {
		commands[index] = value
	}

	cmd := exec.Command(Path, commands...)
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

	fmt.Printf("%s\n", bOut)
	err = json.Unmarshal(bOut, &result)
	if err != nil {
		logrus.Error(err)
		return bOut
	}

	if err := cmd.Wait(); err != nil {
		logrus.Error(err)
	}

	return result
}
