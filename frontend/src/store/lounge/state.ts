import type { LoungeChannelResponse } from '@/entities/channel/models/types';

export interface ActiveItem {
    loungeId: string;
    listType: string;
    itemId: string;
}

export interface ActiveItems {
    [loungeId: string]: ActiveItem;
}

export interface ChannelsCache {
    [loungeId: string]: LoungeChannelResponse;
}

export interface CurrentLoungeInfo {
    id: string;
    name: string;
    icon: string;
    type: string;
    channels: LoungeChannelResponse;
}

export interface LoungeState {
    // loungeId별 listType별 마지막 활성 아이템
    activeItems: ActiveItems;
    currentLoungeInfo: CurrentLoungeInfo;
    // loungeId → API로 불러온 전체 LoungeChannelResponse 를 캐싱
    channelsCache: ChannelsCache;
}

export const state = (): LoungeState => ({
    activeItems: {},
    currentLoungeInfo: {
        id: '',
        name: '',
        icon: '',
        type: '',
        channels: {
            items: [],
        },
    },
    channelsCache: {},
});
