// src/shared/lib/panelUtils.ts
// 패널 관련 공통 유틸리티 함수들

/**
 * 패널에서 사용되는 그룹 인터페이스
 */
export interface PanelGroup {
    id: number;
    label: string;
    order: number;
}

/**
 * 패널에서 사용되는 아이템 인터페이스
 */
export interface PanelItem {
    id: number;
    label: string;
    order: number;
    icon?: string;
    groupId: number | null;
}

/**
 * 그룹 및 아이템 목록에서 첫 번째 아이템을 찾는 범용 함수
 *
 * 이 함수는 패널 리스트의 모든 종류(채널, 캘린더 등)에서 재사용할 수 있습니다.
 *
 * @param groups 그룹 목록
 * @param items 아이템 목록
 * @returns 첫 번째 아이템 ID 또는 찾지 못한 경우 null
 */
export function findFirstItem<G extends PanelGroup, I extends PanelItem>(
    groups: G[],
    items: I[],
): number | null {
    if (!items || items.length === 0) {
        return null;
    }

    // 그룹 및 아이템을 order 기준으로 정렬
    const sortedGroups = [...groups].sort(
        (a, b) => (a.order || 0) - (b.order || 0),
    );
    const sortedItems = [...items].sort(
        (a, b) => (a.order || 0) - (b.order || 0),
    );

    // 그룹이 있는 경우
    if (sortedGroups.length > 0) {
        // 가장 첫 번째 그룹 찾기
        const firstGroup = sortedGroups[0];

        // 첫 번째 그룹에 속한 아이템들 찾기
        const itemsInFirstGroup = sortedItems
            .filter((item) => item.groupId === firstGroup.id)
            .sort((a, b) => (a.order || 0) - (b.order || 0));

        // 첫 번째 그룹에 아이템이 있으면 그 중 첫 번째 아이템 반환
        if (itemsInFirstGroup.length > 0) {
            return itemsInFirstGroup[0].id;
        }
    }

    // 그룹이 없거나 첫 번째 그룹에 아이템이 없는 경우, 비그룹 아이템 중 첫 번째 아이템 찾기
    const ungroupedItems = sortedItems
        .filter((item) => item.groupId === null)
        .sort((a, b) => (a.order || 0) - (b.order || 0));

    if (ungroupedItems.length > 0) {
        return ungroupedItems[0].id;
    }

    // 어떤 아이템도 없는 경우 null 반환
    return null;
}

/**
 * 그룹 및 아이템 데이터 구조 변환 함수
 *
 * 패널 리스트 컴포넌트에서 사용하는 형식으로 데이터를 변환합니다.
 *
 * @param groups 그룹 목록
 * @param items 아이템 목록
 * @returns 병합된 아이템 목록 (그룹 및 비그룹 항목 포함)
 */
export function mergePanelItems<G extends PanelGroup, I extends PanelItem>(
    groups: G[],
    items: I[],
): ((G & { type: 'group'; items: I[] }) | (I & { type: 'ungrouped' }))[] {
    // 그룹에 속한 아이템들을 해당 그룹 객체 안에 items 배열로 추가
    const mergedGroups = groups.map((group) => {
        const groupItems = items
            .filter((item) => item.groupId === group.id)
            .sort((a, b) => (a.order || 0) - (b.order || 0));
        return { ...group, type: 'group', items: groupItems } as G & {
            type: 'group';
            items: I[];
        };
    });

    // 그룹에 속하지 않은 일반 아이템
    const ungroupedItems = items
        .filter((item) => item.groupId === null)
        .map(
            (item) =>
                ({ ...item, type: 'ungrouped' } as I & { type: 'ungrouped' }),
        );

    // 전체를 order 값 기준으로 정렬
    return [...mergedGroups, ...ungroupedItems].sort(
        (a, b) => (a.order || 9999) - (b.order || 9999),
    );
}
