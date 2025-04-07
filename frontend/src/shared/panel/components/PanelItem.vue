<!-- src/shared/panel/components/PanelItem.vue -->
<template>
    <div
        class="panel-item p-1 cursor-pointer flex items-center hover:bg-gray-750 text-sm p-2 rounded-lg mx-2 my-0.5"
        :class="[
            item.class || '',
            isItemActive
                ? 'bg-gray-750 text-gray-50'
                : 'text-gray-250 hover:text-gray-50',
        ]"
    >
        <i :class="iconClass"></i>
        <span class="ml-1">{{ item.label }}</span>
    </div>
</template>

<script setup>
import { computed } from 'vue';
import { useRoute } from 'vue-router';
import { channelIcons } from '../types/ChannelIcons'; // 필요에 따라 아이콘 설정 로직 재사용 또는 일반화
import { getParamName } from '@/widgets/notice/services/navigationService';

const route = useRoute();

const props = defineProps({
    item: {
        type: Object,
        required: true,
    },
    isActive: {
        type: Boolean,
        default: false,
    },
    listType: {
        type: String,
        default: 'channel', // 기본값으로 'channel' 설정
    },
});

const iconClass = computed(() => {
    // 아이콘 설정 로직은 필요에 따라 분기 처리 가능
    return channelIcons[props.item.icon] || '';
});

// URL 기반으로 현재 아이템이 활성 상태인지 확인
const isUrlActive = computed(() => {
    // 지원되는 모든 아이템 타입에 대해 URL의 ID 파라미터와 아이템 ID 비교
    const paramName = getParamName(props.listType);
    return route.params[paramName] === String(props.item.id);
});

// isActive prop과 URL 기반 활성 상태 중 하나라도 true면 아이템을 활성화
const isItemActive = computed(() => {
    return props.isActive || isUrlActive.value;
});
</script>

<style scoped>
/* 필요에 따른 추가 스타일 */
</style>
