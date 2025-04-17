import { ref } from 'vue';
import { removeToken, removeUser } from '@/entities/user';

export function useLogout(emit: (event: 'success') => void) {
    const loading = ref(false);
    const error = ref('');

    const logout = async () => {
        try {
            loading.value = true;

            // 토큰 및 사용자 정보 삭제
            removeToken();
            removeUser();

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
