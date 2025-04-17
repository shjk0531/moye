// src/store/index.ts
import { createPinia } from 'pinia';

// 각 모듈 스토어 내보내기
export { useAppStore } from './app';
export { useUserStore } from './user';
export { useStudyStore } from './study';

// Pinia 인스턴스 생성
const pinia = createPinia();

export default pinia;
