// src/store/user/getters.ts
import type { UserState } from './state';

export const getters = {
    // 사용자 인증 여부 확인
    isAuthenticated: (state: UserState): boolean => {
        return !!state.user;
    },

    // 사용자 이름 (없으면 'Guest' 반환)
    userNickname: (state: UserState): string => {
        return state.user?.nickname || 'Guest';
    },

    // 사용자 이메일 (없으면 빈 문자열 반환)
    userEmail: (state: UserState): string | null => {
        return state.user?.email || null;
    },

    // 사용자 ID (없으면 null 반환)
    userId: (state: UserState): number | null => {
        return state.user?.id || null;
    },

    // 사용자 프로필 이미지 (없으면 null 반환)
    userProfile: (state: UserState): string | null => {
        return state.user?.profile || null;
    },
};
