# 通过 sh 命令进行构建, 工作路径为项目根目录
FROM alpine:latest
WORKDIR /tmp
COPY ./service_demo/media/media ./media
COPY ./config/service_demo/media/config_2.json ./config.json
ENTRYPOINT ["./media"]
CMD ["-config_file","./config.json"]