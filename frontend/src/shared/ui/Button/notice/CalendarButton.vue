<template>
    <span :class="iconClasses" @click="handleClick" :title="title"></span>
</template>

<script>
import {
    navigateToItem,
    isItemActive,
} from '@/widgets/notice/services/navigationService';
import { useRouter, useRoute } from 'vue-router';

export default {
    name: 'CalendarButton',
    setup() {
        const router = useRouter();
        const route = useRoute();
        const isActive = route.path.includes('/calendar');

        return { router, route, isActive };
    },
    computed: {
        iconClasses() {
            const baseClasses = `mdi mdi-calendar-blank hover:text-gray-200 cursor-pointer`;
            return `${baseClasses} ${
                this.isActive ? 'text-gray-150' : 'text-gray-400'
            }`;
        },
        title() {
            return '캘린더';
        },
    },
    methods: {
        async handleClick() {
            const studyId = this.route.params.studyId;

            if (!studyId) {
                console.error('유효한 studyId가 없습니다.');
                return;
            }

            await navigateToItem('calendar', studyId, this.router);
        },
    },
};
</script>

<style scoped>
/* 필요한 스타일 추가 */
</style>
