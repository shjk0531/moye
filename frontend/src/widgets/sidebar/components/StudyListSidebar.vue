<!-- src/widgets/sidebar/components/StudyListSidebar.vue -->
<template>
    <!-- Title Icon: '/me'로 이동 -->
    <div class="w-(--custom-study-list-width) flex flex-col">
        <div
            class="flex flex-col items-center justify-center w-full gap-2 h-(--custom-notice-bar-height)"
        >
            <div
                class="relative flex items-center justify-center w-full group cursor-pointer"
                @click="handleTitleIconClick"
            >
                <!-- indicator: 현재 경로가 '/me'면 상시 활성 -->
                <div
                    :class="[
                        'absolute left-0 top-1/2 -translate-y-1/2 rounded transition-all duration-300',
                        isTitleActive()
                            ? 'w-1 h-[70%] bg-gray-50'
                            : 'w-0 group-hover:w-1 h-[50%] bg-gray-150',
                    ]"
                ></div>
                <div
                    class="w-10 h-10 mx-2 my-1 overflow-hidden border-2 rounded-lg"
                >
                    <img
                        src="/logo.svg"
                        alt="logo"
                        class="w-full h-full bg-gray-50"
                    />
                </div>
            </div>
        </div>

        <!-- 스터디 아이콘 리스트 -->
        <StudyIconList />
    </div>
</template>

<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router';
import { StudyIconList } from '@/entities/study';
import { PATHS } from '@/router/paths';

const router = useRouter();
const route = useRoute();

function handleTitleIconClick() {
    router.push({ path: PATHS.ROOT });
}

const isTitleActive = () => route.path === `/${PATHS.ME}`;
</script>

<style scoped lang="scss"></style>
