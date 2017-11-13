# Sample Golang app for CI
This sample app is made to help other developers start working with docker and jenkins

Its a work in progress so any comment and suggestion is well received

## Build base image
```
make build-base
```

## Build test image
```
make build-test
```

## Run tests (make build-test first)
```
make test-units
```

## Run project
``` bash
make run
```
default port is 8080
you can change the port using `PORT` argument
> make run PORT=8081

<!-- ## Deploy to kubernetes
```
kubectl create -f webapp.yml
```

## Deploy to kubernetes service
```
kubectl create -f webapp-service.yml
``` -->

<!-- ## Describe service
```
kubectl describe svc golang-app-service
``` -->