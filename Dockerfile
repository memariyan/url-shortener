FROM golang:1.23 AS build

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir -p /src

WORKDIR /src

COPY . ./
RUN go build -o url-shortener

FROM debian:bullseye AS local

ARG BUILD_PATH

ENV TZ=Asia/Tehran \
    PATH="/app:${PATH}"

WORKDIR /app

COPY --from=build /src ./

EXPOSE 8001

CMD ["./url-shortener", "serve", "-p", "8001"]
