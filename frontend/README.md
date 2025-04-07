# Moye Frontend

## 프로젝트 구조

이 프로젝트는 Feature-Sliced Design(FSD) 아키텍처 패턴을 사용하여 구성되었습니다.

### 레이어 구조

```
src/
├── app/          # 애플리케이션 진입점, 설정
├── processes/    # 비즈니스 프로세스 (로그인, 인증 등)
├── pages/        # 페이지 컴포넌트
├── widgets/      # 복잡한 독립 블록 (사이드바, 알림 등)
├── features/     # 사용자 상호작용 기능 (인증, 채팅 등)
├── entities/     # 비즈니스 엔티티 (사용자, 채널 등)
├── shared/       # 공유 유틸리티, UI 컴포넌트, API 클라이언트 등
└── store/        # 상태 관리
```

### 레이어 별 책임

-   **app**: 애플리케이션 초기화, 글로벌 스타일, 프로바이더 설정
-   **processes**: 복잡한 비즈니스 프로세스 처리
-   **pages**: 라우팅과 레이아웃
-   **widgets**: 페이지를 구성하는 독립적인 블록
-   **features**: 사용자 기능, 인터랙션
-   **entities**: 비즈니스 모델, 타입 정의
-   **shared**: 공통 유틸리티, UI 라이브러리, API 클라이언트

## 개발 가이드

### 레이어 간 의존성 규칙

1. 각 레이어는 자신보다 아래 레벨의 레이어만 임포트할 수 있습니다.
2. 상위 레이어는 하위 레이어에 의존할 수 없습니다.
3. 같은 레벨의 슬라이스 간에는 직접 의존성이 없어야 합니다.

```
app → processes → pages → widgets → features → entities → shared
```

### 컴포넌트 네이밍 컨벤션

-   컴포넌트 이름은 파스칼 케이스(PascalCase)로 작성합니다.
-   각 레이어의 모듈은 케밥 케이스(kebab-case)로 작성합니다.

## 기술 스택

-   Vue 3
-   TypeScript
-   Vuex
-   Vue Router
-   PrimeVue (UI 컴포넌트 라이브러리)
-   Axios
