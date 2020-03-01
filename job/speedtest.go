package job

import s "github.com/speeedy/adapters/speedtest"

func SpeedTest() {
	r := s.SpeedTest(s.UnitOfMeasure, s.UnitMibps, s.SetFormat, s.JsonPrettyFormat)
	r.ToDB(0)
}
