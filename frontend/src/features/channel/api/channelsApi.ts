// src/features/channel/api/channelsApi.ts
// 채널 관련 API 호출 모듈 (예시 데이터를 반환)

// 그룹 타입 정의
export interface ChannelGroup {
    id: number;
    label: string;
    order: number;
}

// 채널 타입 정의
export interface Channel {
    id: number;
    label: string;
    order: number;
    icon: string;
    groupId: number | null;
}

export async function fetchChannelsGrouped(): Promise<ChannelGroup[]> {
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

export async function fetchChannelsUngrouped(): Promise<Channel[]> {
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

/**
 * 그룹 및 채널 목록에서 첫 번째 채널을 찾는 함수
 * 1. 먼저 그룹 및 채널들을 order 기준으로 정렬
 * 2. 그룹이 있다면 가장 첫 번째 그룹의 가장 첫 번째 채널 반환
 * 3. 그룹이 없거나 그룹에 채널이 없다면 비그룹 채널 중 첫 번째 채널 반환
 * @param groups 그룹 목록
 * @param channels 채널 목록
 * @returns 첫 번째 채널 ID 또는 null
 */
export function findFirstChannel(
    groups: ChannelGroup[],
    channels: Channel[],
): number | null {
    if (!channels || channels.length === 0) {
        return null;
    }

    // 그룹 및 채널을 order 기준으로 정렬
    const sortedGroups = [...groups].sort(
        (a, b) => (a.order || 0) - (b.order || 0),
    );
    const sortedChannels = [...channels].sort(
        (a, b) => (a.order || 0) - (b.order || 0),
    );

    // 그룹이 있는 경우
    if (sortedGroups.length > 0) {
        // 가장 첫 번째 그룹 찾기
        const firstGroup = sortedGroups[0];

        // 첫 번째 그룹에 속한 채널들 찾기
        const channelsInFirstGroup = sortedChannels
            .filter((channel) => channel.groupId === firstGroup.id)
            .sort((a, b) => (a.order || 0) - (b.order || 0));

        // 첫 번째 그룹에 채널이 있으면 그 중 첫 번째 채널 반환
        if (channelsInFirstGroup.length > 0) {
            return channelsInFirstGroup[0].id;
        }
    }

    // 그룹이 없거나 첫 번째 그룹에 채널이 없는 경우, 비그룹 채널 중 첫 번째 채널 찾기
    const ungroupedChannels = sortedChannels
        .filter((channel) => channel.groupId === null)
        .sort((a, b) => (a.order || 0) - (b.order || 0));

    if (ungroupedChannels.length > 0) {
        return ungroupedChannels[0].id;
    }

    // 어떤 채널도 없는 경우 null 반환
    return null;
}
