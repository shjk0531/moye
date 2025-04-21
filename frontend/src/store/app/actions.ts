// src/store/app/actions.ts

import type { AppState } from './state';

export const actions = {
    toggleMemberList(this: AppState) {
        this.isMemberListVisible = !this.isMemberListVisible;
    },

    toggleMic(this: AppState) {
        const micOn = this.micOn;
        this.micOn = !micOn;
    },

    toggleHeadset(this: AppState) {
        const headsetOn = this.headsetOn;
        this.headsetOn = !headsetOn;
    },

    saveLastRoute(this: AppState, route: string) {
        this.lastRoute = route;
    },
};
