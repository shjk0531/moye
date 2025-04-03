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
        <!-- 그룹이 열렸을 때: 모든 채널 렌더링 -->
        <transition name="panel-transition">
            <div v-if="expanded" class="panel-body">
                <PanelItem
                    v-for="item in sortedItems"
                    :key="item.id"
                    :item="item"
                    :isActive="item.id === activeChannelId"
                    @click="handleItemClick(item)"
                />
            </div>
        </transition>
        <!-- 그룹이 닫혔을 때, 현재 활성 채널만 표시 -->
        <div v-if="!expanded && activeItem">
            <PanelItem
                :item="activeItem"
                :isActive="true"
                @click="handleItemClick(activeItem)"
            />
        </div>
    </div>
</template>

<script setup>
import { ref, computed, defineProps, defineEmits } from 'vue';
import PanelItem from './PanelItem.vue';

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
    activeChannelId: {
        type: [Number, String],
        default: null,
    },
});

const emits = defineEmits(['channel-click', 'toggle']);

const expanded = ref(props.defaultExpanded);

// 그룹 내 채널들을 order 기준으로 정렬
const sortedItems = computed(() => {
    return [...props.items].sort((a, b) => (a.order || 0) - (b.order || 0));
});

// 현재 그룹에서 활성 채널 찾기
const activeItem = computed(() => {
    return props.items.find((item) => item.id === props.activeChannelId);
});

function toggle() {
    expanded.value = !expanded.value;
    emits('toggle', expanded.value);
}

function handleItemClick(item) {
    emits('channel-click', item);
}
</script>

<style scoped>
.rotate--90 {
    transform: rotate(-90deg);
}
</style>
