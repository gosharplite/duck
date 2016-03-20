FROM ubuntu:14.04
ADD duck duck
EXPOSE 80
ENTRYPOINT ["/duck"]
