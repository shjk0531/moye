import type { UserState } from './state';

export const getters = {
    // 사용자 인증 여부 확인
    isAuthenticated: (state: UserState): boolean => {
        return !!state.user;
    },

    // 사용자 이름 (없으면 'Guest' 반환)
    userName: (state: UserState): string => {
        return state.user?.name || 'Guest';
    },

    // 사용자 이메일 (없으면 빈 문자열 반환)
    userEmail: (state: UserState): string => {
        return state.user?.email || '';
    },

    // 사용자 ID (없으면 null 반환)
    userId: (state: UserState): number | null => {
        return state.user?.id || null;
    },
};
