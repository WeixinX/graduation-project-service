# 通过 sh 命令进行构建, 工作路径为项目根目录
FROM alpine:latest
WORKDIR /tmp
COPY ./service_demo/compose-post/compose-post ./compose-post
COPY ./config/test/service_demo/compose_post/config_3.json ./config.json
ENTRYPOINT ["./compose-post"]
CMD ["-config_file","./config.json"]