<template>
    <div class="flex justify-center flex-col gap-4">
        <form
            @submit.prevent="login"
            class="flex flex-col gap-4 bg-gray-850 rounded-lg p-4 border border-gray-700"
        >
            <div class="flex flex-col gap-1">
                <label class="text-gray-50 text-sm">이메일</label>
                <input
                    v-model="email"
                    type="email"
                    required
                    class="bg-gray-950 text-white border border-gray-700 rounded-md px-2 py-1 text-sm w-(--custom-auth-input-width)"
                    placeholder="이메일"
                />
            </div>

            <div class="flex flex-col gap-1">
                <div class="flex flex-row gap-1 justify-between items-center">
                    <label class="text-gray-50 text-sm">비밀번호</label>
                    <router-link
                        to="/auth/find-password"
                        class="text-blue-500 text-xs hover:underline"
                    >
                        비밀번호를 잊으셨나요?
                    </router-link>
                </div>
                <div class="relative w-(--custom-auth-input-width)">
                    <input
                        v-model="password"
                        :type="showPassword ? 'text' : 'password'"
                        required
                        autocomplete="current-password"
                        class="bg-gray-950 text-white border border-gray-700 rounded-md px-2 py-1 text-sm w-full"
                        placeholder="비밀번호"
                    />
                    <button
                        type="button"
                        class="absolute inset-y-0 right-0 flex items-center px-2 text-gray-400 hover:text-gray-200"
                        @click="showPassword = !showPassword"
                    >
                        <span
                            class="mdi"
                            :class="showPassword ? 'mdi-eye-off' : 'mdi-eye'"
                        ></span>
                    </button>
                </div>
            </div>

            <button
                type="submit"
                class="bg-green-700 text-white border border-gray-700 rounded-md py-1 hover:bg-green-600 cursor-pointer"
                :disabled="loading"
            >
                로그인
            </button>

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
                >
                    회원가입
                </router-link>
            </p>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useLogin } from '../composables/useLogin';

const emit = defineEmits<{ (e: 'success', user: any): void }>();

const { email, password, login, loading, error } = useLogin(emit);

const showPassword = ref(false);
</script>

<style scoped>
input:-webkit-autofill,
input:-webkit-autofill:hover,
input:-webkit-autofill:focus,
input:-webkit-autofill:active {
    -webkit-box-shadow: 0 0 0 rgb(3 7 18) inset !important;
    -webkit-text-fill-color: white !important;
    transition: background-color 5000s ease-in-out 0s;
}
</style>
