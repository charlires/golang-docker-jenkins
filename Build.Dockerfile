ARG BASE_TAG
ARG IMAGE_NAME
FROM ${IMAGE_NAME}-base:${BASE_TAG}

# Compile binary
ARG APP_NAME
RUN go install ${APP_NAME}/app

# Run program
ENTRYPOINT ["/go/bin/app"]

# Document that the service listens on port 80.
EXPOSE 80
