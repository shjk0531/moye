import { ref } from 'vue';
import { signupApi } from '@/entities/user';

export function useSignup(
    emit: (event: 'success' | 'error', message: string) => void,
) {
    const email = ref('');
    const password = ref('');
    const confirmPassword = ref('');
    const nickname = ref('');
    const loading = ref(false);
    const error = ref('');

    const validateInput = (): boolean => {
        if (
            !email.value ||
            !password.value ||
            !confirmPassword.value ||
            !nickname.value
        ) {
            error.value = '모든 필수 항목을 입력해주세요.';
            return false;
        }

        if (password.value !== confirmPassword.value) {
            error.value = '비밀번호가 일치하지 않습니다.';
            return false;
        }

        // 기본적인 이메일 형식 검증
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        if (!emailRegex.test(email.value)) {
            error.value = '유효한 이메일 주소를 입력해주세요.';
            return false;
        }

        // 비밀번호 강도 검증 (최소 8자 이상)
        if (password.value.length < 8) {
            error.value = '비밀번호는 최소 8자 이상이어야 합니다.';
            return false;
        }

        return true;
    };

    const signup = async () => {
        if (!validateInput()) return;

        try {
            loading.value = true;
            error.value = '';

            const result = await signupApi({
                email: email.value,
                password: password.value,
                nickname: nickname.value,
            });

            emit('success', '회원가입이 완료되었습니다. 로그인해주세요.');
        } catch (e: any) {
            console.error('회원가입 오류:', e);
            error.value =
                e.response?.data?.message || '회원가입 중 오류가 발생했습니다.';
            emit('error', error.value);
        } finally {
            loading.value = false;
        }
    };

    return {
        email,
        password,
        confirmPassword,
        nickname,
        signup,
        loading,
        error,
    };
}
