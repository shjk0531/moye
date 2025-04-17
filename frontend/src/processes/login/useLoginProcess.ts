import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useUserStore } from '@/store';

export function useLoginProcess() {
    const router = useRouter();
    const message = ref('');
    const userStore = useUserStore();

    const handleLoginSuccess = (user: any) => {
        // userStore.setUser(user);
        // message.value = `${user.name}님 환영합니다!`;
        // setTimeout(() => {
        //     router.push('/');
        // }, 1500);
    };

    return {
        message,
        handleLoginSuccess,
    };
}
