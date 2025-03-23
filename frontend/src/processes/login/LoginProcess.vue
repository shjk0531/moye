<template>
  <section>
    <h1>로그인</h1>
    <LoginForm @success="onSuccess" />
    <p v-if="message">{{ message }}</p>
  </section>
</template>

<script setup lang="ts">
import { LoginForm } from "@/features/auth";
import { useLoginProcess } from "./useLoginProcess";

const { handleLoginSuccess, message } = useLoginProcess();

function onSuccess(user: any) {
  handleLoginSuccess(user);
}
</script>
