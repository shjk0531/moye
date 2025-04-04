// src/store/index.ts
import { createStore } from 'vuex';

const store = createStore({
    state: () => ({
        activeChannelMap: {} as Record<string, string>,
        studyName: '',
        studyIcon: '',
    }),
    mutations: {
        setActiveChannel(state, { studyId, channelId }) {
            state.activeChannelMap[studyId] = channelId;
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
