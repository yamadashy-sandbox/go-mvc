# キオクムシのAPIサーバー(?)

## フォルダ構成
Laravelっぽい構成

- app
    - database
        - models
    - http
        - controllers
        - middlewares
        - structs
- gae ... GAE関連
- public
- resources 
    - views ... テンプレート関連
- routes ... ルーティング関連
- tests

## 開発

### 環境構築
必要なライブラリを読み込む
```
sudo go get -u github.com/gin-gonic/gin
```

gcloudでdepなどが使えないので、一旦グローバルにインストールする

### デバッグ
ローカルサーバー起動
```
dev_appserver.py gae/app.yaml
```

### デプロイ
```
$ cd gae
$ gcloud app deploy --version <バージョン番号>
```
