FROM golang:1.16-alpine as builder

# FROM golang:latest

# docker run -p 5000:8080 web4api

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o usrmgmt4api .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/usrmgmt4api .
COPY --from=builder /app/.env .    

EXPOSE 8080

CMD [ "./usrmgmt4api" ] 