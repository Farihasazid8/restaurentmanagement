FROM golang:1.15-alpine
#FROM golang:latest as builder
#RUN apt-get update && apt-get install -y nocache git ca-certificates && update-ca-certificates
#WORKDIR /app
#/github.com/Farihasazid8/restaurentManagement
#WORKDIR $GOPATH/home/fariha/go/src/restaurentmanagement
WORKDIR /src/restaurentmanagement
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
#RUN go get -d -v ./...
#RUN go install -v ./...
#RUN go env -w GOPROXY="https://goproxy.io,direct"
#
#ENV SERVICE_PORT=8080
#ENV MONGO_SERVER=localhost
#ENV MONGO_PORT=27017
#ENV MONGO_AUTH_DATABASE=restaurent_management
#ENV MONGO_USERNAME=mongoAdmin
#ENV MONGO_PASSWORD=abc123
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/restaurantmanagement .
RUN chmod +x /app/bin/restaurantmanagement
EXPOSE 8078
CMD ["/bin/sh", "-c", "/app/bin/restaurantmanagement"]