# duck
Web server as sitting duck.

$ docker build -t gosharplite/duck:v5 .

$ docker push gosharplite/duck:v5

$ docker run --publish 8093:8092 gosharplite/duck:v5 -port=8092