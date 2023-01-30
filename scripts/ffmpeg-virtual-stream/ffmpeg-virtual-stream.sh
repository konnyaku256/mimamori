ffmpeg \
  -f video4linux2 -input_format mjpeg -video_size 1280x720 -framerate 30 -i /dev/video0 \
  -f video4linux2 -vcodec copy /dev/video2 \
  -f video4linux2 -vcodec copy /dev/video3