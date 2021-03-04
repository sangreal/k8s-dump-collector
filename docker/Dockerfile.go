FROM openjdk:15-slim
WORKDIR /
COPY output/dump-tool /dump-tool
ENTRYPOINT ["dump-tool"]