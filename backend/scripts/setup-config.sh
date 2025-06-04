#!/bin/bash

# setup-config.sh: 환경 설정 파일(config.yaml)을 생성합니다.

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
CONFIG_DIR="$SCRIPT_DIR/../configs"
TARGET_FILE="$CONFIG_DIR/config.yaml"

# 환경 변수에서 값 읽기(기본값 제공)
SERVER_PORT="${SERVER_PORT:-:8080}"
DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-5432}"
DB_USER="${DB_USER:-postgres}"
DB_PASSWORD="${DB_PASSWORD:-postgres}"
DB_NAME="${DB_NAME:-moye}"
DB_SSLMODE="${DB_SSLMODE:-disable}"
MONGO_URI="${MONGO_URI:-mongodb://mongo_user:password@localhost:27017/moye?authSource=mongo_user}"
MONGO_DB="${MONGO_DB:-moye}"
JWT_ACCESS_TOKEN_SECRET="${JWT_ACCESS_TOKEN_SECRET:-your_access_token_secret}"
JWT_REFRESH_TOKEN_SECRET="${JWT_REFRESH_TOKEN_SECRET:-your_refresh_token_secret}"
JWT_ACCESS_DURATION="${JWT_ACCESS_DURATION:-15m}"
JWT_REFRESH_DURATION="${JWT_REFRESH_DURATION:-168h}"

mkdir -p "$CONFIG_DIR"

cat > "$TARGET_FILE" <<EOCONF
server_port: '${SERVER_PORT}'

db_host: '${DB_HOST}'
db_port: '${DB_PORT}'
db_user: '${DB_USER}'
db_password: '${DB_PASSWORD}'
db_name: '${DB_NAME}'
db_sslmode: '${DB_SSLMODE}'

mongo_uri: '${MONGO_URI}'
mongo_db: '${MONGO_DB}'

jwt_access_token_secret: '${JWT_ACCESS_TOKEN_SECRET}'
jwt_refresh_token_secret: '${JWT_REFRESH_TOKEN_SECRET}'
jwt_access_duration: '${JWT_ACCESS_DURATION}'
jwt_refresh_duration: '${JWT_REFRESH_DURATION}'
EOCONF

echo "생성된 설정 파일: $TARGET_FILE"
