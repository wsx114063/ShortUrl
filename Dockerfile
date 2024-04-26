FROM golang:1.22.2-bullseye

LABEL author="kuro"

COPY . /startgo

ENV APIKey testApiKey

#RUN groupadd -r group1 && useradd -r -g group1 user1
RUN ["/bin/bash", "-c", "echo hello"]

#ENTRYPOINT ["docker-entrypoint.sh"]

WORKDIR /startgo
# VOLUME 建立儲存空間
VOLUME ["/data/app"]

# USER 建立使用者
#USER user1

EXPOSE 8080
# CMD 執行命令
CMD ["go" ,"run", "."]

# ONBUILD 用來給其他IMAGE當來源時，會執行的指令