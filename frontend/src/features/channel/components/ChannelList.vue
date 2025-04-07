<!-- src/entities/channel/components/ChannelList.vue -->
<template>
    <PanelList
        :groups="channelsGrouped"
        :items="channelsUngrouped"
        :activeItemId="activeItemId"
        listType="channel"
        @item-click="handleItemClick"
    />
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useActiveItem } from '@/shared/composables';
import { PanelList } from '@/shared/panel';
import {
    fetchChannelsGrouped,
    fetchChannelsUngrouped,
} from '@/features/channel';

const route = useRoute();
const router = useRouter();

const channelsGrouped = ref([]);
const channelsUngrouped = ref([]);

// 'channel' 리스트 타입을 지정하여 composable을 사용합니다.
const { activeItemId, checkAndRedirect, currentStudyId } =
    useActiveItem('channel');

onMounted(async () => {
    await loadChannelData();
    await redirectToActiveChannel();
});

// 채널 데이터 로드 함수
async function loadChannelData() {
    try {
        channelsGrouped.value = await fetchChannelsGrouped();
        channelsUngrouped.value = await fetchChannelsUngrouped();
    } catch (error) {
        console.error('채널 데이터 로드 중 오류 발생:', error);
    }
}

// 활성 채널로 리다이렉트 함수
async function redirectToActiveChannel() {
    try {
        await checkAndRedirect(`/study/${currentStudyId.value}/channel`);
    } catch (error) {
        console.error('채널 리다이렉트 중 오류 발생:', error);
    }
}

// 채널 아이템 클릭 핸들러
function handleItemClick(item) {
    // activeItemId 업데이트 및 해당 채널 URL로 이동
    activeItemId.value = item.id;
    navigateToChannel(item.id);
}

// 채널 페이지로 이동하는 함수
function navigateToChannel(channelId) {
    router.push(`/study/${currentStudyId.value}/channel/${channelId}`);
}
</script>
