# Sample Docker files

## Build dev image
`docker build --rm -f DockerfileDev -t image_name:latest .`

## Run dev image
`docker run --rm -it -p 8080:8080 -v `\`pwd\``:/go/src/webapp image_name:latest`