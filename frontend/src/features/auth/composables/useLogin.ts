import { ref } from 'vue';
import { loginApi } from '@/entities/user';
import type { User } from '@/entities/user';
import { saveToken, saveUser } from '@/entities/user';

export function useLogin(emit: (event: 'success', user: User) => void) {
    const email = ref('');
    const password = ref('');
    const loading = ref(false);
    const error = ref('');

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
            const result = await loginApi(email.value, password.value);

            // 토큰 저장
            saveToken(result.access_token);

            // 사용자 정보 생성 및 이벤트 발생
            const userData: User = {
                email: email.value,
                id: 0, // API 응답에서 사용자 ID를 얻을 수 있다면 여기에 설정
                name: email.value.split('@')[0], // 임시로 이메일 아이디를 이름으로 사용
            };

            // 사용자 정보 저장
            saveUser(userData);

            emit('success', userData);
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
