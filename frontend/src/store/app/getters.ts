// src/store/app/getters.ts
import type { AppState } from './state';

export const getters = {
    isMicOn: (state: AppState) => state.micOn,
    isHeadsetOn: (state: AppState) => state.headsetOn,

    getLastRoute: (state: AppState) => state.lastRoute,
};
