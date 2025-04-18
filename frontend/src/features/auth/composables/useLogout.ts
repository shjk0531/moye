import { useUserStore } from '@/store/user';
import { ref } from 'vue';

export function useLogout(emit: (event: 'success') => void) {
    const loading = ref(false);
    const error = ref('');
    const userStore = useUserStore();

    const logout = async () => {
        try {
            loading.value = true;

            userStore.logout();

            emit('success');
        } catch (e) {
            console.error('로그아웃 오류:', e);
            error.value = '로그아웃 중 오류가 발생했습니다.';
        } finally {
            loading.value = false;
        }
    };

    return {
        logout,
        loading,
        error,
    };
}
