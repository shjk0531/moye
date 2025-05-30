import type { StudyState } from './state';

export const getters = {
    // studyId, listType으로 저장된 활성 itemId 조회
    getActiveItemByStudyAndType:
        (state: StudyState) =>
        (studyId: string, listType: string): string | null => {
            return state.activeItems[studyId]?.[listType] ?? null;
        },

    // 현재 스터디 이름/아이콘
    currentStudyInfo: (state: StudyState) => ({
        name: state.studyName,
        icon: state.studyIcon,
    }),
};
