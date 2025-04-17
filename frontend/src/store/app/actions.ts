import type { AppState } from './state';

export const actions = {
    toggleMemberList(this: AppState) {
        this.isMemberListVisible = !this.isMemberListVisible;
    },
};
