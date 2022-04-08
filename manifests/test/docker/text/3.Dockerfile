# 通过 sh 命令进行构建, 工作路径为项目根目录
FROM alpine:latest
WORKDIR /tmp
COPY ./service_demo/text/text ./text
COPY ./config/test/service_demo/text/config_3.json ./config.json
ENTRYPOINT ["./text"]
CMD ["-config_file","./config.json"]