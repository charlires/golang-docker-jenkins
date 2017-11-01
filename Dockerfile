FROM golang:1.9.1

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64 && chmod +x /usr/local/bin/dep

# install the dependencies without checking for go code# Start from a Debian image with the latest version of Go installed
# Copy the local package files to the container's workspace.
WORKDIR /go/src/webapp
ADD main.go ./

ADD Gopkg.toml Gopkg.lock ./
#copies the Gopkg.toml and Gopkg.lock to WORKDIR

# Install dependencies
RUN dep ensure -vendor-only

# Compile binary
RUN go install webapp

# Run program
ENTRYPOINT ["/go/bin/webapp"]

# Document that the service listens on port 8080.
EXPOSE 8080
