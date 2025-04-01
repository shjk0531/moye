<!-- src/features/sidebar/Sidebar.vue -->
<template>
    <div class="sidebar">
        <!-- 왼쪽 아이콘 리스트 (StudyIconList로 통합) -->
        <div class="icon-list">
            <!-- Title Icon: '/me'로 이동 -->
            <div class="">
                <div
                    class="flex flex-col items-center justify-start w-full gap-2"
                >
                    <div
                        class="relative flex items-center justify-center w-full group cursor-pointer"
                        @click="handleTitleIconClick"
                    >
                        <!-- indicator: 현재 경로가 '/me'면 상시 활성 -->
                        <div
                            :class="[
                                'absolute left-0 top-1/2 -translate-y-1/2 rounded transition-all duration-300',
                                isTitleActive()
                                    ? 'w-1 h-[70%] bg-gray-50'
                                    : 'w-0 group-hover:w-1 h-[50%] bg-gray-150',
                            ]"
                        ></div>
                        <div
                            class="w-10 h-10 mx-2 my-1 overflow-hidden border-2 rounded-lg"
                        >
                            <img
                                src="https://picsum.photos/200/300?random=1"
                                alt="title icon"
                                class="w-full h-full"
                            />
                        </div>
                    </div>
                </div>
            </div>

            <!-- 스터디 아이콘 리스트 -->
            <StudyIconList />
        </div>

        <!-- 채널 영역 -->
        <div
            class="study-channel-detail rounded-tl-lg border border-t-gray-600 border-l-gray-600"
        >
            <div class="notice panel text-white p-2">
                <p>공지사항</p>
            </div>
            <!-- 채널 목록 컴포넌트 -->
            <ChannelList />
        </div>
    </div>

    <!-- 사용자 패널 -->
    <div class="user-panel pb-4 px-2 border-rounded">
        <UserPanel />
    </div>
</template>

<script>
import { StudyIconList } from '@/entities/study';
import { ChannelList } from '@/entities/channel';
import { UserPanel } from '@/widgets/userPanel';

export default {
    name: 'Sidebar',
    components: {
        StudyIconList,
        ChannelList,
        UserPanel,
    },
    methods: {
        handleTitleIconClick() {
            // '/me'로 이동
            this.$router.push({ path: '/me' });
        },
        isTitleActive() {
            // 현재 경로가 '/me'인지 확인
            return this.$route.path === '/me';
        },
    },
};
</script>

<style scoped>
.sidebar {
    display: grid;
    grid-template-rows: subgrid;
    grid-template-columns: subgrid;
    grid-column: start / channelEnd;
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

.study-channel-detail {
    display: grid;
    grid-template-rows: subgrid;
    grid-template-columns: subgrid;
    grid-column: studiesEnd / channelEnd;
    grid-row: titlebarEnd / contentEnd;
}

.channel-list {
    grid-area: chatRoomList;
    width: var(--custom-chat-room-list-width);
    flex: 1;
    overflow-y: auto;
}
.user-panel {
    grid-area: userPanel;
}
</style>
