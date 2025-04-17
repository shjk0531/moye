import type { StudyState } from './state';

export interface SetActiveItemPayload {
    studyId: string;
    listType: string;
    itemId: string;
}

export const actions = {
    setActiveItem(this: StudyState, payload: SetActiveItemPayload) {
        const { studyId, listType, itemId } = payload;
        if (!this.activeItems[studyId]) {
            this.activeItems[studyId] = {};
        }
        this.activeItems[studyId][listType] = itemId;
    },

    setStudyName(this: StudyState, name: string) {
        this.studyName = name;
    },

    setStudyIcon(this: StudyState, icon: string) {
        this.studyIcon = icon;
    },
};
