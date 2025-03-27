<!-- src/shared/layout/components/AppLayout.vue -->
<template>
    <div class="grid-container">
        <!-- 상단바 -->
        <div class="titleBar">
            <p>상단바</p>
        </div>

        <!-- 사이드바 -->
        <div class="sideBar">
            <div class="sideList">
                <nav class="iconList">
                    <div class="titleIcon">
                        <li class="nav-item">
                            <p>타이틀틀아이콘</p>
                        </li>
                    </div>
                    <ul class="studyList">
                        <li class="nav-item">
                            <p>스터디리스트</p>
                        </li>
                    </ul>
                </nav>
                <div class="studyChats">
                    <div class="notice">
                        <p class="font-bold p-2">공지사항</p>
                    </div>
                    <div class="chatRoomList">
                        <p>채팅방리스트</p>
                    </div>
                </div>
            </div>
            <div class="userPannel">
                <p>유저패널</p>
            </div>
        </div>

        <!-- 본문 -->
        <div class="contents">
            <div class="notice">
                <p>공지사항</p>
            </div>
            <div class="page">
                <div class="chatmessages">
                    <p>본문</p>
                </div>
                <div class="chat">
                    <p>채팅입력</p>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'AppLayout',
};
</script>

<style scoped lang="scss">
.grid-container {
    display: grid;
    grid-template-columns:
        [start] min-content
        [studiesEnd] min-content
        [chatRoomsEnd] 1fr
        [end];
    grid-template-rows:
        [top] var(--custom-app-top-bar-height)
        [titleBarEnd] min-content
        [noticeEnd] 1fr
        [contentEnd] var(--custom-app-bottom-pannel-height)
        [end];
    grid-template-areas:
        '.            titleBar      titleBar'
        'titleIcon    notice        notice'
        'studyList    chatRoomList  page'
        'userPannel   userPannel    page';
    position: relative;
    overflow: hidden;
    height: 100vh;
}

/* 각 영역에 grid-area 지정 */
.titleBar {
    grid-area: titleBar;
}

.sideBar {
    display: grid;
    grid-template-rows: subgrid;
    grid-template-columns: subgrid;
    grid-column: start / chatRoomsEnd;
    grid-row: titleBarEnd / end;
}

.sideList {
    display: grid;
    grid-template-rows: subgrid;
    grid-template-columns: subgrid;
    grid-column: start / chatRoomsEnd;
    grid-row: titleBarEnd / contentEnd;
}

.iconList {
    display: grid;
    grid-template-rows: subgrid;
    grid-template-columns: subgrid;
    grid-column: start / studiesEnd;
    grid-row: titleBarEnd / contentEnd;
}

.studyChats {
    display: grid;
    grid-template-rows: subgrid;
    grid-template-columns: subgrid;
    grid-column: studiesEnd / chatRoomsEnd;
    grid-row: titleBarEnd / contentEnd;

    .chatRoomList {
        grid-area: chatRoomList;
        width: var(--custom-chat-room-list-width);
    }
}

.contents {
    display: grid;
    grid-template-rows: subgrid;
    grid-template-columns: subgrid;
    grid-column: chatRoomsEnd / end;
    grid-row: titleBarEnd / end;
}

.page {
    grid-area: page;
    display: grid;
    grid-template-rows: subgrid;
    grid-template-columns: subgrid;
}

.notice {
    grid-area: notice;
}
.userPannel {
    grid-area: userPannel;
}

.titleIcon {
    grid-area: titleIcon;
}

.studyList {
    grid-area: studyList;
    // width: var(--custom-study-list-width);
}

.nav-item {
    width: 5rem;
    height: 5rem;
    background-color: #f0f0f0;
}
</style>
