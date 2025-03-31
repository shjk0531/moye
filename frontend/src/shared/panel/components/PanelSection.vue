<template>
    <div class="panel-section my-3">
        <div
            class="p-1 flex justify-between items-center cursor-pointer text-gray-250 hover:text-gray-50 text-sm m-2"
            @click="toggle"
        >
            <div class="flex items-center justify-center">
                <span class="ml-1">{{ label }}</span>
                <span class="items-center justify-center flex">
                    <i
                        class="pi pi-angle-up transition-transform duration-300 text-sm ml-1"
                        :class="{ 'rotate-180': !expanded }"
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
                />
            </div>
        </transition>
        <div v-if="!expanded && activeItem">
            <PanelItem :item="activeItem" :isActive="true" />
        </div>
    </div>
</template>

<script>
import PanelItem from './PanelItem.vue';

export default {
    name: 'PanelSection',
    components: {
        PanelItem,
    },
    props: {
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
            default: false,
        },
        // 현재 사용자가 들어가 있는 채널의 id
        activeChannelId: {
            type: [Number, String],
            default: null,
        },
    },
    data() {
        return {
            expanded: this.defaultExpanded,
        };
    },
    computed: {
        // 그룹 내 채널을 order 기준으로 정렬
        sortedItems() {
            return [...this.items].sort(
                (a, b) => (a.order || 0) - (b.order || 0),
            );
        },
        // 그룹 내에서 active 채널을 찾습니다.
        activeItem() {
            return this.items.find((item) => item.id === this.activeChannelId);
        },
    },
    methods: {
        toggle() {
            this.expanded = !this.expanded;
            this.$emit('toggle', this.expanded);
        },
    },
};
</script>

<style scoped></style>
