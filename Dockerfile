FROM golang:alpine as builder
RUN apk add make git
COPY ./ /usr/src/app/
WORKDIR /usr/src/app
RUN make clean && make

FROM alpine:latest
ARG uid=1000
ARG gid=1000
RUN apk add file
RUN adduser -u ${uid} -g ${gid} -h /app -D gouser
USER gouser
WORKDIR /app
COPY --from=builder /usr/src/app/gomrefdash .
COPY --from=builder /usr/src/app/frontend/spa/ frontend/spa/
EXPOSE 3000
ENTRYPOINT ./gomrefdash