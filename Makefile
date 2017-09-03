build:
	go build
docker build: build
	docker build -t registry.cn-hangzhou.aliyuncs.com/wolfogre-hub/bingwall:$version .
docker push:
	docekr push registry.cn-hangzhou.aliyuncs.com/wolfogre-hub/bingwall:$version
clean:
	rm -f bingwall