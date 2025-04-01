<template>
    <div
        class="panel-item p-1 cursor-pointer flex items-center hover:bg-gray-750 text-sm p-2 rounded-lg mx-2 my-0.5"
        :class="[
            item.class ? item.class : '',
            isActive
                ? 'bg-gray-750 text-gray-50'
                : 'text-gray-250 hover:text-gray-50',
        ]"
    >
        <i :class="iconClass"></i>
        <span class="ml-1">{{ item.label }}</span>
    </div>
</template>

<script>
import { channelIcons } from '../types/ChannelIcons';

export default {
    name: 'PanelItem',
    props: {
        item: {
            type: Object,
            required: true,
        },
        isActive: {
            type: Boolean,
            default: false,
        },
    },
    computed: {
        // item.icon은 키 값으로 들어오므로 매핑 객체에서 실제 클래스를 가져옵니다.
        iconClass() {
            return channelIcons[this.item.icon] || '';
        },
    },
};
</script>

<style scoped>
/* 추가 스타일이 필요하면 작성 */
</style>
