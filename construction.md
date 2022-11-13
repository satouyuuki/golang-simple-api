デプロイ手順

- ec2インスタンスを作成する
    - セキュリティグループはインバウンドルールにssh, httpを追加する
- ec2にsshで接続する
- goをダウンロードして解凍する(usr/localに)

```
$ cd /tmp
$ wget https://go.dev/dl/go1.19.3.linux-amd64.tar.gz
$ tar -xvf go1.19.3.linux-amd64.tar.gz 
$ sudo mv go /usr/local
```

- goの環境変数を設定する(pathを通す)
    - GOPATHを`~/go`に, GOROOTを`/usr/local/go`に設定し、PATHに`$GOPATH`, `$GOROOT/bin`を追加して反映する
- ~/goにbinとpkgディレクトリを作成する
- 自分が作成したモジュールのバイナリをGOPATHにインストールする

```
$ go install github.com/satouyuuki/golang-simple-api@latest
# 何かしらの理由で失敗したらcacheを消す
$ go clean -modcache
```
- nginxをインストールする
    - default.confファイルを編集する

```
events {
    worker_connections 1024;
}

http {
        server {
                listen       80;
                listen       [::]:80;
                server_name  _;

                location / {
                        proxy_pass http://127.0.0.1:8080;
                }
        }
}
```
- systemctlにデーモンの設定ファイルを作成する

```
[Unit]
Description=A simple rest api that runs on Go.

[Service]
WorkingDirectory=/home/ec2-user/go
ExecStart=/home/ec2-user/go/bin/golang-simple-api
Restart=always

[Install]
WantedBy=multi-user.target
```
    - デーモンを自動起動の設定をしスタートをする