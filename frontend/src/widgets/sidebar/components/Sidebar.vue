<template>
    <div class="sidebar">
        <!-- 왼쪽 아이콘 리스트 -->
        <div class="icon-list">
            <div
                class="title-icon w-10 h-10 mx-2 my-1 overflow-hidden cursor-pointer border-2 border-transparent hover:border-blue-500 rounded-lg"
            >
                <img
                    src="https://picsum.photos/200/300?random=1"
                    alt="title icon"
                    class="w-full h-full"
                />
            </div>
            <!-- 구분선 -->
            <div class="divider w-full border-t border-gray-600"></div>
            <div class="study-list">
                <!-- 서버 아이콘 목록 -->
                <div
                    class="study-item w-10 h-10 mx-2 my-0.5 overflow-hidden cursor-pointer border-2 border-transparent hover:border-blue-500 rounded-lg"
                    v-for="study in studies"
                    :key="study.id"
                >
                    <img
                        :src="study.icon"
                        alt="study icon"
                        class="w-full h-full"
                    />
                </div>
            </div>
        </div>

        <!-- 채널 영역 -->
        <div
            class="study-channel-detail rounded-tl-lg border border-t-gray-600 border-l-gray-600"
        >
            <div class="notice panel text-white p-2">
                <p>공지사항</p>
            </div>
            <div class="channel-list">
                <!-- 그룹 채널과 비그룹 채널을 order에 따라 병합하여 렌더링 -->
                <div class="merged-channels">
                    <div v-for="channel in mergedChannels" :key="channel.id">
                        <PanelSection
                            v-if="channel.type === 'group'"
                            :label="channel.label"
                            :items="channel.items"
                            :activeChannelId="activeChannelId"
                        />
                        <PanelItem
                            v-else
                            :item="channel"
                            :isActive="channel.id === activeChannelId"
                        />
                    </div>
                </div>
            </div>
        </div>

        <!-- 사용자 패널 -->
        <div class="user-panel mb-4 mx-2 border-rounded">
            <UserPanel />
        </div>
    </div>
</template>

<script>
import { PanelSection, PanelItem } from '@/shared/panel';
import UserPanel from './UserPanel.vue';

export default {
    name: 'Sidebar',
    components: {
        PanelSection,
        PanelItem,
        UserPanel,
    },
    data() {
        return {
            studies: [
                { id: 1, icon: 'https://picsum.photos/200/300?random=2' },
                { id: 2, icon: 'https://picsum.photos/200/300?random=3' },
                { id: 3, icon: 'https://picsum.photos/200/300?random=4' },
            ],
            // 그룹에 속하는 채널 (PanelSection으로 렌더링)
            channelsGrouped: [
                {
                    id: 1,
                    label: '그룹1',
                    order: 1,
                    items: [
                        {
                            id: 11,
                            label: '채팅방 A',
                            order: 2,
                            icon: 'comments',
                        },
                        {
                            id: 12,
                            label: '채팅방 B',
                            order: 1,
                            icon: 'comment',
                        },
                    ],
                },
                {
                    id: 2,
                    label: '그룹2',
                    order: 2,
                    items: [
                        {
                            id: 21,
                            label: '채팅방 C',
                            order: 2,
                            icon: 'microphone',
                        },
                        {
                            id: 22,
                            label: '채팅방 D',
                            order: 1,
                            icon: 'gamepad',
                        },
                    ],
                },
            ],
            // 그룹에 속하지 않는 채널 (PanelItem으로 렌더링)
            channelsUngrouped: [
                { id: 3, label: '채팅1', order: 3, icon: 'bell' },
                { id: 4, label: '채팅2', order: 4, icon: 'newspaper' },
            ],
            // 현재 사용자가 들어가 있는 채팅방 id (예시로 채팅방 B가 active)
            activeChannelId: 11,
        };
    },
    computed: {
        // 그룹 채널과 비그룹 채널을 병합하고 order에 따라 정렬합니다.
        mergedChannels() {
            const groups = this.channelsGrouped.map((group) => ({
                ...group,
                type: 'group',
            }));
            const ungrouped = this.channelsUngrouped.map((channel) => ({
                ...channel,
                type: 'ungrouped',
            }));
            // order 값이 없으면 9999로 처리하여 항상 뒤로 배치합니다.
            return [...groups, ...ungrouped].sort(
                (a, b) => (a.order || 9999) - (b.order || 9999),
            );
        },
    },
};
</script>

<style scoped>
.sidebar {
    display: grid;
    grid-template-rows: subgrid;
    grid-template-columns: subgrid;
    grid-column: start / chatRoomsEnd;
    grid-row: titleBarEnd / end;
}

.icon-list {
    display: grid;
    grid-template-rows: subgrid;
    grid-template-columns: subgrid;
    grid-column: start / studiesEnd;
    grid-row: titleBarEnd / contentEnd;
}

.study-channel-detail {
    display: grid;
    grid-template-rows: subgrid;
    grid-template-columns: subgrid;
    grid-column: studiesEnd / channelEnd;
    grid-row: titleBarEnd / contentEnd;
}

.study-list {
    display: flex;
    flex-direction: column;
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
