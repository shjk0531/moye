import type { StudyChannelResponse } from '@/entities/channel/models/types';

export interface ActiveItem {
    studyId: string;
    listType: string;
    itemId: string;
}

export interface ActiveItems {
    [studyId: string]: ActiveItem;
}

export interface ChannelsCache {
    [studyId: string]: StudyChannelResponse;
}

export interface CurrentStudyInfo {
    id: string;
    name: string;
    icon: string;
    type: string;
    channels: StudyChannelResponse;
}

export interface StudyState {
    // studyId별 listType별 마지막 활성 아이템
    activeItems: ActiveItems;
    currentStudyInfo: CurrentStudyInfo;
    // studyId → API로 불러온 전체 StudyChannelResponse 를 캐싱
    channelsCache: ChannelsCache;
}

export const state = (): StudyState => ({
    activeItems: {},
    currentStudyInfo: {
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
