<!-- src/entities/lounge/components/LoungeIconList.vue -->
<template>
    <div
        v-for="lounge in lounges"
        :key="lounge.id"
        class="relative flex items-center justify-center w-full group cursor-pointer"
        @click="handleLoungeClick(lounge)"
    >
        <!-- indicator -->
        <div
            :class="[
                'absolute left-0 top-1/2 -translate-y-1/2 rounded transition-all duration-300',
                String(lounge.id) === currentLoungeId
                    ? 'w-1 h-[70%] bg-gray-50'
                    : 'w-0 group-hover:w-1 h-[50%] bg-gray-150',
            ]"
        ></div>
        <div class="w-10 h-10 overflow-hidden border-2 rounded-lg">
            <img
                :src="lounge.profile_url"
                alt="lounge icon"
                class="w-full h-full"
            />
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import type { LoungeIcon } from '@/entities/lounge';
import { useLoungeStore } from '@/store/lounge';
import { PATHS } from '@/router/paths';
import { fetchLoungeIcons } from '@/entities';

// router, store 세팅
const router = useRouter();
const route = useRoute();
const loungeStore = useLoungeStore();

// 라운지 리스트
const lounges = ref<LoungeIcon[]>([]);

// 현재 선택된 loungeId (URL 파라미터)
const currentLoungeId = computed(() => {
    const raw = route.params[PATHS.LOUNGE_PARAM];
    return Array.isArray(raw) ? raw[0] : raw ?? '';
});

// 라운지 아이콘 로드
async function loadLoungeData() {
    try {
        const { icons } = await fetchLoungeIcons();
        lounges.value = icons;
    } catch {
        lounges.value = [];
    }
}

// Helpers: 경로 빌더
function loungeBasePath(loungeId: string) {
    return `${PATHS.LOUNGE_BASE}/${loungeId}`;
}
function channelPath(loungeId: string, channelId: string) {
    return `${loungeBasePath(loungeId)}/${PATHS.LOUNGE_CHANNEL}/${channelId}`;
}

// 클릭 핸들러
async function handleLoungeClick(lounge: LoungeIcon) {
    const loungeId = String(lounge.id);

    // 1) 스토어에 라운지 정보 저장
    loungeStore.setCurrentLoungeInfo({
        id: lounge.id,
        name: lounge.name,
        icon: lounge.profile_url,
        type: 'channel',
        channels: {
            items: [],
        },
    });

    // 2) 이미 활성화된 채널이 있으면 그걸 쓰고, 없으면 첫 번째 채널을 API 호출로 가져와서
    const activeChannel = loungeStore.getActiveItemByLoungeAndType(
        loungeId,
        'channel',
    );
    const channelId =
        activeChannel || (await loungeStore.getFirstChannelId(loungeId)) || ''; // 채널이 하나도 없으면 빈문자열

    console.log('channelId:', channelId);
    // 3) 스토어에도 저장하고, 라우팅
    loungeStore.setActiveItem({
        loungeId,
        listType: 'channel',
        itemId: channelId,
    });
    if (channelId) {
        router.push(channelPath(loungeId, channelId));
    } else {
        // 혹시 채널이 하나도 없는 라운지라면 기본 페이지로
        router.push(loungeBasePath(loungeId));
    }
}

// 마운트 시 아이콘 로드
onMounted(loadLoungeData);
</script>

<style scoped></style>
