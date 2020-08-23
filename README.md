# reserva

## 環境構築

```sh
git clone git@github.com:rymiyamoto/reserva.git
docker-compose up -d
```

## 実行方法

```sh
# コンテナ起動(すでに立ち上がっていたら不要)
docker-compose up -d
```

### コンテナ外で実行する場合

```sh
docker-compose exec app go run main.go
```

### コンテナ内で実行する場合

```sh
docker-compose exec app sh
go run main.go
```

### linter実行

```sh
docker-compose exec app golangci-lint run
```
