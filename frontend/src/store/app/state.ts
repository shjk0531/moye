export interface AppState {
    // 멤버 리스트 표시 여부
    isMemberListVisible: boolean;
}

export const state = (): AppState => ({
    isMemberListVisible: false,
});
