package utils

const (
	// MSG
	TestInProgressMsg     = "Speed Test In Progress. Please Waiting... 🤓"
	TestFailMsg           = "Speed Test Failed. 😇"
	TestFailWithServerMsg = "Please only use Server ID. 🧐"
	TestUnexpectedMsg     = "Unexpected Command. 🥴"
	StartMsg              = "Welcome to my channel. 🥳\n" +
		"If you want to learn bot commands please write /help."
	HelpMsgTest = "If you want to Speed Test please write /test.\n" +
		"/test --> default basic test\n" +
		"/test details --> speed test with extras argument\n" +
		"/test image --> send image with speed test\n" +
		"/test server 1234 --> speed test with specific server"
	HelpMsgServer  = "If you want to learn Server List please write /servers."
	HelpMsgLastBestWorst = "If you want to learn Past test results please write /last.\n" +
		"/last --> last result\n" +
		"/last 5 --> last 5 results\n" +
		"/best --> best result\n" +
		"/worst --> worst result"
	NotFoundCmdMsg = "Hmm... What are you doing here? 🤔"
	LastFailMsg    = "Please only use number. 🧐"
	DBFailMsg      = "Hmm... Record not found. 🤔"

	// MSG Format
	TestF = "Download: %.1f Mbps - Upload: %.1f Mbps\n" +
		"Jitter: %d ms - Latency: %d ms"
	TestDetailsF = "Download/Upload: %.2f/%.2f Mbps\n" +
		"Jitter/Latency: %.2f/%.2f ms\n" +
		"ISP: %s\n" +
		"Server: %s - Location: %s/%s\n" +
		"External/Internal IP: %s/%s - VPN: %s\n" +
		"Result: %s"
	ServerF = "%d) Server: %s\n" +
		"ID: %d - Location: %s/%s\n"

	// Flag
	Enable  = "Enable"
	Disable = "Disable"
)
