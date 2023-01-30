const API_URL = "http://192.168.0.100:3000/api";
const API_VERSION = "v1";
const BASE_URL = API_URL + "/" + API_VERSION;

function changeCameraMode(cameraMode) {
  fetch(BASE_URL + "/exec/change-camera-mode" + "?mode=" + cameraMode).catch(
    (error) => console.log(error)
  );
}

function captureScreen() {
  fetch(BASE_URL + "/exec/capture-screen").catch((error) => console.log(error));
}
