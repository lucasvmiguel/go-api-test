##
## BUILD
## 
FROM golang:alpine AS build

WORKDIR /app

COPY . .
RUN go mod download

RUN go build /app/cmd/api/main.go

##
## RUN
##
FROM alpine

WORKDIR /

COPY --from=build /app/main /main

ENTRYPOINT ["/main"]