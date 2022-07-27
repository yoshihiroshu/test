# 環境構築

依存モジュールのインストール

```
go mod download
```

imageをビルドする

```
docker-compose up --build
```

docker-networkの作成

```
docker network create test-backend-network
```

バックグラウンドで立ち上げる

```
docker-compose up -d
```

確認

```
curl localhost:80/
```
