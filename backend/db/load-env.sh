#!/bin/bash

# config.yaml 파일의 경로
CONFIG_PATH="../configs/config.yaml"

# PostgreSQL 설정 추출
POSTGRES_USER=$(grep "db_user:" $CONFIG_PATH | awk -F "'" '{print $2}')
POSTGRES_PASSWORD=$(grep "db_password:" $CONFIG_PATH | awk -F "'" '{print $2}')
POSTGRES_DB=$(grep "db_name:" $CONFIG_PATH | awk -F "'" '{print $2}')
POSTGRES_PORT=$(grep "db_port:" $CONFIG_PATH | awk -F "'" '{print $2}')

# MongoDB 설정 추출
MONGO_URI=$(grep "mongo_uri:" $CONFIG_PATH | awk -F "'" '{print $2}')
MONGO_DB=$(grep "mongo_db:" $CONFIG_PATH | awk -F "'" '{print $2}')
# MongoDB URI에서 사용자 정보 추출
MONGO_USER=$(echo $MONGO_URI | sed -n 's/.*:\/\/\([^:]*\):.*/\1/p')
MONGO_PASSWORD=$(echo $MONGO_URI | sed -n 's/.*:\/\/[^:]*:\([^@]*\).*/\1/p')
MONGO_PORT=$(echo $MONGO_URI | sed -n 's/.*@[^:]*:\([0-9]*\)\/.*/\1/p')

# .env 파일 생성
cat > .env << EOF
# PostgreSQL 설정
POSTGRES_USER=$POSTGRES_USER
POSTGRES_PASSWORD=$POSTGRES_PASSWORD
POSTGRES_DB=$POSTGRES_DB
POSTGRES_PORT=$POSTGRES_PORT

# MongoDB 설정
MONGO_INITDB_ROOT_USERNAME=$MONGO_USER
MONGO_INITDB_ROOT_PASSWORD=$MONGO_PASSWORD
MONGO_INITDB_DATABASE=$MONGO_DB
MONGO_PORT=$MONGO_PORT
EOF

echo "환경 변수가 config.yaml에서 .env 파일로 성공적으로 추출되었습니다." 