import { apiClient } from '@/shared/api';
import { findFirstItem } from '@/shared/lib/panelUtils';
import type { ChannelGroup, Channel } from '../models/types';

/**
 * 채널 그룹 목록을 가져옵니다.
 */
export async function fetchChannelsGrouped(): Promise<ChannelGroup[]> {
    try {
        // 실제 환경에서는 아래 API 호출을 사용
        // const response = await apiClient.get('/channels/groups');
        // return response.data;

        // 예시 데이터를 반환합니다.
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
    } catch (error) {
        console.error('채널 그룹을 가져오는 중 오류 발생:', error);
        throw error;
    }
}

/**
 * 채널 목록을 가져옵니다.
 */
export async function fetchChannelsUngrouped(): Promise<Channel[]> {
    try {
        // 실제 환경에서는 아래 API 호출을 사용
        // const response = await apiClient.get('/channels');
        // return response.data;

        // 예시 데이터를 반환합니다.
        return [
            { id: 0, label: '채팅1', order: 1, icon: 'comments', groupId: 1 },
            { id: 1, label: '채팅2', order: 2, icon: 'newspaper', groupId: 1 },
            { id: 2, label: '채팅3', order: 1, icon: 'comment', groupId: null },
            {
                id: 3,
                label: '채팅4',
                order: 4,
                icon: 'microphone',
                groupId: null,
            },
            { id: 4, label: '채팅5', order: 5, icon: 'gamepad', groupId: null },
            { id: 5, label: '채팅6', order: 6, icon: 'bell', groupId: null },
            { id: 6, label: '채팅7', order: 1, icon: 'newspaper', groupId: 2 },
            { id: 7, label: '채팅8', order: 2, icon: 'comments', groupId: 2 },
            { id: 8, label: '채팅9', order: 3, icon: 'comment', groupId: 2 },
            {
                id: 9,
                label: '채팅10',
                order: 4,
                icon: 'microphone',
                groupId: 2,
            },
            { id: 10, label: '채팅11', order: 5, icon: 'gamepad', groupId: 2 },
        ];
    } catch (error) {
        console.error('채널 목록을 가져오는 중 오류 발생:', error);
        throw error;
    }
}

/**
 * 그룹 및 채널 목록에서 첫 번째 채널을 찾는 함수
 * 내부적으로 공통 findFirstItem 함수를 사용합니다.
 *
 * @param groups 그룹 목록
 * @param channels 채널 목록
 * @returns 첫 번째 채널 ID 또는 null
 */
export function findFirstChannel(
    groups: ChannelGroup[],
    channels: Channel[],
): number | null {
    return findFirstItem(groups, channels);
}
