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
 * 비그룹 아이템과 그룹 내 아이템을 order 값에 따라 정렬하여 첫 번째 아이템을 찾습니다.
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

    // 그룹과 아이템 복사
    const allGroups = [...groups];
    const allItems = [...items];

    // 모든 아이템 항목 생성 (비그룹 아이템 + 그룹 내 첫 번째 아이템)
    const itemEntries: Array<{ id: number; order: number }> = [];

    // 비그룹 아이템 추가
    const ungroupedItems = allItems.filter((item) => item.groupId === null);
    ungroupedItems.forEach((item) => {
        itemEntries.push({
            id: item.id,
            order: item.order || 0,
        });
    });

    // 그룹 내 첫 번째 아이템 추가
    allGroups.forEach((group) => {
        // 해당 그룹의 아이템 찾기
        const groupItems = allItems
            .filter((item) => item.groupId === group.id)
            .sort((a, b) => (a.order || 0) - (b.order || 0));

        // 그룹에 아이템이 있으면 첫 번째 아이템을 그룹의 order로 등록
        if (groupItems.length > 0) {
            itemEntries.push({
                id: groupItems[0].id,
                order: group.order || 0,
            });
        }
    });

    // 전체 항목을 order로 정렬하고 첫 번째 아이템 ID 반환
    if (itemEntries.length > 0) {
        itemEntries.sort((a, b) => a.order - b.order);
        return itemEntries[0].id;
    }

    // 어떤 아이템도 없는 경우 null 반환
    return null;
}

/**
 * 그룹 및 아이템 데이터 구조 변환 함수
 *
 * 패널 리스트 컴포넌트에서 사용하는 형식으로 데이터를 변환합니다.
 * 비그룹 아이템과 그룹을 order 값에 따라 정렬합니다.
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
        return {
            ...group,
            type: 'group',
            items: groupItems,
        } as G & {
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

    // 비그룹 아이템과 그룹을 모두 order 값에 따라 정렬
    return [...ungroupedItems, ...mergedGroups].sort(
        (a, b) => (a.order || 0) - (b.order || 0),
    );
}
