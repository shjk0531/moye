import { ref } from 'vue';
import { loginApi } from '../api/authApi';

export function useLogin(emit: (event: 'success', user: any) => void) {
    const email = ref('');
    const password = ref('');
    const loading = ref(false);
    const error = ref('');

    const login = async () => {
        try {
            loading.value = true;
            error.value = '';
            const result = await loginApi(email.value, password.value);
            localStorage.setItem('token', result.token);
            emit('success', result.user);
        } catch (e) {
            error.value = '로그인 실패.';
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
