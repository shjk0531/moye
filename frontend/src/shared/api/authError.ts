import { useUserStore } from '@/store';
import router from '@/router';

export interface AuthErrorResponse {
    error: string;
    code?: string;
}

/**
 * 인증 오류를 처리하는 유틸리티 함수
 * 토큰 만료나 인증 실패 시 사용자를 로그아웃 처리하고 로그인 페이지로 리다이렉트
 */
export function handleAuthError(error: any): void {
    const userStore = useUserStore();

    // 이미 로그인 페이지에 있는 경우는 리다이렉트 하지 않음
    if (router.currentRoute.value.path.includes('/auth/login')) {
        return;
    }

    let errorMessage = '인증이 필요합니다';

    // 서버에서 오는 오류 코드에 따라 사용자에게 표시할 메시지 결정
    if (error && error.code) {
        switch (error.code) {
            case 'token_expired':
                errorMessage = '로그인이 만료되었습니다';
                break;
            case 'token_invalid':
                errorMessage = '유효하지 않은 인증입니다';
                break;
            case 'token_missing':
                errorMessage = '인증 정보가 없습니다';
                break;
            default:
                errorMessage = error.error || errorMessage;
        }
    }

    // 사용자를 로그아웃 처리
    userStore.logout();

    // 로그인 페이지로 이동 (query parameter로 이유 전달)
    router.push({
        path: '/auth/login',
        query: { reason: 'auth_required', message: errorMessage },
    });
}
