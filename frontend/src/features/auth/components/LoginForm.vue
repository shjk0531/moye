<template>
  <form @submit.prevent="login">
    <input v-model="email" type="email" placeholder="이메일" required />
    <input v-model="password" type="password" placeholder="비밀번호" required />
    <button type="submit" :disabled="loading">로그인</button>
    <p v-if="error">{{ error }}</p>
  </form>
</template>

<script setup lang="ts">
import { useLogin } from "../composables/useLogin";

const emit = defineEmits<{
  (e: "success", user: any): void;
}>();

const { email, password, login, loading, error } = useLogin(emit);
</script>

<style scoped>
form {
  grid-column: span 12;
}
</style>
