# 注意把 service-a 和 service-b build 到这个目录下
FROM alpine:latest
WORKDIR /tmp
COPY ./service-a ./service-a
COPY ./service-b ./service-b
COPY ./test_service_a_b.sh ./run.sh
ENTRYPOINT ["sh","/tmp/run.sh"]
