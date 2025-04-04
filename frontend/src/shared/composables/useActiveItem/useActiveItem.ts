// src/shared/composables/useActiveItem/useActiveItem.ts
import { computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useStore } from 'vuex';

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
     * @param defaultRoute - 예: '/study/1/channel'
     */
    async function checkAndRedirect(defaultRoute: string) {
        // route.params[`${listType}Id`]가 undefined 또는 null일 때,
        // activeItemId.value가 0과 같이 falsy 값이라도 명시적으로 값이 있으면 리다이렉트합니다.
        if (
            (route.params[`${listType}Id`] === undefined ||
                route.params[`${listType}Id`] === null) &&
            activeItemId.value !== null &&
            activeItemId.value !== undefined
        ) {
            await router.replace(`${defaultRoute}/${activeItemId.value}`);
        }
    }

    return { activeItemId, checkAndRedirect, currentStudyId };
}
