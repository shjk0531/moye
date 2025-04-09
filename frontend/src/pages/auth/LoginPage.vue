<template>
    <div class="auth-page">
        <LoginProcess />
    </div>
</template>

<script lang="ts">
import { LoginProcess } from '@/processes/login';

export default {
    name: 'LoginPage',
    components: {
        LoginProcess,
    },
};
</script>

<style scoped>
.auth-page {
    grid-column: span 12;
}
</style>
