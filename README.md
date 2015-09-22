# duck
Web server as sitting duck.

$ docker build -t gosharplite/duck:v4 .

$ docker push gosharplite/duck:v4

$ docker run --publish 8093:8092 gosharplite/duck:v4 -port=8092