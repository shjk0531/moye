// src/store/notice/actions.ts

import type { NoticeState } from './state';

export const actions = {
    toggleMemberList(this: NoticeState) {
        this.memberListVisible = !this.memberListVisible;
    },
};
