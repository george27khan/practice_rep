для работы с кафка нужна librakafka
лучше ставить через wsl
```
sudo apt update
sudo apt install -y librdkafka-dev
rdkafka-config --version
```

для запуска программы на го нужен флаг
```
 CGO_ENABLED=1 go run main.go
```