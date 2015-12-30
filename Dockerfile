FROM ubuntu:latest
RUN apt-get update && apt-get install -y \
    vim \
    curl \
	dnsutils
ADD duck duck
EXPOSE 80
ENTRYPOINT ["/duck"]