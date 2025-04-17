// src/store/app/state.ts

export interface AppState {
    // 멤버 리스트 표시 여부
    isMemberListVisible: boolean;

    // 마이크 켜짐 여부
    micOn: boolean;

    // 헤드셋 켜짐 여부
    headsetOn: boolean;
}

export const state = (): AppState => ({
    isMemberListVisible: false,
    micOn: true,
    headsetOn: true,
});
