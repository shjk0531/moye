<template>
    <span :class="iconClasses" @click="handleClick" :title="title"></span>
</template>

<script>
import {
    navigateToItem,
    isItemActive,
} from '@/widgets/notice/services/navigationService';

export default {
    name: 'ChannelButton',
    data() {
        return {
            // url에 따라 활성화 여부 결정
            // true: /study/:studyId/channel
            // false: /study/:studyId/다른 페이지
            isActive: this.$route.path.includes('/channel'),
        };
    },
    computed: {
        iconClasses() {
            const baseClasses = `mdi mdi-message hover:text-gray-200 cursor-pointer`;
            return `${baseClasses} ${
                this.isActive ? 'text-gray-150' : 'text-gray-400'
            }`;
        },
        title() {
            return '채널';
        },
    },
    methods: {
        async handleClick() {
            const store = this.$store;
            const router = this.$router;
            const studyId = this.$route.params.studyId;

            if (!studyId) {
                console.error('유효한 studyId가 없습니다.');
                return;
            }

            await navigateToItem('channel', studyId, store, router);
        },
    },
};
</script>

<style scoped>
/* 필요한 스타일 추가 */
</style>
