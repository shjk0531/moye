<template>
    <div class="login-container">
        <Toast />
        <div
            v-if="showAuthMessage"
            class="auth-message p-3 mb-4 surface-card border-round"
        >
            <i class="pi pi-exclamation-circle mr-2 text-yellow-500"></i>
            {{ authMessage }}
        </div>
        <LoginForm @success="onLoginSuccess" />
    </div>
</template>

<script lang="ts" setup>
import { LoginForm } from '@/features/auth';
import { onMounted, ref, computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useToast } from 'primevue/usetoast';
import type { User } from '@/entities/user';

const router = useRouter();
const route = useRoute();
const toast = useToast();

// 인증 메시지 관련 상태
const authMessage = ref('');
const showAuthMessage = computed(() => !!authMessage.value);

// 로그인 성공 처리
const onLoginSuccess = (user: User) => {
    router.push('/');
};

// 인증 오류로 인한 리다이렉트 처리
onMounted(() => {
    if (route.query.reason === 'auth_required' && route.query.message) {
        authMessage.value = route.query.message as string;

        // 3초 후 메시지 자동 제거
        setTimeout(() => {
            authMessage.value = '';
        }, 5000);
    }
});
</script>

<style scoped>
.login-container {
    max-width: 450px;
    margin: 0 auto;
    padding: 2rem;
}

.auth-message {
    background-color: var(--yellow-50);
    color: var(--yellow-800);
    border: 1px solid var(--yellow-200);
}
</style>
