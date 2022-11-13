### ローカル環境
go version go1.19.2 darwin/arm64

### モジュールを初期化する
```
$ go mod init github.com/satouyuuki/golang-simple-api
```

go.modは依存関係を管理するファイル(package.json的なやつ)

### 依存関係をgo.modに追加する
```
$ go get .
$ go get github.com/stretchr/testify/assert
```

go.sumファイルができるがpackage.lock.json的なやつ(だと思う)

### example.httpファイルの使い方
vscodeにREST ClientというExtensionsをインストールして使ってください

### ec2へのデプロイの仕方はconstruction.mdに記載しました。