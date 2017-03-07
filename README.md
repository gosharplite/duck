# duck
Web server as sitting duck.

$ CGO_ENABLED=0 GOOS=linux go build -o duck -a -tags netgo -ldflags '-w' src/cmd/main.go

$ docker build -t gosharplite/duck:v15 .

$ docker push gosharplite/duck:v15

$ docker run --publish 8093:8092 gosharplite/duck:v15 -port=8092
