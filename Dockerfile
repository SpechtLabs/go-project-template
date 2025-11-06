FROM alpine:latest

LABEL org.opencontainers.image.title="PROJECT_NAME"
LABEL org.opencontainers.image.source="https://github.com/OWNER/PROJECT_NAME"
LABEL org.opencontainers.image.description="PROJECT_NAME application"
LABEL org.opencontainers.image.licenses="Apache 2.0"

COPY ./PROJECT_NAME-server /bin/PROJECT_NAME-server

ENTRYPOINT ["/bin/PROJECT_NAME-server"]
CMD [""]

EXPOSE 8080
