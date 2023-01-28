package v4l2

type CameraControl struct {
	CameraMode string
	CommandName string
	CommandArguments []string
}

var (
	Noon = CameraControl {
		CameraMode: "noon",
		CommandName: "v4l2-ctl",
		CommandArguments: []string{"-d", "/dev/video0", "-c", "exposure_auto=3"},
	}
	Night = CameraControl {
		CameraMode: "night",
		CommandName: "v4l2-ctl",
		CommandArguments: []string{"-d", "/dev/video0", "-c", "exposure_auto=1", "-c", "exposure_absolute=9900"},
	}
)

func SelectCameraControl(cameraMode string) CameraControl {
	switch cameraMode {
	case Noon.CameraMode:
		return Noon
	case Night.CameraMode:
		return Night
	default:
		return Noon
	}
}