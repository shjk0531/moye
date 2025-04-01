<!-- src/shared/layout/components/AppLayout.vue -->
<template>
    <div class="grid-container bg-gray-950 h-100vh">
        <!-- 타이틀바 -->
        <TitleBar />

        <!-- 사이드바 -->
        <Sidebar />

        <!-- 본문 -->
        <div class="page">
            <router-view />
        </div>
    </div>
</template>

<script>
import { StudyIconList } from '@/entities/study';
import { Sidebar } from '@/widgets/sidebar';
import { TitleBar } from '@/widgets/titlebar';
import { UserPanel } from '@/widgets/userPanel';

export default {
    name: 'AppLayout',
    components: {
        Sidebar,
        TitleBar,
        StudyIconList,
        UserPanel,
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
        [end];
    grid-template-rows:
        [top]
        min-content
        [titlebarEnd]
        min-content
        [noticeEnd]
        1fr
        [contentEnd]
        min-content
        [end];
    grid-template-areas:
        'titlebar     titlebar      titlebar'
        'notice       notice        notice'
        'studyList    chatRoomList  page'
        'userPanel    userPanel     page';
    position: relative;
    overflow: hidden;
    height: 100vh;
}

.page {
    display: grid;
    grid-template-rows: subgrid;
    grid-template-columns: subgrid;
    grid-column: channelEnd / end;
    grid-row: titlebarEnd / end;
}

.notice {
    grid-area: notice;
}

.user-panel {
    grid-area: userPanel;
}
</style>
