FROM alpine:latest

LABEL org.opencontainers.image.title="{{REPO_NAME}}"
LABEL org.opencontainers.image.source="https://github.com/SpechtLabs/{{REPO_NAME}}"
LABEL org.opencontainers.image.description="{{REPO_DESCRIPTION}}"
LABEL org.opencontainers.image.licenses="Apache 2.0"

COPY ./{{REPO_NAME}} /bin/{{REPO_NAME}}

ENTRYPOINT ["/bin/{{REPO_NAME}}"]
CMD [""]

EXPOSE 8099
EXPOSE 50051

