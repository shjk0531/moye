<template>
  <div class="auth-page">
    <h1>로그인</h1>
    <LoginProcess />
  </div>
</template>

<script setup lang="ts">
import { LoginProcess } from "@/processes/login";
</script>

<style scoped>
.auth-page {
  grid-column: span 12;
}
</style>
