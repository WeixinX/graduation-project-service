# 通过 sh 命令进行构建, 工作路径为项目根目录
FROM alpine:latest
WORKDIR /tmp
COPY ./service_demo/compose-post/compose-post ./compose-post
COPY ./config/service_demo/compose_post/config_3.json ./config.json
COPY ./anomaly_injector/cpu.cpp ./cpu.cpp

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk add g++ && \
    g++ -fopenmp -o cpu /tmp/cpu.cpp && \
    apk del g++ && \
    apk add libgcc libstdc++ libgomp \

ENTRYPOINT ["./compose-post"]
CMD ["-config_file","./config.json"]