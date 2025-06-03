<template>
    <span :class="iconClasses" @click="handleClick" :title="title"></span>
</template>

<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router';
import { computed } from 'vue';

const router = useRouter();
const route = useRoute();

// active 여부는 로컬에서만 계산
const isActive = computed(() => route.path.includes('/calendar'));

const iconClasses = computed(
    () =>
        `mdi mdi-calendar-blank hover:text-gray-200 cursor-pointer ${
            isActive.value ? 'text-gray-150' : 'text-gray-400'
        }`,
);

const title = '캘린더';

async function handleClick() {
    const loungeId = route.params.loungeId as string;
    if (!loungeId) {
        console.error('유효한 loungeId가 없습니다.');
        return;
    }
    await router.push(`/lounge/${loungeId}/calendar`);
}
</script>

<style scoped>
/* 필요하면 스타일 추가 */
</style>
