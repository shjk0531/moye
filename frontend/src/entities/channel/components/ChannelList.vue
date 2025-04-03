<!-- /src/entities/channel/components/ChannelList.vue -->
<template>
    <div class="channel-list">
        <div class="flex flex-col w-full">
            <div v-for="channel in mergedChannels" :key="channel.id">
                <!-- 그룹 채널이면 PanelSection 컴포넌트 사용 -->
                <PanelSection
                    v-if="channel.type === 'group'"
                    :label="channel.label"
                    :items="channel.items"
                    :activeChannelId="activeChannelId"
                    @channel-click="handleChannelClick"
                />
                <!-- 비그룹 채널이면 PanelItem 컴포넌트 사용 -->
                <PanelItem
                    v-else
                    :item="channel"
                    :isActive="channel.id === activeChannelId"
                    @click="handleChannelClick(channel)"
                />
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import {
    fetchChannelsGrouped,
    fetchChannelsUngrouped,
} from '@/features/channel';

import { PanelSection, PanelItem } from '@/shared/panel';

// 전역 상태 (예: Vuex 혹은 Pinia)를 사용하는 경우 아래와 같이 import
import { useStore } from 'vuex';

const route = useRoute();
const router = useRouter();
const store = useStore();

const channelsGrouped = ref([]);
const channelsUngrouped = ref([]);

// API 호출 후 데이터 저장
onMounted(async () => {
    channelsGrouped.value = await fetchChannelsGrouped();
    channelsUngrouped.value = await fetchChannelsUngrouped();
});

// 그룹 데이터와 채널 데이터를 결합하여 렌더링 리스트 구성
const mergedChannels = computed(() => {
    // 각 그룹에 해당하는 채널들을 items 배열에 할당
    const groups = channelsGrouped.value.map((group) => {
        const items = channelsUngrouped.value
            .filter((ch) => ch.groupId === group.id)
            .sort((a, b) => (a.order || 0) - (b.order || 0));
        return { ...group, type: 'group', items };
    });
    // 그룹에 속하지 않은 채널
    const ungrouped = channelsUngrouped.value
        .filter((ch) => ch.groupId === null)
        .map((ch) => ({ ...ch, type: 'ungrouped' }));
    // 전체를 order 값 기준으로 정렬
    return [...groups, ...ungrouped].sort(
        (a, b) => (a.order || 9999) - (b.order || 9999),
    );
});

// URL의 studyId (문자열) 추출
const currentStudyId = computed(() => route.params.studyId);

// 전역 상태에서 studyId별 activeChannelId 관리 (예: Vuex 스토어)
const activeChannelId = computed({
    get() {
        return store.state.activeChannelMap[currentStudyId.value] || null;
    },
    set(val) {
        store.commit('setActiveChannel', {
            studyId: currentStudyId.value,
            channelId: val,
        });
    },
});

// 채널 클릭 핸들러: 전역 activeChannelId 업데이트 후 라우터 이동
function handleChannelClick(channel) {
    activeChannelId.value = channel.id;
    router.push(`/study/${currentStudyId.value}/channel/${channel.id}`);
}
</script>

<style scoped lang="scss">
.channel-list {
    grid-area: chatRoomList;
    width: var(--custom-chat-room-list-width);
    flex: 1;
    overflow-y: auto;
}
</style>
