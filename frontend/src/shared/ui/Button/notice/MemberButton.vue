<template>
    <span :class="iconClasses" @click="handleClick" :title="title"></span>
</template>

<script>
export default {
    name: 'MemberButton',
    data() {
        return {
            // vuex에서 멤버 리스트 보여줄지 안 보여줄지 로드
            isMemberListVisible: this.$store.state.isMemberListVisible,
        };
    },
    computed: {
        iconClasses() {
            const baseClasses = `mdi mdi-account-group hover:text-gray-200 cursor-pointer`;
            return `${baseClasses} ${
                this.isMemberListVisible ? 'text-gray-150' : 'text-gray-400'
            }`;
        },
        title() {
            return '멤버';
        },
    },
    methods: {
        handleClick() {
            // vuex에 멤버 리스트 보여줄지 안보여줄지 저장(toggleMemberList)
            this.$store.commit('toggleMemberList');
        },
    },
};
</script>

<style scoped>
/* 필요한 스타일 추가 */
</style>
