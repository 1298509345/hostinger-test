FROM golang:1.23.4

COPY app .
COPY conf/sim.toml ./conf/

ENTRYPOINT ["./app"]

