import type { LoungeState } from './state';

export const getters = {
    // loungeId, listType으로 저장된 활성 itemId 조회
    getActiveItemByLoungeAndType:
        (state: LoungeState) =>
        (loungeId: string, listType: string): string | null => {
            const activeItem = state.activeItems[loungeId];
            if (!activeItem) return null;
            if (activeItem.listType === listType) {
                return activeItem.itemId;
            }
            return null;
        },

    // 현재 라운지 이름/아이콘
    getCurrentLoungeInfo: (state: LoungeState) => ({
        id: state.currentLoungeInfo.id,
        name: state.currentLoungeInfo.name,
        icon: state.currentLoungeInfo.icon,
        type: state.currentLoungeInfo.type,
        channels: state.currentLoungeInfo.channels,
    }),
};
