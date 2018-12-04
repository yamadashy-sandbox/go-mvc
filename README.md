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

### クローン
```
$ go get -u github.com/stivan622/kiokumushi-api
```

### 環境構築
Dockerのインストール
```
$ brew install docker
$ brew cask install docker
```

dep初期化
```
$ docker-compose run --rm dep init
```

gcloudでdepなどが使えないので、一旦グローバルにインストールする

### デバッグ
ローカルサーバー起動
```
$ docker-compose up app
```

### デプロイ
```
$ gcloud app deploy --version <バージョン番号>
```
