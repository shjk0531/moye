<!-- src/shared/panel/components/PanelList.vue -->
<template>
    <div class="w-(--custom-item-list-width) flex flex-col">
        <div v-for="(item, index) in mergedList" :key="`${item.id} - ${index}`">
            <!-- 그룹 데이터면 PanelSection 컴포넌트 사용 -->
            <PanelSection
                v-if="item.type === 'group'"
                :label="item.label"
                :items="item.items"
                :activeItemId="activeItemId"
                :listType="listType"
                @item-click="handleItemClick"
            />
            <!-- 일반 아이템이면 PanelItem 컴포넌트 사용 -->
            <PanelItem
                v-else
                :item="item"
                :isActive="item.id === activeItemId"
                :listType="listType"
                @click="handleItemClick(item)"
            />
        </div>
    </div>
</template>

<script setup>
import { computed } from 'vue';
import PanelSection from './PanelSection.vue';
import PanelItem from './PanelItem.vue';
import { mergePanelItems } from '@/shared/lib/panelUtils';

const props = defineProps({
    groups: {
        type: Array,
        default: () => [],
    },
    items: {
        type: Array,
        default: () => [],
    },
    activeItemId: {
        type: [Number, String],
        default: null,
    },
    listType: {
        type: String,
        default: 'channel',
    },
});

const emits = defineEmits(['item-click']);

// 공통 유틸리티 함수를 사용하여 그룹화 및 정렬된 아이템 목록 생성
const mergedList = computed(() => {
    return mergePanelItems(props.groups, props.items);
});

function handleItemClick(item) {
    emits('item-click', item);
}
</script>

<style scoped></style>
