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
import { channelIcons } from '../types/ChannelIcons';

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
    return channelIcons[props.item.icon] || '';
});
</script>

<style scoped>
/* 추가 스타일이 필요하면 작성 */
</style>
