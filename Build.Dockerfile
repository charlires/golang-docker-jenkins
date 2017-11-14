ARG BASE_TAG
ARG IMAGE_NAME
FROM ${IMAGE_NAME}-base:${BASE_TAG}

# Compile binary
ARG APP_NAME
RUN go install ${APP_NAME}

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 \
  CMD curl -f http://localhost/healthz || exit 1

# Run program
ARG APP_NAME
ENTRYPOINT ["/go/bin/${APP_NAME}"]

# Document that the service listens on port 80.
EXPOSE 80
