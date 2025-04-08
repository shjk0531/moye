<template>
    <div class="grid-container bg-gray-950 h-100vh pt-3">
        <!-- 타이틀바 -->
        <!-- <TitleBar /> -->

        <!-- 스터디 아이콘 리스트 -->
        <div class="icon-list">
            <StudyListSidebar />
        </div>

        <!-- 사이드바 -->
        <div class="left-sidebar">
            <component :is="leftSideComponent" v-if="leftSideComponent" />
        </div>

        <!-- 사용자 패널 -->
        <div class="user-panel">
            <UserPanel />
        </div>

        <!-- 본문 -->
        <div class="page bg-gray-800">
            <router-view />
        </div>

        <!-- 멤버 리스트 -->
        <div class="right-sidebar">
            <MemberList class="w-50" v-if="isMemberListVisible" />
        </div>
    </div>
</template>

<script>
import { StudyNotice } from '@/widgets/notice';
import { MemberList, StudyListSidebar } from '@/widgets/sidebar';
import { TitleBar } from '@/widgets/titlebar';
import { UserPanel } from '@/widgets/userPanel';
import { mapState, mapMutations } from 'vuex';

export default {
    name: 'AppLayout',
    components: {
        TitleBar,
        StudyListSidebar,
        UserPanel,
        MemberList,
        StudyNotice,
    },
    computed: {
        ...mapState(['isMemberListVisible']),
        leftSideComponent() {
            return this.$route.meta.leftSide;
        },
        noticeComponent() {
            return this.$route.meta.notice;
        },
    },
    methods: {
        ...mapMutations(['toggleMemberList']),
    },
};
</script>

<style scoped lang="scss">
.grid-container {
    display: grid;
    grid-template-columns:
        [start] min-content
        [studiesEnd] min-content
        [channelEnd] 1fr
        [pageEnd] min-content
        [end];
    grid-template-rows:
        [top] min-content
        [titlebarEnd] min-content
        [noticeEnd] 1fr
        [contentEnd] min-content
        [end];
    grid-template-areas:
        'titlebar     titlebar      titlebar    titlebar'
        'notice       notice        notice      notice'
        'studyList    chatRoomList  page        memberList'
        'userPanel    userPanel     page        memberList';
    position: relative;
    overflow: hidden;
    height: 100vh;
}

.left-sidebar {
    display: grid;
    grid-template-rows: subgrid;
    grid-template-columns: subgrid;
    grid-column: studiesEnd / channelEnd;
    grid-row: titlebarEnd / contentEnd;
}

.page {
    display: grid;
    grid-template-rows: subgrid;
    grid-template-columns: subgrid;
    grid-column: channelEnd / end;
    grid-row: titlebarEnd / end;
}

.icon-list {
    display: grid;
    grid-template-rows: subgrid;
    grid-template-columns: subgrid;
    grid-column: start / studiesEnd;
    grid-row: titlebarEnd / contentEnd;
    width: var(--custom-icon-list-width);
}

.user-panel {
    grid-area: userPanel;
}
.right-sidebar {
    grid-area: memberList;
}
</style>
