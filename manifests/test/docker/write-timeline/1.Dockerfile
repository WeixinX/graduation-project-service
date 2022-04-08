# 通过 sh 命令进行构建, 工作路径为项目根目录
FROM alpine:latest
WORKDIR /tmp
COPY ./service_demo/write-timeline/write-timeline ./write-timeline
COPY ./config/test/service_demo/write_timeline/config_1.json ./config.json
ENTRYPOINT ["./write-timeline"]
CMD ["-config_file","./config.json"]