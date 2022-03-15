#!/bin/bash
# if [ ! -n "$TEAMGRAM_HOST" ]; then
#   echo ">>> Plase set environment variable TEAMGRAM_HOST to your own server IP. <<<"
#   exit 1
# fi

export TEAMGRAM_HOST=${TEAMGRAM_HOST:-"0.0.0.0"}
export ETCD_URL=${ETCD_URL:-"etcd:2379"}
export REDIS_HOST=${REDIS_HOST:-"redis:6379"}
export KAFKA_HOST=${REDIS_HOST:-"broker:9092"}
export MYSQL_URI=${MYSQL_URI:-"teamgram:teamgram@tcp(mysql:3306)/teamgram?charset=utf8mb4"}
export MINIO_URI=${MYSQL_URI:-"minio:9000"}
export MINIO_KEY=${MYSQL_URI:-"minio"}
export MINIO_SECRET=${MYSQL_URI:-"miniostorage"}
export MINIO_SSL=${MYSQL_URI:-"false"}

# create configs from config templates.
createConfigs() {
  CONFIG_TARGET_DIR=/app/etc
  CONFIG_TEMPLATES_DIR=/app/etc
  for file in `ls $CONFIG_TEMPLATES_DIR`; do
    cat $CONFIG_TEMPLATES_DIR/$file \
      | sed 's#"ListenOn: 127.0.0.1#ListenOn: $TEAMGRAM_HOST"#g' \
      | sed "s#127.0.0.1:2379#$ETCD_URL#g" \
      | sed "s#127.0.0.1:6379#$REDIS_HOST#g" \
      | sed "s#root:@tcp(127.0.0.1:3306)/teamgram?charset=utf8mb4#$MYSQL_URI#g" \
      | sed "s#AccessKeyID: minio#AccessKeyID: $MINIO_KEY#g" \
      | sed "s#SecretAccessKey: miniostorage#SecretAccessKey: $MINIO_SECRET#g" \
      | sed "s#UseSSL: false#UseSSL: $MINIO_SSL#g" \
      | sed "s#localhost:9000#$MINIO_URI#g" \
      | sed "s#127.0.0.1:9092#$KAFKA_HOST#g" \
      | cat > $CONFIG_TARGET_DIR/$file
  done
}

runService() {
    echo "starting $1..."
    ./$1 >> /tmp/teamgram-$1.log &
    echo "running $1!"
}

createConfigs

runService idgen
runService status
runService authsession
runService dfs
runService media
runService biz
runService msg
runService sync
runService bff
runService session
runService gateway


tail -f /tmp/teamgram-*.log

