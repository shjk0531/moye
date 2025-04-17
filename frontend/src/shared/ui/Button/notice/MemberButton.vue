<template>
    <span :class="iconClasses" @click="handleClick" :title="title"></span>
</template>

<script>
import { useAppStore } from '@/store';

export default {
    name: 'MemberButton',
    setup() {
        const appStore = useAppStore();
        return { appStore };
    },
    computed: {
        iconClasses() {
            const baseClasses = `mdi mdi-account-group hover:text-gray-200 cursor-pointer`;
            return `${baseClasses} ${
                this.appStore.isMemberListVisible
                    ? 'text-gray-150'
                    : 'text-gray-400'
            }`;
        },
        title() {
            return '멤버';
        },
    },
    methods: {
        handleClick() {
            this.appStore.toggleMemberList();
        },
    },
};
</script>

<style scoped>
/* 필요한 스타일 추가 */
</style>
