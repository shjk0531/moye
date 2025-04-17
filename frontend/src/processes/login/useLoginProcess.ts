import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useUserStore } from '@/store';
import type { User } from '@/entities/user';

export function useLoginProcess() {
    const router = useRouter();
    const message = ref('');
    const userStore = useUserStore();

    const handleLoginSuccess = (user: User) => {
        userStore.setUser(user);
        message.value = `${user.nickname}님 환영합니다!`;
        setTimeout(() => {
            router.push('/');
        }, 1500);
    };

    return {
        message,
        handleLoginSuccess,
    };
}
