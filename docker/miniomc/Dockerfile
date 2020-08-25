FROM minio/mc:latest
ENV TZ Asia/Tokyo
COPY ./run.sh /
RUN chmod 744 /run.sh
ENTRYPOINT ["/run.sh"]
