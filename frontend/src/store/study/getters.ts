import type { StudyState } from './state';

export const getters = {
    // studyId, listType으로 저장된 활성 itemId 조회
    getActiveItemByStudyAndType:
        (state: StudyState) =>
        (studyId: string, listType: string): string | null => {
            const activeItem = state.activeItems[studyId];
            if (!activeItem) return null;
            if (activeItem.listType === listType) {
                return activeItem.itemId;
            }
            return null;
        },

    // 현재 스터디 이름/아이콘
    getCurrentStudyInfo: (state: StudyState) => ({
        id: state.currentStudyInfo.id,
        name: state.currentStudyInfo.name,
        icon: state.currentStudyInfo.icon,
        type: state.currentStudyInfo.type,
        channels: state.currentStudyInfo.channels,
    }),
};
