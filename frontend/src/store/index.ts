// store/index.js 또는 store/index.ts
import { createStore } from 'vuex';

const store = createStore({
    state: () => ({
        activeChannelMap: {} as Record<string, string>,
    }),
    mutations: {
        setActiveChannel(state, { studyId, channelId }) {
            state.activeChannelMap[studyId] = channelId;
        },
    },
});

export default store;
