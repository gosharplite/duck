# duck
Web server as sitting duck.

$ docker build -t gosharplite/duck:v12 .

$ docker push gosharplite/duck:v12

$ docker run --publish 8093:8092 gosharplite/duck:v12 -port=8092
