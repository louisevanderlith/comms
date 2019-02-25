FROM alpine:latest

COPY comms .
COPY conf conf

EXPOSE 8085

ENTRYPOINT [ "./comms" ]