package exec_command

type cameraMode struct {
	cameraMode string
	command Command
}

var (
	noon = cameraMode {
		cameraMode: "noon",
		command: Command {
			Name: "v4l2-ctl",
			Arguments: []string{"-d", "/dev/video0", "-c", "exposure_auto=3"},
		},
	}
	night = cameraMode {
		cameraMode: "night",
		command: Command {
			Name: "v4l2-ctl",
			Arguments: []string{"-d", "/dev/video0", "-c", "exposure_auto=1", "-c", "exposure_absolute=9900"},
		},
	}
)

func SelectCommandByCameraMode(cameraMode string) Command {
	switch cameraMode {
	case noon.cameraMode:
		return noon.command
	case night.cameraMode:
		return night.command
	default:
		return noon.command
	}
}