import type { StudyChannelResponse } from '@/entities/channel/models/types';

export interface StudyState {
    // studyId별 listType별 마지막 활성 아이템
    activeItems: Record<string, Record<string, string>>;
    studyName: string;
    studyIcon: string;
    // studyId → API로 불러온 전체 StudyChannelResponse 를 캐싱
    channelsCache: Record<string, StudyChannelResponse>;
}

export const state = (): StudyState => ({
    activeItems: {},
    studyName: '',
    studyIcon: '',
    channelsCache: {},
});
