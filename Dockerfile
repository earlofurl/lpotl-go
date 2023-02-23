## Build stage
#FROM golang:1.19-alpine3.16 AS builder
#WORKDIR /app
#COPY . .
#RUN go build -o main main.go
## RUN apk add curl
## RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
#
## Run stage
#FROM alpine:3.16
#WORKDIR /app
#COPY --from=builder /app/main .
## COPY --from=builder /app/migrate.linux-amd64 ./migrate
## COPY app.env .
##COPY --from=builder /app/start.sh .
#COPY start.sh .
#RUN chmod +x start.sh
## RUN chmod +x /app/start.sh
#COPY db/migration ./db/migration
#
#EXPOSE 8080
#CMD [ "/app/main" ]
#ENTRYPOINT [ "start.sh" ]

##
## Build Stage
##
FROM golang:1.19-alpine3.16 AS builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY cmd/lpotl-go/main.go ./
COPY . ./
COPY app.env ./
COPY postgres/migration ./postgres/migration

RUN go build -o /lptl-go

##
## Deploy Stage
##

FROM alpine:3.16

WORKDIR /

# Copy the binary from the build stage
COPY --from=builder /lpotl-go /lpotl-go

# Copy the app.env file
COPY --from=builder /app/app.env ./

EXPOSE 8080

# Set non root user
#RUN adduser -D -g '' pixelthc
#USER pixelthc:pixelthc

ENTRYPOINT [ "/lpotl-go" ]
