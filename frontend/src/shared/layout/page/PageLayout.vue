<template>
    <div class="page-layout flex flex-col w-full h-full dark:bg-gray-850">
        <!-- 상단 알림 영역 -->
        <div
            class="flex dark:bg-gray-900 px-4 py-2 h-(--custom-notice-bar-height)"
        >
            <slot name="notice"></slot>
        </div>
        <!-- 메인 콘텐츠 영역 -->
        <div class="flex flex-row flex-1 overflow-hidden">
            <!-- 주요 콘텐츠 영역 -->
            <div class="content-area flex flex-col flex-1 overflow-hidden">
                <slot name="content"></slot>
            </div>
            <!-- 오른쪽 사이드바 영역 -->
            <div
                v-if="$slots.sidebar && isNoticeVisible"
                class="right-sidebar dark:bg-gray-900 w-(--custom-right-sidebar-width)"
            >
                <slot name="sidebar"></slot>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useNoticeStore } from '@/store';
import { computed } from 'vue';

const noticeStore = useNoticeStore();

const isNoticeVisible = computed(() => noticeStore.isMemberListVisible);
</script>

<style scoped>
.right-sidebar {
    border-left: 1px solid var(--surface-border);
}
</style>
