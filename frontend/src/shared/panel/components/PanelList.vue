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
                @item-click="handleItemClick"
            />
            <!-- 일반 아이템이면 PanelItem 컴포넌트 사용 -->
            <PanelItem
                v-else
                :item="item"
                :isActive="item.id === activeItemId"
                @click="handleItemClick(item)"
            />
        </div>
    </div>
</template>

<script setup>
import { computed, defineProps, defineEmits } from 'vue';
import PanelSection from './PanelSection.vue';
import PanelItem from './PanelItem.vue';

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
});

const emits = defineEmits(['item-click']);

const mergedList = computed(() => {
    // 그룹에 속한 아이템들을 해당 그룹 객체 안에 items 배열로 추가
    const mergedGroups = props.groups.map((group) => {
        const groupItems = props.items
            .filter((item) => item.groupId === group.id)
            .sort((a, b) => (a.order || 0) - (b.order || 0));
        return { ...group, type: 'group', items: groupItems };
    });
    // 그룹에 속하지 않은 일반 아이템
    const ungroupedItems = props.items
        .filter((item) => item.groupId === null)
        .map((item) => ({ ...item, type: 'ungrouped' }));
    // 전체를 order 값 기준으로 정렬
    return [...mergedGroups, ...ungroupedItems].sort(
        (a, b) => (a.order || 9999) - (b.order || 9999),
    );
});

function handleItemClick(item) {
    emits('item-click', item);
}
</script>

<style scoped></style>
