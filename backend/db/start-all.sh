#!/bin/bash

# 스크립트 실행 디렉토리로 이동
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
cd "$SCRIPT_DIR"

# config.yaml에서 환경 변수를 로드합니다
bash ./load-env.sh

# PostgreSQL 컨테이너 시작
docker-compose -f ./docker-compose-postgres.yml up -d
echo "PostgreSQL 컨테이너가 시작되었습니다."

# MongoDB 컨테이너 시작
docker-compose -f ./docker-compose-mongo.yml up -d
echo "MongoDB 컨테이너가 시작되었습니다."

echo "모든 데이터베이스 컨테이너가 성공적으로 시작되었습니다." 