// src/features/auth/composables/useLogin.ts
import { ref } from 'vue';
import { loginApi } from '@/entities/user';
import type { User } from '@/entities/user';
import { useUserStore } from '@/store/user';

export function useLogin(emit: (event: 'success', user: User) => void) {
    const email = ref('');
    const password = ref('');
    const loading = ref(false);
    const error = ref('');
    const userStore = useUserStore();

    const validateInput = (): boolean => {
        if (!email.value || !password.value) {
            error.value = '이메일과 비밀번호를 모두 입력해주세요.';
            return false;
        }
        return true;
    };

    const login = async () => {
        if (!validateInput()) return;

        try {
            loading.value = true;
            error.value = '';
            const result = await loginApi({
                email: email.value,
                password: password.value,
            });

            userStore.login(result.user, result.access_token);

            emit('success', result.user);
        } catch (e) {
            console.error('로그인 오류:', e);
            error.value = '이메일 또는 비밀번호가 올바르지 않습니다.';
        } finally {
            loading.value = false;
        }
    };

    return {
        email,
        password,
        login,
        loading,
        error,
    };
}
