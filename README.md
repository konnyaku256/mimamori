# mimamori

「mimamori」は赤ちゃんのための見守りカメラシステムです。
Web カメラが捉えた映像と音声を WebRTC で配信し、ブラウザからストリーミングできるようにしています。

![](./docs/img/web-frontend-preview.png)

## 依存関係

- ハードウェア
  - [Raspberry Pi 4 Model B](https://www.raspberrypi.com/products/raspberry-pi-4-model-b/)
  - [Logicool C270n HD ウェブカメラ](https://www.logicool.co.jp/ja-jp/products/webcams/hd-webcam-c270n.960-001265.html)
- ソフトウェア
  - [WebRTC Native Client Momo](https://github.com/shiguredo/momo)

## システム構成図

![](./docs/img/mimamori-architecture.png)

## 導入手順

### 0 注意事項

1. Raspberry Pi に Raspberry Pi OS はインストール済みとします。
2. Raspberry Pi のローカル IP アドレスは「192.168.0.10」が割り当てられていることを前提としています。
   環境に合わせて適宜読み替えてください。

### 1 WebRTC Native Client Momo の導入

次の 2 つの公式ドキュメントに従って、WebRTC Native Client Momo を Raspberry Pi に導入・動作確認します。

- [Raspberry Pi (Raspberry-Pi-OS) で Momo を使ってみる](https://github.com/shiguredo/momo/blob/develop/doc/SETUP_RASPBERRY_PI.md)
- [テストモードを利用して Momo を動かしてみる](https://github.com/shiguredo/momo/blob/develop/doc/USE_TEST.md)

WebRTC Native Client Momo のバイナリおよび関連ファイルは /home/pi/mimamori 下に展開しておきます。

### 2 Web カメラ制御 API サーバーの導入

mimamori-camera-controller は [v4l2-utils](https://git.linuxtv.org/v4l-utils.git) に含まれる v4l2-ctl コマンドを使用して、Web カメラの設定を制御するための API サーバーです。

まず、v4l2-ctl コマンドを Raspberry Pi で使えるようにするため、次のコマンドを実行して v4l2-utils をインストールします。

```
pi@raspberrypi:~ $ sudo apt install v4l-utils
```

次に Raspberry Pi 向けにビルド済みの API サーバーのバイナリを scp で mac から Raspberry Pi に転送します。（[./backend/Makefile](./backend/Makefile) の scp の項目を参照ください。）

```
$ make scp
```

### 3 Web Frontend の導入

この Web Frontend は WebRTC Native Client Momo が WebRTC で配信する映像と音声をストリーミングします。

手順 1 で導入した WebRTC Native Client Momo のバイナリを配置したディレクトリの html 下に [./frontend](./frontend/) のファイルを配置します。

### 3 Systemd でサービス化

次の 2 つの Systemd のサービスを /etc/systemd/system 下に配置します。

- [/etc/systemd/system/mimamori-webrtc-server.service](./etc/systemd/system/mimamori-webrtc-server.service)
  - 手順 1 で導入した WebRTC Native Client Momo のバイナリを test モードで実行するサービス
- [/etc/systemd/system/mimamori-camera-controller.service](./etc/systemd/system/mimamori-camera-controller.service)
  - 手順 2 で導入した Web カメラ制御 API サーバーのバイナリを実行するサービス

次のコマンドを実行して、サービスを再読み込みします。

```
pi@raspberrypi:~ $ sudo systemctl daemon-reload
```

次のコマンドを実行して、OS の起動時にサービスが自動的に実行されるようにします。

```
pi@raspberrypi:~ $ sudo systemctl enable mimamori-webrtc-server.service
pi@raspberrypi:~ $ sudo systemctl enable mimamori-camera-controller.service
```

### 4 ディレクトリ構成の確認

最終的なディレクトリ構成は次のようになります。

Web Frontend と Backend

```
pi@raspberrypi:~/mimamori $ tree .
.
|-- html
|   |-- camera-controller-client.js
|   |-- favicon-96x96.png
|   |-- mimamori.html
|   `-- webrtc.js
|-- mimamori-camera-controller
`-- momo
```

Systemd のサービス

```
pi@raspberrypi:/etc/systemd/system $ ls
mimamori-webrtc-server.service
mimamori-camera-controller.service
```

## 使い方

Raspberry Pi と同じローカルエリアネットワークに接続されているデバイスから http://192.168.0.10:8080/html/mimamori.html にアクセスします。
画面右上「接続」ボタンを押して、ストリーミングを開始します。
また、画面右上「昼モード」、「夜モード」ボタンを押して、カメラの設定を切り替えます。
