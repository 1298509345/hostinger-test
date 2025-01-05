FROM golang:1.23.4

COPY fab-dev .
COPY conf/sim.toml ./conf/

ENTRYPOINT ["./fab-dev"]

ENV TZ=Asia/Shanghai