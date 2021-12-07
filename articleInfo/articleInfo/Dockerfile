FROM golang:1.16-alpine as go-builder

# FROM golang:latest

# docker build -t afzal442/webapisl:v1 .
# docker run -p 5000:8080 afzal442/webapisl:v1 [host port:container port]t]

WORKDIR /app
COPY . /app/
ADD aws-lambda-rie /aws-lambda-rie
RUN mkdir -p /opt/extensions
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM scratch
COPY --from=go-builder /opt/extensions /opt/extensions
COPY --from=go-builder /aws-lambda-rie /aws-lambda-rie
COPY --from=go-builder /app/articleInfo /articleInfo
ENTRYPOINT ["/articleInfo"]