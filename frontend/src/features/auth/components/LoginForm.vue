<template>
    <div class="flex justify-center">
        <form @submit.prevent="login" class="flex flex-col gap-4">
            <InputText
                v-model="email"
                v-tooltip="'이메일'"
                placeholder="Email"
                class="!bg-gray-950 !text-white !border-gray-600 !border !rounded-md !px-2 !py-1 !text-sm !w-(--custom-auth-input-width)"
                required
            />
            <Password
                name="password"
                placeholder="Password"
                required
                class="!bg-gray-950 !text-white !border-gray-600 !border !rounded-md !px-2 !py-1 !text-sm !w-(--custom-auth-input-width)"
            />
            <Button
                type="submit"
                class="!bg-green-600 !text-white !border-gray-600 !border !rounded-md !p-2"
                :disabled="loading"
            >
                로그인
            </Button>
            <p v-if="error" class="text-red-600 text-center text-sm">
                {{ error }}
            </p>
        </form>
    </div>
</template>

<script setup lang="ts">
import { useLogin } from '../composables/useLogin';

const emit = defineEmits<{
    (e: 'success', user: any): void;
}>();

const { email, password, login, loading, error } = useLogin(emit);
</script>

<style scoped>
form {
    grid-column: span 12;
}

/* 자동완성 스타일 수정 */
:deep(input:-webkit-autofill),
:deep(input:-webkit-autofill:hover),
:deep(input:-webkit-autofill:focus),
:deep(input:-webkit-autofill:active) {
    -webkit-box-shadow: 0 0 0 30px rgb(3 7 18) inset !important;
    -webkit-text-fill-color: white !important;
    transition: background-color 5000s ease-in-out 0s;
}
</style>
