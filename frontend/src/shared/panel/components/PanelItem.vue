<!-- src/shared/panel/components/PanelItem.vue -->
<template>
    <div
        class="panel-item p-1 cursor-pointer flex items-center hover:bg-gray-750 text-sm p-2 rounded-lg mx-2 my-0.5"
        :class="[
            item.class || '',
            isActive
                ? 'bg-gray-750 text-gray-50'
                : 'text-gray-250 hover:text-gray-50',
        ]"
    >
        <i :class="iconClass"></i>
        <span class="ml-1">{{ item.label }}</span>
    </div>
</template>

<script setup>
import { computed, defineProps } from 'vue';
import { channelIcons } from '../types/ChannelIcons'; // 필요에 따라 아이콘 설정 로직 재사용 또는 일반화

const props = defineProps({
    item: {
        type: Object,
        required: true,
    },
    isActive: {
        type: Boolean,
        default: false,
    },
});

const iconClass = computed(() => {
    // 아이콘 설정 로직은 필요에 따라 분기 처리 가능
    return channelIcons[props.item.icon] || '';
});
</script>

<style scoped>
/* 필요에 따른 추가 스타일 */
</style>
