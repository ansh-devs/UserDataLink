FROM golang:alpine as BuildStage

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main cmd/main.go

FROM alpine:latest
COPY --from=BuildStage /app/main .
EXPOSE 8080
CMD ["./main"]