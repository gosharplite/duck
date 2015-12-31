# duck
Web server as sitting duck.

$ docker build -t gosharplite/duck:v7 .

$ docker push gosharplite/duck:v7

$ docker run --publish 8093:8092 gosharplite/duck:v7 -port=8092