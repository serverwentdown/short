FROM golang:1.9-alpine as go

WORKDIR /go/src/short
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -ldflags '-extldflags "-static"' -o short


FROM scratch

EXPOSE 8080
COPY --from=go /go/src/short/short /short

ENTRYPOINT ["/short"]
