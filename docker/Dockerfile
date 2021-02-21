FROM openjdk:11-slim
WORKDIR /
COPY com.smartnews.dump.collector-0.1.0.jar /app.jar
CMD ["java", "-jar", "-Xmx10m", "-Xms1m", "-XX:+HeapDumpOnOutOfMemoryError", "-XX:HeapDumpPath=/tmp/dumps", "app.jar"]