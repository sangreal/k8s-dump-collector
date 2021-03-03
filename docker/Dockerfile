FROM openjdk:15-slim
WORKDIR /
RUN mkdir -p /coredumps
COPY build/libs/com.smartnews.dump.collector-0.1.0.jar /app.jar
CMD ["java", "-jar", "-Xmx10m", "-Xms1m", "-XX:+HeapDumpOnOutOfMemoryError", "-XX:HeapDumpPath=/coredumps", "app.jar"]