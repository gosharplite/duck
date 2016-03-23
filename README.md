# duck
Web server as sitting duck.

$ docker build -t gosharplite/duck:v11 .

$ docker push gosharplite/duck:v11

$ docker run --publish 8093:8092 gosharplite/duck:v11 -port=8092
