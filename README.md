### モジュールを初期化する
$ go mod init example.com/m

- go.modは依存関係を管理するファイル(package.json的な)
- goには簡単にreat apiを作成できるginというパッケージがある

### 依存関係をgo.modに追加する
$ go get .

- go.sumファイルがpackage.lock.json的なやつ

### example.httpファイルの使い方
vscodeにREST ClientというExtensionsをインストールして使ってください