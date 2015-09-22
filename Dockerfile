FROM google/debian:wheezy
ADD duck duck
EXPOSE 80
ENTRYPOINT ["/duck"]