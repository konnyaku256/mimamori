const serverURL = "http://192.168.0.10:3000/api/v1/exec-v4l2";

function setCameraMode(cameraMode) {
  fetch(serverURL + "?mode=" + cameraMode).catch((error) => console.log(error));
}
