<template>
    <div class="grid-container bg-gray-950 h-100vh">
        <!-- 타이틀바 -->
        <TitleBar />

        <!-- 스터디 아이콘 리스트 -->
        <div class="icon-list">
            <StudyListSidebar />
        </div>

        <!-- 사이드바 -->
        <div class="left-sidebar">
            <router-view name="leftSide" />
        </div>

        <!-- 사용자 패널 -->
        <div class="user-panel pb-4 px-2 border-rounded">
            <UserPanel />
        </div>

        <!-- 공지사항 (isMemberListVisible prop 전달 및 이벤트 리스너 추가) -->
        <Notice
            :isMemberListVisible="showMemberList"
            @toggle-member-list="toggleMemberList"
        />

        <!-- 본문 -->
        <div class="page bg-gray-800">
            <router-view name="page" />
        </div>

        <!-- 멤버 리스트 (showMemberList 값에 따라 렌더링) -->
        <MemberList v-if="showMemberList" />
    </div>
</template>

<script>
import { Notice } from '@/widgets/notice';
import { MemberList, StudyListSidebar } from '@/widgets/sidebar';
import { TitleBar } from '@/widgets/titlebar';
import { UserPanel } from '@/widgets/userPanel';

export default {
    name: 'AppLayout',
    components: {
        TitleBar,
        StudyListSidebar,
        UserPanel,
        MemberList,
        Notice,
    },
    data() {
        return {
            showChatPage: true, // 채팅 페이지 표시 여부
            showMemberList: false, // 멤버 리스트 표시 여부
        };
    },
    methods: {
        toggleChatPage() {
            // 채팅 페이지의 표시 상태를 토글함
            this.showChatPage = !this.showChatPage;
        },
        toggleMemberList() {
            // account-group 아이콘 클릭 시 멤버 리스트의 표시 상태를 토글함
            this.showMemberList = !this.showMemberList;
        },
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
    grid-column: channelEnd / pageEnd;
    grid-row: noticeEnd / end;
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
</style>
