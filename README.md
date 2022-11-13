### ローカル環境
go version go1.19.2 darwin/arm64

### モジュールを初期化する
$ go mod init example.com/m

- go.modは依存関係を管理するファイル(package.json的な)
- goには簡単にreat apiを作成できるginというパッケージがある

### 依存関係をgo.modに追加する
$ go get .

- go.sumファイルがpackage.lock.json的なやつ

### example.httpファイルの使い方
vscodeにREST ClientというExtensionsをインストールして使ってください

### module pathが違うと怒られた(go mod init example.com/mのせいで)
以下だとダメ
go mod edit -replace example.com/m@v0.1=github.com/satouyuuki/golang-simple-api@v0.1
以下参考にファイルを削除して再作成した
https://stackoverflow.com/questions/65512353/how-to-undo-go-mod-init
