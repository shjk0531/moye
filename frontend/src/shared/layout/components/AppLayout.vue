<template>
    <div class="grid-container bg-gray-950 h-100vh">
        <!-- 타이틀바 -->
        <TitleBar />

        <!-- 사이드바 영역: 이름이 지정된 router-view -->
        <router-view name="sidebar">
            <!-- fallback: 라우트에서 sidebar가 지정되지 않으면 기본 Sidebar 컴포넌트를 사용 -->
            <Sidebar />
        </router-view>

        <!-- 공지사항 -->
        <Notice
            :isMemberListVisible="showMemberList"
            @toggle-member-list="toggleMemberList"
        />

        <!-- 본문 영역: 이름이 지정된 router-view -->
        <div class="page">
            <router-view name="page" />
        </div>

        <!-- 멤버 리스트 (showMemberList 값에 따라 렌더링) -->
        <MemberList v-if="showMemberList" />
    </div>
</template>

<script lang="ts">
import { StudyIconList } from '@/entities/study';
import { Notice } from '@/widgets/notice';
import { MemberList, Sidebar } from '@/widgets/sidebar';
import { TitleBar } from '@/widgets/titlebar';
import { UserPanel } from '@/widgets/userPanel';

export default {
    name: 'AppLayout',
    components: {
        Sidebar,
        TitleBar,
        StudyIconList,
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

.page {
    display: grid;
    grid-template-rows: subgrid;
    grid-template-columns: subgrid;
    grid-column: channelEnd / pageEnd;
    grid-row: noticeEnd / end;
}
</style>
