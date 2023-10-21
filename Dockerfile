FROM golang:1.20 as build
WORKDIR /app
COPY go.mod go.sum ./
COPY Makefile ./
COPY ./cmd ./cmd
COPY ./internal ./internal
RUN make build
FROM debian:12-slim
WORKDIR /root
COPY --from=build /app/.bin/app ./app
CMD ["./app"]
