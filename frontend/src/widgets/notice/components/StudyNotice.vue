<template>
    <div class="flex flex-row justify-between items-center">
        <div class="notice-header">
            <h2 class="text-lg font-semibold text-gray-50">Notice</h2>
        </div>
        <div class="notice-icons flex flex-row gap-4 text-2xl">
            <StudyNoticeIcon
                v-for="(item, idx) in icons"
                :key="idx"
                :item="item"
                :is-member-list-visible="isMemberListVisible"
                @toggle-member-list="$emit('toggle-member-list')"
                @icon-click="handleIconClick"
            />
        </div>
    </div>
</template>

<script>
import StudyNoticeIcon from './StudyNoticeIcon.vue';

export default {
    name: 'StudyNotice',
    components: {
        StudyNoticeIcon,
    },
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
                    icon: 'mdi-message',
                    type: 'channel',
                    color: 'text-gray-400',
                    hover: 'hover:text-gray-200',
                    active: 'text-gray-150',
                },
                {
                    icon: 'mdi-calendar-check',
                    type: 'calendar',
                    color: 'text-gray-400',
                    hover: 'hover:text-gray-200',
                    active: 'text-gray-150',
                },
                {
                    icon: 'mdi-post',
                    type: 'post',
                    color: 'text-gray-400',
                    hover: 'hover:text-gray-200',
                    active: 'text-gray-150',
                },
                {
                    icon: 'mdi-note',
                    type: 'note',
                    color: 'text-gray-400',
                    hover: 'hover:text-gray-200',
                    active: 'text-gray-150',
                },
                {
                    icon: 'mdi-account-group',
                    type: 'member',
                    color: 'text-gray-400',
                    hover: 'hover:text-gray-200',
                    active: 'text-gray-150',
                },
            ],
        };
    },
    methods: {
        handleIconClick(item) {
            console.log(
                `${item.type} 아이콘이 클릭되었지만 아직 처리 기능이 구현되지 않았습니다.`,
            );
            // 필요한 경우 추가 처리 로직 구현
        },
    },
};
</script>

<style scoped></style>
