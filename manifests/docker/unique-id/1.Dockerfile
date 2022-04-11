# 通过 sh 命令进行构建, 工作路径为项目根目录
FROM alpine:latest
WORKDIR /tmp
COPY ./service_demo/unique-id/unique-id ./unique-id
COPY ./config/service_demo/unique_id/config_1.json ./config.json
COPY ./service_demo/unique-id/unique_id.json ./unique_id.json
COPY ./anomaly_injector/cpu.cpp ./cpu.cpp

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk add g++ && \
    g++ -fopenmp -o cpu /tmp/cpu.cpp && \
    apk del g++ && \
    apk add libgcc libstdc++ libgomp

ENTRYPOINT ["./unique-id"]
CMD ["-config_file","./config.json"]