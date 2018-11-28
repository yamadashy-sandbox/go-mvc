キオクムシのapiサーバー

# 環境構築
必要なライブラリを読み込む
```
sudo go get -u github.com/gin-gonic/gin
```

gcloudでdepなどが使えないので、一旦グローバルにインストールする

# デバッグ
ローカルサーバー起動
```
dev_appserver.py gae/app.yaml
```

# デプロイ
```
$ cd gae
$ gcloud app deploy --version <バージョン番号>
```
