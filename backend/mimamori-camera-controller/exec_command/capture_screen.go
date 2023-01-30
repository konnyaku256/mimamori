package exec_command

var (
	CaptureScreen = Command {
		Name: "v4l2-ctl",
		Arguments: []string{"-d", "/dev/video0", "-c", "exposure_auto=3"},
	}
)