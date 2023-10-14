FROM golang:bookworm as builder
RUN apt-get install -y make git
COPY ./ /usr/src/app/
WORKDIR /usr/src/app
RUN make clean && make

FROM debian:bookworm-slim
ARG uid=1000
ARG gid=1000
ARG user=gouser
ARG userhome=/app
RUN groupadd -f -g ${gid} ${user}
RUN useradd -l -u ${uid} -g ${gid} -d "${userhome}" -r ${user}
USER ${user}
WORKDIR ${userhome}
COPY --from=builder /usr/src/app/gomrefdash .
COPY --from=builder /usr/src/app/frontend/spa/ frontend/spa/
EXPOSE 3000
ENTRYPOINT ./gomrefdash
