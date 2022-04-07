# 通过 sh 命令进行构建, 工作路径为项目根目录
FROM alpine:latest
WORKDIR /tmp
COPY ./service_demo/unique-id/unique-id ./unique-id
COPY ./config/service_demo/unique_id/config_1.json ./config.json
COPY ./service_demo/unique-id/unique_id.json ./unique_id.json
ENTRYPOINT ["./unique-id"]
CMD ["-config_file","./config.json"]