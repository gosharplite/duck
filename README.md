# duck
Web server as sitting duck.

$ docker build -t gosharplite/duck:v9 .

$ docker push gosharplite/duck:v9

$ docker run --publish 8093:8092 gosharplite/duck:v9 -port=8092
