<!-- src/shared/layout/components/AppLayout.vue -->
<template>
    <div class="grid-container bg-gray-950 h-100vh">
        <!-- 타이틀바 -->
        <div class="titleBar flex flex-row items-center justify-center">
            <div class="flex flex-row items-center justify-center">
                <div class="titleIcon bg-gray-400 w-2 h-2"></div>
                <div class="text-gray-200 text-2xs">타이틀</div>
            </div>
        </div>

        <!-- 사이드바 -->
        <div class="sideBar">
            <div class="sideList">
                <nav class="iconList">
                    <div class="titleIcon">
                        <li class="w-8 h-8 bg-gray-400 rounded-full">
                            <p>타이틀아이콘</p>
                        </li>
                    </div>
                    <ul class="studyList">
                        <li class="nav-item">
                            <p>스터디리스트</p>
                        </li>
                    </ul>
                </nav>
                <div
                    class="studyChats border-l-1 border-t-1 border-gray-700 rounded-tl-lg"
                >
                    <div class="notice">
                        <p class="font-bold p-2 text-gray-50">공지사항</p>
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
        <div class="page">
            <router-view />
        </div>
    </div>
    <div class="flex flex-row">
        <div class="flex-1 bg-gray-950 text-white p-2">
            <p>950</p>
        </div>
        <div class="flex-1 bg-gray-900 text-white p-2">
            <p>900</p>
        </div>
        <div class="flex-1 bg-gray-850 text-white p-2">
            <p>850</p>
        </div>
        <div class="flex-1 bg-gray-800 text-white p-2">
            <p>800</p>
        </div>
        <div class="flex-1 bg-gray-750 text-white p-2">
            <p>750</p>
        </div>
        <div class="flex-1 bg-gray-700 text-white p-2">
            <p>700</p>
        </div>
        <div class="flex-1 bg-gray-650 text-white p-2">
            <p>650</p>
        </div>
        <div class="flex-1 bg-gray-600 text-white p-2">
            <p>600</p>
        </div>
        <div class="flex-1 bg-gray-550 text-white p-2">
            <p>550</p>
        </div>
        <div class="flex-1 bg-gray-500 text-white p-2">
            <p>500</p>
        </div>
        <div class="flex-1 bg-gray-450 text-white p-2">
            <p>450</p>
        </div>
        <div class="flex-1 bg-gray-400 text-white p-2">
            <p>400</p>
        </div>
        <div class="flex-1 bg-gray-350 text-white p-2">
            <p>350</p>
        </div>
        <div class="flex-1 bg-gray-300 text-white p-2">
            <p>300</p>
        </div>
        <div class="flex-1 bg-gray-250 text-white p-2">
            <p>250</p>
        </div>
        <div class="flex-1 bg-gray-200 text-white p-2">
            <p>200</p>
        </div>
        <div class="flex-1 bg-gray-150 text-white p-2">
            <p>150</p>
        </div>
        <div class="flex-1 bg-gray-100 text-white p-2">
            <p>100</p>
        </div>
        <div class="flex-1 bg-gray-50 text-white p-2">
            <p>50</p>
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
        [top]
        min-content
        [titleBarEnd]
        min-content
        [noticeEnd]
        1fr
        [contentEnd]
        min-content
        [end];
    grid-template-areas:
        'titleBar     titleBar      titleBar'
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

.page {
    display: grid;
    grid-template-rows: subgrid;
    grid-template-columns: subgrid;
    grid-column: chatRoomsEnd / end;
    grid-row: titleBarEnd / end;
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
</style>
