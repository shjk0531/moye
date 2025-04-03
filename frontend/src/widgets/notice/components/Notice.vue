<template>
    <div class="notice-panel">
        <div
            class="flex flex-row justify-between items-center bg-gray-850 px-4 py-2"
        >
            <div class="notice-header">
                <h2 class="text-lg font-semibold text-gray-50">Notice</h2>
            </div>
            <div class="notice-icons flex flex-row gap-4 text-2xl">
                <span
                    v-for="(item, idx) in icons"
                    :key="idx"
                    :class="getIconClass(item)"
                    @click="handleClick(item)"
                ></span>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
export default {
    name: 'Notice',
    props: {
        isMemberListVisible: {
            type: Boolean,
            default: false,
        },
    },
    data() {
        return {
            icons: [
                {
                    icon: 'mdi mdi-message',
                    color: 'text-gray-400',
                    hover: 'hover:text-gray-200',
                    active: 'text-gray-150',
                },
                {
                    icon: 'mdi-calendar-check',
                    color: 'text-gray-400',
                    hover: 'hover:text-gray-200',
                    active: 'text-gray-150',
                },
                {
                    icon: 'mdi-post',
                    color: 'text-gray-400',
                    hover: 'hover:text-gray-200',
                    active: 'text-gray-150',
                },
                {
                    icon: 'mdi-note',
                    color: 'text-gray-400',
                    hover: 'hover:text-gray-200',
                    active: 'text-gray-150',
                },
                {
                    icon: 'mdi-account-group',
                    color: 'text-gray-400',
                    hover: 'hover:text-gray-200',
                    active: 'text-gray-150',
                },
            ],
        };
    },
    methods: {
        handleClick(item: { icon: string; color: string; hover: string }) {
            // account-group 아이콘 클릭 시 toggle 이벤트 emit
            if (item.icon === 'mdi-account-group') {
                this.$emit('toggle-member-list');
            }
            // 다른 아이콘 클릭 시 필요한 동작 추가 가능
        },
        getIconClass(item: {
            icon: string;
            color: string;
            hover: string;
            active: string;
        }) {
            // 기본 클래스 문자열 구성
            let classes = `mdi ${item.icon} ${item.hover} cursor-pointer `;
            // account-group 아이콘의 경우 isMemberListVisible 값에 따라 색상 변경
            if (item.icon === 'mdi-account-group') {
                classes += this.isMemberListVisible ? item.active : item.color;
            } else {
                classes += item.color;
            }
            return classes;
        },
    },
};
</script>

<style scoped>
.notice-panel {
    grid-row: titlebarEnd / noticeEnd;
    grid-column: channelEnd / end;
}
</style>
