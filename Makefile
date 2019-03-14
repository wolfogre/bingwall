NAME=bingwall
IMAGE_NAME=registry.aliyuncs.com/wolfogre/$(NAME)
UNIX_TIME=$(shell git log -1 --format="%at")
MAIN_VERSION=$(shell grep -E "[0-9][0-9]*\.[0-9][0-9]*" internal/version/version.go -o)

check:
	git diff HEAD --quiet || exit 1

build: check
	mkdir -p bin
	GOOS=linux GOARCH=amd64 go build -ldflags="-X $(NAME)/internal/version.timeVersion=$(UNIX_TIME)" -o bin/$(NAME)

image: build
	docker build -t $(IMAGE_NAME):$(MAIN_VERSION).$(UNIX_TIME) .

push:
	docker push $(IMAGE_NAME):$(MAIN_VERSION).$(UNIX_TIME)

clean:
	rm -rf bin

all: image push clean