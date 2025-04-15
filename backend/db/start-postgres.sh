#!/bin/bash

# config.yaml에서 환경 변수를 로드합니다
./load-env.sh

# 환경 변수를 적용하여 docker-compose를 실행합니다
docker-compose -f docker-compose-postgres.yml up -d

echo "PostgreSQL 컨테이너가 성공적으로 시작되었습니다." 