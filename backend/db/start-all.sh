#!/bin/bash

# config.yaml에서 환경 변수를 로드합니다
./load-env.sh

# PostgreSQL 컨테이너 시작
docker-compose -f docker-compose-postgres.yml up -d
echo "PostgreSQL 컨테이너가 시작되었습니다."

# MongoDB 컨테이너 시작
docker-compose -f docker-compose-mongo.yml up -d
echo "MongoDB 컨테이너가 시작되었습니다."

echo "모든 데이터베이스 컨테이너가 성공적으로 시작되었습니다." 