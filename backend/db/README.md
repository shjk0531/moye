# 데이터베이스 설정

이 디렉토리에는 PostgreSQL 및 MongoDB를 Docker로 실행하기 위한 설정 파일이 포함되어 있습니다.

## 설정 방법

1. 모든 민감한 정보(사용자 이름, 비밀번호, 포트 등)는 `/backend/configs/config.yaml`에 정의되어 있습니다.
2. `load-env.sh` 스크립트는 config.yaml에서 환경 변수를 추출하여 `.env` 파일을 생성합니다.
3. 각 Docker Compose 파일은 이 환경 변수를 사용합니다.

## 사용 방법

### 모든 데이터베이스 실행

```bash
./start-all.sh
```

### PostgreSQL만 실행

```bash
./start-postgres.sh
```

### MongoDB만 실행

```bash
./start-mongo.sh
```

### 직접 환경 변수 생성 후 실행

```bash
./load-env.sh
docker-compose -f docker-compose-postgres.yml up -d
docker-compose -f docker-compose-mongo.yml up -d
```

## 주의사항

-   민감한 정보를 직접 Docker Compose 파일에 하드코딩하지 마세요.
-   `.env` 파일은 `.gitignore`에 추가하여 Git에 포함되지 않도록 해야 합니다.
