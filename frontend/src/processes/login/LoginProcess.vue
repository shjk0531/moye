<template>
    <section class="login-process">
        <LoginForm @success="onSuccess" />
        <Transition name="fade">
            <p v-if="message" class="welcome-message">{{ message }}</p>
        </Transition>
    </section>
</template>

<script setup lang="ts">
import { LoginForm } from '@/features/auth';
import { useLoginProcess } from './useLoginProcess';
import type { User } from '@/entities/user';

const { handleLoginSuccess, message } = useLoginProcess();

function onSuccess(user: User) {
    handleLoginSuccess(user);
}
</script>

<style scoped>
.login-process {
    display: flex;
    flex-direction: column;
    align-items: center;
}

.welcome-message {
    margin-top: 1rem;
    color: #10b981;
    font-weight: 500;
}

.fade-enter-active,
.fade-leave-active {
    transition: opacity 0.5s ease;
}

.fade-enter-from,
.fade-leave-to {
    opacity: 0;
}
</style>
