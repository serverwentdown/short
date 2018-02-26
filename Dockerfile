FROM golang:1.10-alpine as go

WORKDIR /go/src/short
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -ldflags '-extldflags "-static" -X main.version=1.0' -o short


FROM scratch

EXPOSE 8080
COPY --from=go /go/src/short/short /short

HEALTHCHECK CMD ["/short", "-healthcheck"]

ENTRYPOINT ["/short"]
