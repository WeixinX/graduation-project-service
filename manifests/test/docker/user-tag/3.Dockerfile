# 通过 sh 命令进行构建, 工作路径为项目根目录
FROM alpine:latest
WORKDIR /tmp
COPY ./service_demo/user-tag/user-tag ./user-tag
COPY ./config/test/service_demo/user_tag/config_3.json ./config.json
COPY ./service_demo/user-tag/user_tag.json ./user_tag.json
ENTRYPOINT ["./user-tag"]
CMD ["-config_file","./config.json"]