# duck
Web server as sitting duck.

$ CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' .

$ docker build -t gosharplite/duck:v13 .

$ docker push gosharplite/duck:v13

$ docker run --publish 8093:8092 gosharplite/duck:v13 -port=8092
