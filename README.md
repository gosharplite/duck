# duck
Web server as sitting duck.

$ docker build -t gosharplite/duck:v1 .

$ docker push gosharplite/duck:v1

$ docker run --publish 8093:8092 gosharplite/duck:v1 -port=8092