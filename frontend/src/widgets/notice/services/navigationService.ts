import { Store } from 'vuex';
import type { Router, Route } from 'vue-router';
import {
    fetchChannelsGrouped,
    fetchChannelsUngrouped,
    findFirstChannel,
} from '@/features/channel';
import {
    fetchCalendarsGrouped,
    fetchCalendarsUngrouped,
    findFirstCalendar,
} from '@/features/calender';

interface ApiFunction {
    fetchGroups: () => Promise<any[]>;
    fetchItems: () => Promise<any[]>;
    findFirst: (groups: any[], items: any[]) => number | null;
    route: (studyId: string, itemId: string | number) => string;
    enabled: boolean; // 기능 활성화 여부
    paramName: string; // URL 파라미터 이름
}

// 아이템 타입별 API 함수 매핑
const apiMapping: Record<string, ApiFunction> = {
    channel: {
        fetchGroups: fetchChannelsGrouped,
        fetchItems: fetchChannelsUngrouped,
        findFirst: findFirstChannel,
        route: (studyId, itemId) => `/study/${studyId}/channel/${itemId}`,
        enabled: true,
        paramName: 'channelId',
    },
    calendar: {
        fetchGroups: fetchCalendarsGrouped,
        fetchItems: fetchCalendarsUngrouped,
        findFirst: findFirstCalendar,
        route: (studyId, itemId) => `/study/${studyId}/calendar/${itemId}`,
        enabled: true,
        paramName: 'calendarId',
    },
    post: {
        fetchGroups: async () => [],
        fetchItems: async () => [],
        findFirst: () => null,
        route: (studyId, itemId) => `/study/${studyId}/post/${itemId}`,
        enabled: false, // 아직 구현 안됨
        paramName: 'postId',
    },
    note: {
        fetchGroups: async () => [],
        fetchItems: async () => [],
        findFirst: () => null,
        route: (studyId, itemId) => `/study/${studyId}/note/${itemId}`,
        enabled: false, // 아직 구현 안됨
        paramName: 'noteId',
    },
    // 추후 다른 아이템 타입이 추가될 경우 여기에 추가
};

/**
 * 아이템 타입이 지원되는지 확인하는 함수
 * @param itemType - 아이템 타입
 * @returns 지원 여부
 */
export function isItemTypeSupported(itemType: string): boolean {
    return !!apiMapping[itemType] && apiMapping[itemType].enabled;
}

/**
 * 사용자가 마지막으로 접속한 아이템으로 이동하는 함수
 * @param itemType - 이동할 아이템 타입 ('channel', 'calendar' 등)
 * @param studyId - 스터디 ID
 * @param store - Vuex 스토어
 * @param router - Vue 라우터
 * @returns Promise<void>
 */
export async function navigateToItem(
    itemType: string,
    studyId: string,
    store: Store<any>,
    router: Router,
): Promise<void> {
    try {
        if (!studyId) {
            console.error('유효한 studyId가 없습니다.');
            return;
        }

        // 해당 아이템 타입에 대한 API 함수 가져오기
        const api = apiMapping[itemType];
        if (!api) {
            console.error(`지원하지 않는 아이템 타입입니다: ${itemType}`);
            return;
        }

        // 아이템 타입이 활성화되지 않은 경우
        if (!api.enabled) {
            console.error(`${itemType} 아이템 타입은 현재 지원되지 않습니다.`);
            return;
        }

        // 마지막 접속 아이템 ID 가져오기
        const lastActiveItemId = store.state.activeItems[studyId]?.[itemType];

        if (lastActiveItemId !== null && lastActiveItemId !== undefined) {
            // 마지막 접속 아이템이 있으면 해당 아이템으로 이동
            router.push(api.route(studyId, lastActiveItemId));
        } else {
            // 마지막 접속 아이템이 없으면 첫 번째 아이템 찾기
            const groups = await api.fetchGroups();
            const items = await api.fetchItems();
            const firstItemId = api.findFirst(groups, items);

            if (firstItemId !== null) {
                // 첫 번째 아이템으로 이동
                router.push(api.route(studyId, firstItemId));

                // 첫 번째 아이템 ID를 activeItemId로 저장
                store.commit('setActiveItem', {
                    studyId,
                    listType: itemType,
                    itemId: String(firstItemId),
                });
            } else {
                console.error(`${itemType} 아이템을 찾을 수 없습니다.`);
            }
        }
    } catch (error) {
        console.error(`${itemType} 이동 중 오류 발생:`, error);
    }
}

/**
 * 현재 라우트와 아이템 타입 및 ID를 비교하여 활성 상태 여부 반환
 * @param itemType - 아이템 타입
 * @param currentPath - 현재 라우트 경로
 * @param route - 현재 라우트 객체 (옵션)
 * @param itemId - 아이템 ID (옵션)
 * @returns 활성 상태 여부
 */
export function isItemActive(
    itemType: string,
    currentPath: string,
    route?: any,
    itemId?: number | string,
): boolean {
    // 해당 아이템 타입이 활성화되어 있는지 확인
    if (!isItemTypeSupported(itemType) && itemType !== 'member') {
        return false;
    }

    // 경로에 해당 타입이 포함되어 있는지 기본 확인
    const typePathMapping: Record<string, string> = {
        channel: '/channel/',
        calendar: '/calendar/',
        post: '/post/',
        note: '/note/',
    };

    const typePath = typePathMapping[itemType] || '';
    if (!typePath || !currentPath.includes(typePath)) {
        return false;
    }

    // 아이템 ID 비교 (라우트 객체와 아이템 ID가 모두 제공된 경우)
    if (route && itemId !== undefined) {
        const paramName = apiMapping[itemType]?.paramName || `${itemType}Id`;
        const routeParamId = route.params[paramName];

        // 라우트에 ID 파라미터가 있는지 확인
        if (routeParamId) {
            return String(routeParamId) === String(itemId);
        }
    }

    // 라우트 객체나 아이템 ID가 없는 경우, 경로 기반으로만 확인
    return true;
}

/**
 * 사용 가능한 아이템 타입 목록 반환
 * @param onlyEnabled - 활성화된 아이템 타입만 반환할지 여부
 * @returns 사용 가능한 아이템 타입 목록
 */
export function getAvailableItemTypes(onlyEnabled: boolean = true): string[] {
    return Object.entries(apiMapping)
        .filter(([_, api]) => !onlyEnabled || api.enabled)
        .map(([type]) => type);
}

/**
 * 아이템 타입별 경로 생성 함수 반환
 * @param itemType - 아이템 타입
 * @returns 경로 생성 함수 또는 null
 */
export function getRouteGenerator(
    itemType: string,
): ((studyId: string, itemId: string | number) => string) | null {
    return apiMapping[itemType]?.route || null;
}

/**
 * 아이템 타입의 URL 파라미터 이름 반환
 * @param itemType - 아이템 타입
 * @returns URL 파라미터 이름
 */
export function getParamName(itemType: string): string {
    return apiMapping[itemType]?.paramName || `${itemType}Id`;
}
