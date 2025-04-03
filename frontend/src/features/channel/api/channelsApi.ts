// src/features/channel/api/channelsApi.ts
// 채널 관련 API 호출 모듈 (예시 데이터를 반환)
export async function fetchChannelsGrouped() {
    return [
        {
            id: 1,
            label: '그룹1',
            order: 2,
        },
        {
            id: 2,
            label: '그룹2',
            order: 3,
        },
    ];
}

export async function fetchChannelsUngrouped() {
    return [
        { id: 0, label: '채팅1', order: 1, icon: 'comments', groupId: 1 },
        { id: 1, label: '채팅2', order: 2, icon: 'newspaper', groupId: 1 },
        { id: 2, label: '채팅3', order: 1, icon: 'comment', groupId: null },
        { id: 3, label: '채팅4', order: 4, icon: 'microphone', groupId: null },
        { id: 4, label: '채팅5', order: 5, icon: 'gamepad', groupId: null },
        { id: 5, label: '채팅6', order: 6, icon: 'bell', groupId: null },
        { id: 6, label: '채팅7', order: 1, icon: 'newspaper', groupId: 2 },
        { id: 7, label: '채팅8', order: 2, icon: 'comments', groupId: 2 },
        { id: 8, label: '채팅9', order: 3, icon: 'comment', groupId: 2 },
        { id: 9, label: '채팅10', order: 4, icon: 'microphone', groupId: 2 },
        { id: 10, label: '채팅11', order: 5, icon: 'gamepad', groupId: 2 },
    ];
}
