build:
	go build
dkbuild: build
	docker build -t registry.cn-hangzhou.aliyuncs.com/wolfogre-hub/bingwall:${version} .
dkpush:
	docekr push registry.cn-hangzhou.aliyuncs.com/wolfogre-hub/bingwall:${version}
clean:
	rm -f bingwall
