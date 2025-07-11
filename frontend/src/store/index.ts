// src/store/index.ts
import { createPinia } from 'pinia';
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate';

// 각 모듈 스토어 내보내기
export { useAppStore } from './app';
export { useUserStore } from './user';
export { useLoungeStore } from './lounge';
export { useNoticeStore } from './notice';
export { useChatStore } from './chat';

// Pinia 인스턴스 생성
const pinia = createPinia();
pinia.use(piniaPluginPersistedstate);

export default pinia;
