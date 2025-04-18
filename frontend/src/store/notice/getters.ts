// src/store/notice/getters.ts
import type { NoticeState } from './state';

export const getters = {
    isMemberListVisible: (state: NoticeState) => state.memberListVisible,
};
