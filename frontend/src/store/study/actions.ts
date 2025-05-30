// src/store/study/actions.ts
import type { StudyChannelResponse } from '@/entities/channel/models/types';
import { fetchStudyChannels } from '@/entities';
import type { CurrentStudyInfo } from './state';

export interface SetActiveItemPayload {
    studyId: string;
    listType: string;
    itemId: string;
}

export const actions = {
    /** studyId에 대한 채널 목록을 한 번만 호출해서 캐싱한 뒤 반환 */
    async loadChannels(
        this: any,
        studyId: string,
    ): Promise<StudyChannelResponse> {
        if (!this.channelsCache[studyId]) {
            const response = await fetchStudyChannels(studyId);
            this.channelsCache[studyId] = response;
            console.log('response:', response);
        }
        return this.channelsCache[studyId];
    },

    /** 활성 항목 저장 */
    setActiveItem(
        this: any,
        { studyId, listType, itemId }: SetActiveItemPayload,
    ) {
        if (!this.activeItems[studyId]) {
            this.activeItems[studyId] = {};
        }
        this.activeItems[studyId][listType] = itemId;
    },

    /** 로드된(또는 캐시된) 채널 목록에서 "첫 번째" 채널 ID 추출 */
    async getFirstChannelId(
        this: any,
        studyId: string,
    ): Promise<string | null> {
        const { items } = await this.loadChannels(studyId);

        // (1) 그룹 내부의 첫 채널 우선 — snake_case 사용
        const groupItem = (items as any[]).find((i) => i.item_type === 'group');
        if (groupItem?.group?.channels?.length) {
            return groupItem.group.channels[0].id;
        }

        // (2) 개별 channel 아이템 중 첫 번째
        const channelItem = (items as any[]).find(
            (i) => i.item_type === 'channel',
        );
        return channelItem?.channel?.id ?? null;
    },

    /**
     * 현재 스터디 정보 설정
     */
    setCurrentStudyInfo(this: any, studyInfo: CurrentStudyInfo) {
        this.currentStudyInfo = {
            id: studyInfo.id,
            name: studyInfo.name,
            icon: studyInfo.icon,
            type: studyInfo.type,
            channels: studyInfo.channels,
        };
    },
};
