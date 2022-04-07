# 通过 sh 命令进行构建, 工作路径为项目根目录
FROM alpine:latest
WORKDIR /tmp
COPY ./service_load_balancer/lb-service ./lb-service
COPY ./config/service_load_balancer/lb_text.json ./config.json
ENTRYPOINT ["./lb-service"]
CMD ["-config_file","./config.json"]