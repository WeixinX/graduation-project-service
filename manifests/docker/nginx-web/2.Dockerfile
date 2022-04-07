# 通过 sh 命令进行构建, 工作路径为项目根目录
FROM alpine:latest
WORKDIR /tmp
COPY ./service_demo/nginx-web/nginx-web ./nginx-web
COPY ./config/service_demo/nginx_web/config_2.json ./config.json
ENTRYPOINT ["./nginx-web"]
CMD ["-config_file","./config.json"]