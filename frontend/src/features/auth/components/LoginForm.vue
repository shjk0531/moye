<template>
    <div class="flex justify-center flex-col gap-4">
        <form
            @submit.prevent="login"
            class="flex flex-col gap-4 bg-gray-850 rounded-lg p-4 border border-gray-700"
        >
            <div class="flex flex-col gap-1">
                <p class="text-gray-50 text-sm">이메일</p>
                <InputText
                    v-model="email"
                    v-tooltip="'이메일'"
                    class="!bg-gray-950 !text-white !border-gray-700 !border !rounded-md !px-2 !py-1 !text-sm !w-(--custom-auth-input-width)"
                    required
                />
            </div>
            <div class="flex flex-col gap-1">
                <div class="flex flex-row gap-1 justify-between items-center">
                    <p class="text-gray-50 text-sm">비밀번호</p>
                    <!-- 비밀번호 찾기 -->
                    <router-link
                        to="/auth/find-password"
                        class="text-blue-500 text-xs hover:underline"
                        >비밀번호를 잊으셨나요?</router-link
                    >
                </div>
                <Password
                    name="password"
                    required
                    class="!bg-gray-950 !text-white !border-gray-700 !border !rounded-md !px-2 !py-1 !text-sm !w-(--custom-auth-input-width)"
                />
            </div>
            <Button
                type="submit"
                class="!bg-green-700 !text-white !border-gray-700 !border !rounded-md !py-1 !hover:bg-green-700"
                :disabled="loading"
            >
                로그인
            </Button>
            <p v-if="error" class="text-red-600 text-center text-sm">
                {{ error }}
            </p>
        </form>
        <div
            class="flex flex-col gap-1 items-center border border-gray-700 p-4 rounded-lg"
        >
            <p class="text-gray-50 text-sm">
                아직 회원이 아니신가요?
                <router-link
                    to="/auth/signup"
                    class="text-blue-500 hover:underline"
                    >회원가입</router-link
                >
            </p>
        </div>
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
