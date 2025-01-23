FROM uhub.service.ucloud.cn/entropypool/alpine:3.16

WORKDIR /app

COPY file-gateway /app

ENTRYPOINT ["/app/file-gateway"]

EXPOSE 60330 60331 60332 60333 60150 60151

CMD ["run"]
