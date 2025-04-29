import { useToast } from 'primevue/usetoast';
import { useRouter } from 'vue-router';
import { useUserStore } from '@/store';

export function useAuthRedirect() {
    const toast = useToast();
    const router = useRouter();
    const userStore = useUserStore();

    /**
     * 인증 만료 처리 및 로그인 페이지로 리다이렉트
     * @param message 표시할 메시지
     */
    const redirectToLogin = (message = '로그인이 필요합니다') => {
        // 이미 로그인 페이지에 있으면 중복 처리 방지
        if (router.currentRoute.value.path.includes('/auth/login')) {
            return;
        }

        // 사용자에게 알림 표시
        toast.add({
            severity: 'warn',
            summary: '인증 만료',
            detail: message,
            life: 3000,
        });

        // 로그아웃 처리
        userStore.logout();

        // 로그인 페이지로 리다이렉트
        router.push('/auth/login');
    };

    return {
        redirectToLogin,
    };
}
