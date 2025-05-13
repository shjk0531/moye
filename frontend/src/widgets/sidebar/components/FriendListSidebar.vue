<template>
    <div class="flex flex-col">
        <div
            class="flex channel-notice text-white p-2 h-(--custom-notice-bar-height) items-center justify-center"
        >
            <p
                class="text-lg font-bold w-full justify-center items-center text-center"
            >
                메인 페이지
            </p>
        </div>
        <div class="w-(--custom-item-list-width) flex flex-col">
            <div
                class="panel-item p-1 cursor-pointer flex items-center hover:bg-gray-750 text-sm p-2 rounded-lg mx-2 my-0.5 text-gray-100"
            >
                테스트
            </div>
        </div>
        <!-- 구분선 -->
        <div class="w-(--custom-item-list-width) flex flex-col">
            <div class="h-px bg-gray-750 w-full"></div>
        </div>
        <!-- 친구 목록 -->
        <div class="w-(--custom-item-list-width) flex flex-col">
            <div
                class="panel-item p-1 cursor-pointer flex items-center hover:bg-gray-750 text-sm p-2 rounded-lg mx-2 my-0.5 text-gray-100"
            >
                테스트
            </div>
        </div>
    </div>
</template>

<script setup lang="ts"></script>
