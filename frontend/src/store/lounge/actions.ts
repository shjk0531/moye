// src/store/lounge/actions.ts
import type { LoungeChannelResponse } from '@/entities/channel/models/types';
import { fetchChannels } from '@/entities';
import type { CurrentLoungeInfo } from './state';

export interface SetActiveItemPayload {
    loungeId: string;
    listType: string;
    itemId: string;
}

export const actions = {
    /** loungeId에 대한 채널 목록을 한 번만 호출해서 캐싱한 뒤 반환 */
    async loadChannels(
        this: any,
        loungeId: string,
    ): Promise<LoungeChannelResponse> {
        if (!this.channelsCache[loungeId]) {
            const response = await fetchChannels(loungeId);
            this.channelsCache[loungeId] = response;
            console.log('lounge store loadChannels response:', response);
        }
        return this.channelsCache[loungeId];
    },

    /** 활성 항목 저장 */
    setActiveItem(
        this: any,
        { loungeId, listType, itemId }: SetActiveItemPayload,
    ) {
        if (!this.activeItems[loungeId]) {
            this.activeItems[loungeId] = {};
        }
        this.activeItems[loungeId][listType] = itemId;
    },

    /** 로드된(또는 캐시된) 채널 목록에서 "첫 번째" 채널 ID 추출 */
    async getFirstChannelId(
        this: any,
        loungeId: string,
    ): Promise<string | null> {
        const { items } = await this.loadChannels(loungeId);

        console.log('items[0]', items[0]);
        const firstItem = items[0];
        if (firstItem.item_type === 'channel') {
            return firstItem.channel.id;
        } else if (firstItem.item_type === 'group') {
            return firstItem.group.channels[0].id;
        }
        return null;
    },

    /**
     * 현재 스터디 정보 설정
     */
    setCurrentLoungeInfo(this: any, loungeInfo: CurrentLoungeInfo) {
        this.currentLoungeInfo = {
            id: loungeInfo.id,
            name: loungeInfo.name,
            icon: loungeInfo.icon,
            type: loungeInfo.type,
            channels: loungeInfo.channels,
        };
    },
};
