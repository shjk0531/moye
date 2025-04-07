// src/shared/composables/useActiveItem/useActiveItem.ts
import { computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useStore } from 'vuex';
import {
    findFirstChannel,
    fetchChannelsGrouped,
    fetchChannelsUngrouped,
} from '@/features/channel';

export function useActiveItem(listType: string) {
    const route = useRoute();
    const router = useRouter();
    const store = useStore();

    // URL에서 studyId를 추출합니다.
    const currentStudyId = computed(() => route.params.studyId as string);

    // Vuex 스토어의 activeItems를 읽거나 업데이트합니다.
    const activeItemId = computed({
        get() {
            return (
                store.state.activeItems[currentStudyId.value]?.[listType] ??
                null
            );
        },
        set(val: string) {
            store.commit('setActiveItem', {
                studyId: currentStudyId.value,
                listType,
                itemId: val,
            });
        },
    });

    /**
     * URL에 목록 항목 ID가 없는 경우, 저장된 마지막 항목 ID로 리다이렉트합니다.
     * 마지막 항목 ID가 없는 경우, 첫 번째 채널로 리다이렉트합니다.
     * @param defaultRoute - 예: '/study/1/channel'
     */
    async function checkAndRedirect(defaultRoute: string) {
        // route.params[`${listType}Id`]가 undefined 또는 null일 때,
        if (
            route.params[`${listType}Id`] === undefined ||
            route.params[`${listType}Id`] === null
        ) {
            // activeItemId.value가 있으면 해당 채널로 리다이렉트
            if (
                activeItemId.value !== null &&
                activeItemId.value !== undefined
            ) {
                await router.replace(`${defaultRoute}/${activeItemId.value}`);
            }
            // activeItemId.value가 없으면 첫 번째 채널로 리다이렉트 (채널 타입인 경우만)
            else if (listType === 'channel') {
                try {
                    // 채널 및 그룹 데이터 로드
                    const groups = await fetchChannelsGrouped();
                    const channels = await fetchChannelsUngrouped();

                    // 첫 번째 채널 ID 가져오기
                    const firstChannelId = findFirstChannel(groups, channels);

                    if (firstChannelId !== null) {
                        // 첫 번째 채널 ID로 리다이렉트
                        await router.replace(
                            `${defaultRoute}/${firstChannelId}`,
                        );

                        // 첫 번째 채널 ID를 activeItemId로 설정
                        activeItemId.value = String(firstChannelId);
                    }
                } catch (error) {
                    console.error('첫 번째 채널을 찾는 중 오류 발생:', error);
                }
            }
        }
    }

    return { activeItemId, checkAndRedirect, currentStudyId };
}
