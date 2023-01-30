import os
import sys

from dotenv import load_dotenv
import cv2
import requests

READED_IMAGE_BY_WEB_CAMERA_PATH = "screen_capture.jpg"

def saveImageByWebCamera():
    # Web カメラのオブジェクトを取得
    camera = cv2.VideoCapture(3) # /dev/video3
    if camera.isOpened() is False:
        print("Web カメラのオープンに失敗しました")
        sys.exit()

    # Web カメラの画像を読み取り
    _, frame = camera.read()

    cv2.imwrite(READED_IMAGE_BY_WEB_CAMERA_PATH, frame)
    

def uploadImageByLineNotify():
    apiUrl = "https://notify-api.line.me/api/notify" 
    apiToken = os.getenv("LINE_NOTIFY_API_TOKEN")
    headers = {"Authorization" : "Bearer "+ apiToken} 
    message =  "みまもりカメラの画面をキャプチャしました" 
    payload = {"message" :  message}
    files = {"imageFile": open(READED_IMAGE_BY_WEB_CAMERA_PATH, "rb")}
    _ = requests.post(
        apiUrl,
        params=payload,
        headers=headers,
        files=files
    )


def removeImage():
    os.remove(READED_IMAGE_BY_WEB_CAMERA_PATH)
    
if __name__ == "__main__":
    load_dotenv()

    saveImageByWebCamera()
    uploadImageByLineNotify()
    removeImage()
