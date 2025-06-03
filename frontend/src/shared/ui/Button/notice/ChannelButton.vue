<template>
    <span :class="iconClasses" @click="handleClick" :title="title"></span>
</template>

<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router';
import { computed } from 'vue';

const router = useRouter();
const route = useRoute();
const isActive = computed(() => route.path.includes('/channel'));

const iconClasses = computed(() => {
    const baseClasses = `mdi mdi-message hover:text-gray-200 cursor-pointer`;
    return `${baseClasses} ${isActive ? 'text-gray-150' : 'text-gray-400'}`;
});

const title = computed(() => {
    return '채널';
});

const handleClick = async () => {
    const loungeId = route.params.loungeId as string;

    if (!loungeId) {
        console.error('유효한 loungeId가 없습니다.');
        return;
    }

    // await navigateToItem('channel', loungeId, router);
};
</script>

<style scoped>
/* 필요한 스타일 추가 */
</style>
