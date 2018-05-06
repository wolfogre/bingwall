version=`git tag -l | tail -n 1`
name=bingwall

build:
	go build
dkbuild: build
	docker build -t reg.qiniu.com/wolfogre/${name}:${version} .
dkpush:
	docker push reg.qiniu.com/wolfogre/${name}:${version}
clean:
	rm -f ${name}