FROM golang:1.11 as builder

WORKDIR /box
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY main.go .
COPY controllers ./controllers
COPY core ./core
COPY routers ./routers

RUN CGO_ENABLED="0" go build

FROM alpine:latest

COPY --from=builder /box/comms .
COPY conf conf

EXPOSE 8085

ENTRYPOINT [ "./comms" ]