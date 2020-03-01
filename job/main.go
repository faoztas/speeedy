package job

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/speeedy/config"
	"github.com/speeedy/utils"
)

func Jobs() {
	logrus.Info("Starting Scheduling jobs")

	if config.Env.SpeedTest.Flag {
		utils.RunRoutineForever(SpeedTest, 1*time.Hour)
	} else {
		logrus.Info("SpeedTest not enabled")
	}
}
