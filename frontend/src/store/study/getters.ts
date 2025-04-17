import type { StudyState } from './state';

export const getters = {
    // 스터디 ID별 활성 항목을 가져오는 게터
    getActiveItemByStudyAndType:
        (state: StudyState) =>
        (studyId: string, listType: string): string | null => {
            return state.activeItems[studyId]?.[listType] ?? null;
        },

    // 현재 스터디 정보 게터
    currentStudyInfo: (state: StudyState) => {
        return {
            name: state.studyName,
            icon: state.studyIcon,
        };
    },
};
