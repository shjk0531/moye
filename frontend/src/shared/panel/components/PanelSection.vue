<!-- src/shared/panel/components/PanelSection.vue -->
<template>
    <div class="panel-section my-3">
        <div
            class="p-1 flex justify-between items-center cursor-pointer text-gray-250 hover:text-gray-50 text-sm m-2"
            @click="toggle"
        >
            <div class="flex items-center justify-center">
                <span class="ml-1">{{ label }}</span>
                <span class="flex items-center justify-center">
                    <i
                        class="mdi mdi-chevron-down transition-transform duration-300 text-md ml-1 font-bold"
                        :class="{ 'rotate--90': !expanded }"
                    ></i>
                </span>
            </div>
        </div>
        <!-- 그룹이 열렸을 때: 모든 아이템 렌더링 -->
        <transition name="panel-transition">
            <div v-if="expanded" class="panel-body">
                <PanelItem
                    v-for="item in sortedItems"
                    :key="item.id"
                    :item="item"
                    :isActive="item.id === activeItemId"
                    :listType="listType"
                    @click="handleItemClick(item)"
                />
            </div>
        </transition>
        <!-- 그룹이 닫혔을 때, 현재 활성 아이템만 표시 -->
        <div v-if="!expanded && activeItemCombined">
            <PanelItem
                :item="activeItemCombined"
                :isActive="true"
                :listType="listType"
                @click="handleItemClick(activeItemCombined)"
            />
        </div>
    </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import { useRoute } from 'vue-router';
import PanelItem from './PanelItem.vue';

const route = useRoute();

const props = defineProps({
    label: {
        type: String,
        required: true,
    },
    items: {
        type: Array,
        default: () => [],
    },
    defaultExpanded: {
        type: Boolean,
        default: true,
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

const emits = defineEmits(['item-click', 'toggle']);

const expanded = ref(props.defaultExpanded);

// 그룹 내 아이템들을 order 기준으로 정렬
const sortedItems = computed(() => {
    return [...props.items].sort((a, b) => (a.order || 0) - (b.order || 0));
});

// ID 기반으로 현재 그룹에서 활성 아이템 찾기
const activeItemById = computed(() => {
    return props.items.find((item) => item.id === props.activeItemId);
});

// URL 기반으로 현재 그룹에서 활성 아이템 찾기
const activeItemByUrl = computed(() => {
    // 리스트 타입에 따라 URL 파라미터 확인
    if (props.listType === 'channel') {
        const channelId = route.params.channelId;
        if (channelId) {
            return props.items.find((item) => String(item.id) === channelId);
        }
    }
    // 추가적인 리스트 타입이 있다면 여기에 확장
    return null;
});

// ID 기반 또는 URL 기반 활성 아이템 중 하나라도 있으면 반환
const activeItemCombined = computed(() => {
    return activeItemByUrl.value || activeItemById.value;
});

function toggle() {
    expanded.value = !expanded.value;
    emits('toggle', expanded.value);
}

function handleItemClick(item) {
    emits('item-click', item);
}
</script>

<style scoped>
.rotate--90 {
    transform: rotate(-90deg);
}
</style>
