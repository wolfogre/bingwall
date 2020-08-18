NAME=bingwall
IMAGE_NAME=registry.aliyuncs.com/wolfogre/$(NAME)
VERSION=$(shell date "+%y.%m").$(shell git rev-list --count --since="$(shell date "+%Y-%m")-01T00:00:00+08:00" HEAD)

check:
	git diff HEAD --quiet || exit 1

build: check
	mkdir -p bin
	GOOS=linux GOARCH=amd64 go build -ldflags="-X $(NAME)/internal/version.injectVersion=$(VERSION)" -o bin/$(NAME)

image: build
	docker build -t $(IMAGE_NAME):$(VERSION) .

push:
	docker push $(IMAGE_NAME):$(VERSION)

clean:
	rm -rf bin

all: image push clean

