// src/store/index.ts
import { createStore } from 'vuex';

const store = createStore({
    state: () => ({
        // studyId별 각 목록(listType별)의 마지막 활성 항목을 저장합니다.
        // 예: { '1': { channel: '채널1의 id', calendar: '캘린더2의 id' } }
        activeItems: {} as Record<string, Record<string, string>>,
        studyName: '',
        studyIcon: '',
    }),
    mutations: {
        setActiveItem(
            state,
            {
                studyId,
                listType,
                itemId,
            }: { studyId: string; listType: string; itemId: string },
        ) {
            if (!state.activeItems[studyId]) {
                state.activeItems[studyId] = {};
            }
            state.activeItems[studyId][listType] = itemId;
        },
        setStudyName(state, name: string) {
            state.studyName = name;
        },
        setStudyIcon(state, icon: string) {
            state.studyIcon = icon;
        },
    },
});

export default store;
