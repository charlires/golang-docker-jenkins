ARG BASE_TAG
ARG IMAGE_NAME
FROM ${IMAGE_NAME}-base:${BASE_TAG}

RUN go get -u github.com/jstemmer/go-junit-report

CMD go test ./... -v 2>&1 | go-junit-report