<template>
    <span :class="iconClasses" @click="handleClick" :title="title"></span>
</template>

<script setup lang="ts">
import { useNoticeStore } from '@/store';
import { computed } from 'vue';
const noticeStore = useNoticeStore();

const iconClasses = computed(() => {
    const baseClasses = `mdi mdi-account-group hover:text-gray-200 cursor-pointer`;
    return `${baseClasses} ${
        noticeStore.isMemberListVisible ? 'text-gray-150' : 'text-gray-400'
    }`;
});

const title = computed(() => '멤버');

const handleClick = () => {
    noticeStore.toggleMemberList();
};
</script>

<style scoped>
/* 필요한 스타일 추가 */
</style>
