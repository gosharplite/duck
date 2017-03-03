FROM scratch
ADD duck /
EXPOSE 8080
ENTRYPOINT ["/duck"]
