http echo
---------

natの設定確認のための簡易webサーバーです。
pingだけではtcpがちゃんと疎通できるのか頼りないので、これをラズパイにインストールして検証したいVLANのアクセスポートに接続して使用してました。
単純にwebサーバーから見たリモートクライアントのipアドレスを返すだけです。
[https://httpbin.org/ip](https://httpbin.org/ip)みたいなサービスがイントラ内に必要だったので実装しました。

`make` は僕のラズパイmodel2向けのバイナリを生成します。
[こちら](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63)を参考に必要に応じて`Makefile`を修正して、`go build`のオプションを修正してください。

-----

## REQUIREMENTS

- GO1.14

## RUN


- 直接実行する
```
go run cmd/main.go
```

- 動作確認
```
$ curl 0.0.0.0
127.0.0.1
```

## DEPLOYMENT

ビルドしたバイナリとunitファイルをラズパイへ配置します。

deployのスクリプトで`remote_user`と`remote_host` を参照します。
[direnv](https://github.com/direnv/direnv)などを利用して環境変数で指定してください。

- 例
  - remote_user
    - sshユーザー名
  - remote_host
    - ラズパイのホスト名 avahiなどで設定している名前

- ビルドしてラズパイへ転送
```
make
make deploy
```

- ラズパイにログインしてバイナリを`/usr/local/bin/`へ移動します。

```
ssh alarm@pi.local
mv echo /usr/local/bin/echo
```

- systemdの登録(初回のみ)

  - unit ファイルをsystemディレクトリに移動して、`systemctl enable`して`start`してください。

```
mv http-echo.service /etc/systemd/system/http-echo.service
systemctl enable http-echo
systemctl start http-echo
```

## REFERENCE

echoの実装はこちらを参考にさせていただきました。
- https://github.com/methane/echoserver
