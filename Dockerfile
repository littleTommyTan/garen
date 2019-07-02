FROM debian:latest
RUN mkdir -p /app
WORKDIR /app

ADD ./dist/garen /app/garen

CMD ["./garen"]