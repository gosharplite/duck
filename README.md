# duck
Web server as sitting duck.

$ docker build -t gosharplite/duck:v6 .

$ docker push gosharplite/duck:v6

$ docker run --publish 8093:8092 gosharplite/duck:v6 -port=8092