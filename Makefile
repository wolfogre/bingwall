version=$(shell git tag -l | tail -n 1`)
name=bingwall

build:
	mkdir -p bin
	GOOS=linux GOARCH=amd64 go build -o $(name)
image: build
	docker build -t reg.qiniu.com/wolfogre/$(name):$(version) .
push:
	docker push reg.qiniu.com/wolfogre/$(name):$(version)
clean:
	rm -rf bin
