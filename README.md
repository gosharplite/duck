# duck
Web server as sitting duck.

$ docker build -t gosharplite/duck:v8 .

$ docker push gosharplite/duck:v8

$ docker run --publish 8093:8092 gosharplite/duck:v8 -port=8092
