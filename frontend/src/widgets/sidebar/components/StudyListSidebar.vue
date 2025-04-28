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

        <div class="study-list">
            <ScrollPanel style="height: 100%">
                <div
                    class="flex flex-col items-center justify-start w-full gap-2"
                >
                    <!-- 스터디 아이콘 리스트 -->
                    <StudyIconList />
                    <!-- 새로운 스터디 생성 버튼 -->
                    <div
                        class="relative flex items-center justify-center w-full group cursor-pointer"
                        @click="handleNewStudyClick"
                    >
                        <div
                            class="w-10 h-10 overflow-hidden border-2 rounded-lg flex justify-center items-center p-1 dark:bg-gray-50 hover:bg-gray-200"
                        >
                            <span class="mdi mdi-plus-circle text-2xl"></span>
                        </div>
                    </div>
                </div>
            </ScrollPanel>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router';
import { StudyIconList } from '@/entities/study';
import ScrollPanel from 'primevue/scrollpanel';
import { PATHS } from '@/router/paths';
import { useAppStore } from '@/store';

const router = useRouter();
const route = useRoute();
const appStore = useAppStore();

function handleTitleIconClick() {
    router.push({ path: PATHS.ROOT });
}

const isTitleActive = () => route.path === `/${PATHS.ME}`;

function handleNewStudyClick() {
    appStore.saveLastRoute(route.path);
    router.push(`/${PATHS.CREATE}`);
}
</script>

<style scoped lang="scss">
.study-list {
    flex: 1;
    overflow-y: auto;
    overflow-x: hidden;
}

.study-list :deep(.p-scrollpanel-content) {
    height: 100%;
}

.study-list :deep(.p-scrollpanel-bar) {
    background-color: var(--custom-scrollbar-color);
    border-radius: 4px;
    transition: background-color 0.3s ease;
}

.study-list :deep(.p-scrollpanel-bar:hover) {
    background-color: var(--custom-scrollbar-color-hover);
}

.study-list :deep(.p-scrollpanel-bar:active) {
    background-color: var(--custom-scrollbar-color-active);
}

.study-list :deep(.p-scrollpanel-bar-y) {
    width: calc(var(--spacing)) !important;
}
</style>
