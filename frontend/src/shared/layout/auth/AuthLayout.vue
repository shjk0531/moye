<template>
    <div class="w-screen h-screen flex items-center flex-col bg-gray-950">
        <!-- 로고 아이콘 -->
        <div class="flex flex-row justify-center items-center py-10">
            <!-- public/logo.svg 파일 사용 -->
            <div class="bg-gray-50 rounded-full p-0.5">
                <img src="/logo.svg" alt="logo" class="w-20 h-20" />
            </div>
        </div>
        <!-- auth  폼 -->
        <div class="h-fit">
            <router-view />
        </div>
        <!-- 하단 바 -->
        <div class="w-full py-20">
            <div
                class="flex flex-row text-gray-200 text-sm justify-center items-center gap-2"
            >
                <p>후원</p>
                <!-- github 링크 바로가기 -->
                <a href="https://github.com/shjk0531" target="_blank">
                    <p>제작자</p>
                </a>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts"></script>

<style scoped></style>
