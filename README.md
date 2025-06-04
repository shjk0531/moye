# Moye

## 프로젝트 개요

본 프로젝트는 실시간 채팅과 커뮤니티 기능을 중심으로 한 라운지 웹 플랫폼입니다. 기존 팀 프로젝트에서 발견한 문제점을 개선하고자 개인 프로젝트로 발전시켰으며, 실시간 텍스트 및 영상 채팅, 사용자 맞춤형 검색 기능을 제공합니다. 사용자는 웹사이트를 통해 자유롭게 그룹을 형성하고 다양한 주제로 소통할 수 있습니다.

---

## 기술 스택

### Frontend

-   **Vue 3**: 사용자 인터페이스 구축 및 관리 (Composition API)
-   **PrimeVue 4 (unstyled mode)**: UI 컴포넌트 라이브러리
-   **Tailwind CSS v4**: 스타일링 및 반응형 UI 레이아웃
-   **Vite**: 빠른 빌드 및 개발 환경 구성
-   **Pinia**: 간편한 상태 관리

### Backend

-   **Go (Golang)**: 고성능 및 경량 백엔드 서버 개발
-   **Gin Framework**: 웹 서버 및 REST API 라우팅
-   **GORM**: PostgreSQL ORM
-   **mgm**: MongoDB 데이터 관리
-   **Redis**: 실시간 채팅 및 온라인 사용자 상태 관리
-   **WebSocket**: 실시간 양방향 통신 지원
-   **Inverted Index (Go)**: Elasticsearch 대체 경량 검색 엔진 구현

### Database

-   **PostgreSQL 17**: 주 데이터 관리 및 FTS(GIN 인덱스)를 이용한 빠른 검색
-   **MongoDB**: 채팅 로그 등 비정형 데이터 저장
-   **Redis**: 빠른 접근 및 실시간 데이터 처리 인메모리 DB

### 인프라

-   **AWS**: 클라우드 인프라 및 서비스 관리
-   **AWS S3**: 이미지 및 파일 스토리지
-   **Docker**: 컨테이너화 및 배포 효율화
-   **Nginx**: 리버스 프록시 및 로드밸런싱
-   **GitHub Actions**: CI/CD를 통한 자동화된 빌드, 테스트 및 배포

---

## 프로젝트 주요 기능

### 사용자 인증

-   JWT 기반의 인증 및 인가
-   안전한 사용자 세션 관리 및 보안 유지

### 커뮤니티 및 채팅 기능

-   디스코드 유사 서버 및 채팅룸 구성 기능
-   실시간 텍스트 및 비디오 채팅
-   Redis 기반 온라인 사용자 실시간 관리

### 고성능 맞춤형 검색 기능

-   Elasticsearch 대신 Go 기반 경량 Inverted Index 직접 구현
-   PostgreSQL Full-Text Search(GIN) 연동
-   저비용·고성능의 빠른 검색 응답(수십 밀리초)

### 데이터 관리

-   민감 정보 보호를 위한 UUID 적용
-   이미지 및 파일을 AWS S3에 안전 저장

---

## 아키텍처

![image](https://github.com/user-attachments/assets/881db6ae-3bd7-4b12-a8a9-59ed2b0bf5f8)

---

## 배포 전략

-   GitHub Actions를 통한 자동화된 빌드 및 테스트, 무중단 배포
-   Docker 컨테이너화로 효율적 관리
-   AWS 및 Nginx 리버스 프록시로 안정적 서비스 운영 및 로드밸런싱

## 환경 설정

`backend/scripts/setup-config.sh` 스크립트를 실행하면 환경 변수 값을 사용해
`backend/configs/config.yaml` 파일을 생성할 수 있습니다.

```bash
cd backend/scripts
./setup-config.sh
```

생성된 `config.yaml` 파일은 `.gitignore`에 포함되어 있어 저장소에 커밋되지 않습니다.

---

## 기술 스택 전환 배경

기존 프로젝트(Spring Boot + React)에서 발견된 한계와 문제점을 극복하기 위해 Go와 Vue 3를 선택했습니다.

### Backend: Go vs. Spring Boot

-   **경량성**: 메모리 사용량 및 컨테이너 크기 대폭 감소
-   **성능**: Goroutine 기반의 효율적인 실시간 처리
-   **운영 단순화**: JVM 튜닝 없이 간편한 서비스 관리

### Frontend: Vue 3 vs. React

-   **개발 생산성**: Composition API를 통한 명확한 로직 및 UI 관리
-   **상태 관리 간소화**: Pinia 기본 지원
-   **빠른 개발 환경**: Vite를 통한 빠른 빌드 및 Hot Module Replacement

### 기대 효과

-   **비용 절감**: 인프라 운영 비용 약 50\~60% 감소
-   **성능 향상**: 빠른 응답 속도로 사용자 경험 개선
-   **유지보수 간편화**: 서비스 운영 및 확장 용이

---

## ELK 기반 검색 아키텍처 개선

### 이전 문제점

-   **높은 운영 비용**: ELK 스택 유지 비용 과다
-   **복잡한 관리**: JVM 튜닝, 장애 대응 복잡성
-   **과도한 기능**: 로그 및 모니터링 불필요

### 요구사항

-   경량 서비스화 및 저비용 운영
-   높은 검색 성능 유지
-   운영 복잡도 감소

### 개선 방안

-   Go 기반 자체 Inverted Index 구현 및 PostgreSQL FTS(GIN 인덱스) 사용
-   Elasticsearch 수준의 빠른 검색 응답
-   비용 효율적이고 관리하기 쉬운 솔루션 제공

### 기대 효과

-   **운영 비용 절감**: 인프라 비용 대폭 감소
-   **유지보수 효율화**: 관리 대상 서비스 단순화
-   **성능 확보**: 검색 응답 시간 최적화

---

## 향후 개발 계획

-   세부적인 커뮤니티 및 관리자 기능 추가
-   점진적 마이크로서비스 구조 확장
-   지속적인 성능 최적화 및 안정성 강화
-   추가적인 사용자 피드백 반영 및 기능 개선
