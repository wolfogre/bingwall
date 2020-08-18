FROM centos:8

ENV TZ=Asia/Shanghai

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY bin/bingwall /usr/local/bin/bingwall

EXPOSE 80

CMD ["bingwall"]
