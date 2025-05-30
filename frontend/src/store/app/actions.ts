// src/store/app/actions.ts

import type { AppState } from './state';

export const actions = {
    toggleMemberList(this: AppState): void {
        this.isMemberListVisible = !this.isMemberListVisible;
    },

    toggleMic(this: AppState): void {
        const micOn = this.micOn;
        this.micOn = !micOn;
    },

    toggleHeadset(this: AppState): void {
        const headsetOn = this.headsetOn;
        this.headsetOn = !headsetOn;
    },

    setLastRoute(this: AppState, route: string): void {
        this.lastRoute = route;
    },
};
