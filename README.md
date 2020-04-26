http echo
---------


natの設定確認のための簡易webサーバーです。
単純にこのwebサーバーから見たリモートクライアントのipアドレスを返すだけです。

pingだけではtcpがちゃんと疎通できるのか頼りないので、これをラズパイにインストールして検証したいVLANのアクセスポートに接続して使用してました。
なので `make` は僕のプライベートラズパイmodel2向けのバイナリを生成します。


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
direnvなどを利用して環境変数で指定してください。

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
