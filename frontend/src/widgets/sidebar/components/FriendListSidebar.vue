<template>
    <div class="flex flex-col gap-1">
        <div
            class="flex channel-notice text-white p-2 h-(--custom-notice-bar-height) items-center justify-center"
        >
            <p
                class="text-lg font-bold w-full justify-center items-center text-center"
            >
                {{ currentNotice }}
            </p>
        </div>
        <div class="w-(--custom-item-list-width) flex flex-col">
            <div
                class="panel-item p-1 cursor-pointer flex items-center hover:bg-gray-750 text-sm p-2 rounded-lg mx-2"
                :class="{
                    'bg-gray-750 text-gray-100 font-bold': isFindFriendsActive,
                    'text-gray-200': !isFindFriendsActive,
                }"
                @click="handleFindFriends"
            >
                친구
            </div>
            <div
                class="panel-item p-1 cursor-pointer flex items-center hover:bg-gray-750 text-sm p-2 rounded-lg mx-2"
                :class="{
                    'bg-gray-750 text-gray-100 font-bold': isFindLoungeActive,
                    'text-gray-200': !isFindLoungeActive,
                }"
                @click="handleFindLounge"
            >
                라운지 찾기
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

<script setup lang="ts">
import { computed } from 'vue';
import { useRouter } from 'vue-router';
import { PATHS } from '@/router/paths';

const router = useRouter();
const currentRoute = computed(() => router.currentRoute.value.path);

const currentNotice = computed(() => {
    if (currentRoute.value === `/${PATHS.FRIENDS}`) {
        return '친구 목록';
    } else if (currentRoute.value === `/${PATHS.LOUNGES}`) {
        return '라운지 찾기';
    }
});

function handleFindLounge() {
    router.push(`/${PATHS.LOUNGES}`);
}

function handleFindFriends() {
    router.push(`/${PATHS.FRIENDS}`);
}

const isFindFriendsActive = computed(() => {
    return currentRoute.value === `/${PATHS.FRIENDS}`;
});

const isFindLoungeActive = computed(() => {
    return currentRoute.value === `/${PATHS.LOUNGES}`;
});
</script>
