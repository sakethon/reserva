#!/bin/sh

until (/usr/bin/mc config host add minio http://minio:9000 ${MINIO_ACCESS_KEY} ${MINIO_SECRET_KEY}) do sleep 1; done;
/usr/bin/mc mb minio/${MINIO_BUCKET_IN}
/usr/bin/mc mb minio/${MINIO_BUCKET_OUT}

if [ -e /data_init ]; then
    /usr/bin/mc cp -r /data_init/ minio/${MINIO_BUCKET_IN}/
    /usr/bin/mc cp -r /data_init/ minio/${MINIO_BUCKET_OUT}/
fi
