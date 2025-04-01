// src/features/channel/api/channelsApi.ts
// 이 파일은 채널 관련 API 호출을 처리하는 모듈입니다.
// 실제 API 호출을 구현하는 대신 예시 데이터를 반환합니다.
// 이 예시에서는 그룹화된 채널과 비그룹화된 채널을 가져오는 두 개의 함수를 제공합니다.
// 이 파일은 TypeScript로 작성되었습니다.
export async function fetchChannelsGrouped() {
    // 실제 API 호출 대신 예시 데이터를 반환합니다.
    return [
        {
            id: 1,
            label: '그룹1',
            order: 1,
            items: [
                { id: 11, label: '채팅방 A', order: 2, icon: 'comments' },
                { id: 12, label: '채팅방 B', order: 1, icon: 'comment' },
            ],
        },
        {
            id: 2,
            label: '그룹2',
            order: 2,
            items: [
                { id: 21, label: '채팅방 C', order: 2, icon: 'microphone' },
                { id: 22, label: '채팅방 D', order: 1, icon: 'gamepad' },
            ],
        },
    ];
}

export async function fetchChannelsUngrouped() {
    return [
        { id: 3, label: '채팅1', order: 3, icon: 'bell' },
        { id: 4, label: '채팅2', order: 4, icon: 'newspaper' },
    ];
}
