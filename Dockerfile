FROM scratch

COPY cmd/cmd .
COPY templates templates

EXPOSE 8085

ENTRYPOINT [ "./cmd" ]