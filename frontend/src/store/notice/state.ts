// src/store/notice/state.ts

export interface NoticeState {
    memberListVisible: boolean;
}

export const state = (): NoticeState => ({
    memberListVisible: false,
});
