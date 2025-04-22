<!-- src/features/auth/components/SignupFrom.vue -->

<template>
    <div class="flex justify-center flex-col gap-4">
        <form
            @submit.prevent="signUp"
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
                <p class="text-gray-50 text-sm">비밀번호</p>
                <Password
                    v-tooltip="'비밀번호'"
                    class="!bg-gray-950 !text-white !border-gray-700 !border !rounded-md !px-2 !py-1 !text-sm !w-(--custom-auth-input-width)"
                    required
                />
            </div>
            <div class="flex flex-col gap-1">
                <p class="text-gray-50 text-sm">비밀번호 확인</p>
                <Password
                    v-tooltip="'비밀번호 확인'"
                    class="!bg-gray-950 !text-white !border-gray-700 !border !rounded-md !px-2 !py-1 !text-sm !w-(--custom-auth-input-width)"
                    required
                />
            </div>
            <div class="flex flex-col gap-1">
                <p class="text-gray-50 text-sm">닉네임</p>
                <InputText
                    v-tooltip="'이름'"
                    class="!bg-gray-950 !text-white !border-gray-700 !border !rounded-md !px-2 !py-1 !text-sm !w-(--custom-auth-input-width)"
                    required
                />
            </div>
            <div class="flex flex-col gap-1">
                <button
                    type="submit"
                    class="bg-green-700 text-white border border-gray-700 rounded-md py-1 hover:bg-green-600 cursor-pointer"
                >
                    회원가입
                </button>
            </div>
        </form>
        <div
            class="flex flex-col gap-1 items-center border border-gray-700 p-4 rounded-lg"
        >
            <p class="text-gray-50 text-sm">
                이미 회원이신가요?
                <router-link
                    to="/auth/login"
                    class="text-blue-500 hover:underline"
                    >로그인</router-link
                >
            </p>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const email = ref('');

const signUp = () => {
    console.log(email.value);
};
</script>

<style scoped>
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
