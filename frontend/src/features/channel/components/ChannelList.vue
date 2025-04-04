<!-- src/entities/channel/components/ChannelList.vue -->
<template>
    <PanelList
        :groups="channelsGrouped"
        :items="channelsUngrouped"
        :activeItemId="activeItemId"
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
    channelsGrouped.value = await fetchChannelsGrouped();
    channelsUngrouped.value = await fetchChannelsUngrouped();

    // URL에 channelId가 없으면 저장된 마지막 채널 ID로 리다이렉트합니다.
    await checkAndRedirect(`/study/${currentStudyId.value}/channel`);
});

// 항목 클릭 시, 활성 항목을 업데이트하고 해당 채널 URL로 이동합니다.
function handleItemClick(item) {
    activeItemId.value = item.id;
    router.push(`/study/${currentStudyId.value}/channel/${item.id}`);
}
</script>
