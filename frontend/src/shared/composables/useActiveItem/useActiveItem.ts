// src/shared/composables/useActiveItem/useActiveItem.ts
import { computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useStudyStore } from '@/store';

// 각 리스트 타입별 API 함수 import
import {
    fetchChannelsGrouped,
    fetchChannelsUngrouped,
    findFirstChannel,
} from '@/features/channel';
import {
    fetchCalendarsGrouped,
    fetchCalendarsUngrouped,
    findFirstCalendar,
} from '@/features/calendar';

// API 함수 타입 정의
interface ApiMapping {
    fetchGroups: () => Promise<any[]>;
    fetchItems: () => Promise<any[]>;
    findFirst: (groups: any[], items: any[]) => number | null;
}

// 리스트 타입별 API 함수 매핑
const apiMapping: Record<string, ApiMapping> = {
    channel: {
        fetchGroups: fetchChannelsGrouped,
        fetchItems: fetchChannelsUngrouped,
        findFirst: findFirstChannel,
    },
    calendar: {
        fetchGroups: fetchCalendarsGrouped,
        fetchItems: fetchCalendarsUngrouped,
        findFirst: findFirstCalendar,
    },
    // 다른 리스트 타입이 추가될 경우 여기에 추가
};

export function useActiveItem(listType: string) {
    const route = useRoute();
    const router = useRouter();
    const studyStore = useStudyStore();

    // 지원하지 않는 listType인 경우 경고 로그 출력
    if (!apiMapping[listType]) {
        console.warn(
            `useActiveItem: 지원하지 않는 리스트 타입입니다: ${listType}`,
        );
    }

    // URL에서 studyId를 추출합니다.
    const currentStudyId = computed(() => route.params.studyId as string);

    // Pinia 스토어의 activeItems를 읽거나 업데이트합니다.
    const activeItemId = computed({
        get() {
            return (
                studyStore.activeItems[currentStudyId.value]?.[listType] ?? null
            );
        },
        set(val: string) {
            studyStore.setActiveItem({
                studyId: currentStudyId.value,
                listType,
                itemId: val,
            });
        },
    });

    /**
     * URL에 목록 항목 ID가 없는 경우, 저장된 마지막 항목 ID로 리다이렉트합니다.
     * 마지막 항목 ID가 없는 경우, 첫 번째 항목으로 리다이렉트합니다.
     * @param defaultRoute - 예: '/study/1/channel'
     */
    async function checkAndRedirect(defaultRoute: string) {
        // route.params[`${listType}Id`]가 undefined 또는 null일 때,
        if (
            route.params[`${listType}Id`] === undefined ||
            route.params[`${listType}Id`] === null
        ) {
            // activeItemId.value가 있으면 해당 항목으로 리다이렉트
            if (
                activeItemId.value !== null &&
                activeItemId.value !== undefined
            ) {
                await router.replace(`${defaultRoute}/${activeItemId.value}`);
            }
            // activeItemId.value가 없으면 첫 번째 항목으로 리다이렉트
            else if (apiMapping[listType]) {
                try {
                    // 해당 타입의 API 함수 가져오기
                    const { fetchGroups, fetchItems, findFirst } =
                        apiMapping[listType];

                    // 그룹 및 아이템 데이터 로드
                    const groups = await fetchGroups();
                    const items = await fetchItems();

                    // 첫 번째 아이템 ID 가져오기
                    const firstItemId = findFirst(groups, items);

                    if (firstItemId !== null) {
                        // 첫 번째 아이템 ID로 리다이렉트
                        await router.replace(`${defaultRoute}/${firstItemId}`);

                        // 첫 번째 아이템 ID를 activeItemId로 설정
                        activeItemId.value = String(firstItemId);
                    }
                } catch (error) {
                    console.error(
                        `첫 번째 ${listType}을(를) 찾는 중 오류 발생:`,
                        error,
                    );
                }
            }
        }
    }

    return { activeItemId, checkAndRedirect, currentStudyId };
}
