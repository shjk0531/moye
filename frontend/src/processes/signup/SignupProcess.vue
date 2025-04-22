<template>
    <section class="">
        <SignupForm @success="handleSignupSuccess" @error="handleSignupError" />
    </section>
</template>

<script setup lang="ts">
import { SignupForm } from '@/features/auth';
import { useRouter } from 'vue-router';
import { useToast } from 'primevue/usetoast';

const router = useRouter();
const toast = useToast();

const handleSignupSuccess = (message: string) => {
    toast.add({
        severity: 'success',
        summary: '회원가입 성공',
        detail: message,
        life: 3000,
    });
    // 로그인 페이지로 리다이렉트
    router.push('/auth/login');
};

const handleSignupError = (message: string) => {
    toast.add({
        severity: 'error',
        summary: '회원가입 실패',
        detail: message,
        life: 5000,
    });
};
</script>
