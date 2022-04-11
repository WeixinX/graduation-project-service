# 注意把 service-a 和 service-b build 到这个目录下
FROM alpine:latest
WORKDIR /tmp
COPY ./service-a ./service-a
COPY ./service-b ./service-b
COPY test_service_a_b.sh ./run.sh
COPY ./cpu.cpp ./cpu.cpp

# 先换源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk add g++ && \
    g++ -fopenmp -o cpu /tmp/cpu.cpp && \
    apk del g++ && \
    apk add libgcc libstdc++ libgomp

ENTRYPOINT ["sh","/tmp/run.sh"]
