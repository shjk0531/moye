<!-- src/features/channel/components/ChannelList.vue -->

<template>
    <PanelList
        :groups="channelsGrouped"
        :items="channelsUngrouped"
        :activeItemId="activeChannelId"
        @item-click="handleChannelClick"
    />
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useStore } from 'vuex';
import { PanelList } from '@/shared/panel';
import {
    fetchChannelsGrouped,
    fetchChannelsUngrouped,
} from '@/features/channel';

const route = useRoute();
const router = useRouter();
const store = useStore();

const channelsGrouped = ref([]);
const channelsUngrouped = ref([]);

onMounted(async () => {
    channelsGrouped.value = await fetchChannelsGrouped();
    channelsUngrouped.value = await fetchChannelsUngrouped();
});

// URL에서 studyId 추출
const currentStudyId = computed(() => route.params.studyId);

// activeChannelId getter에서 nullish 연산자를 사용하여 0도 올바르게 인식하게 함
const activeChannelId = computed({
    get() {
        return store.state.activeChannelMap[currentStudyId.value] ?? null;
    },
    set(val) {
        store.commit('setActiveChannel', {
            studyId: currentStudyId.value,
            channelId: val,
        });
    },
});

// 채널 클릭 시 activeChannelId 업데이트 후 라우터 이동
function handleChannelClick(item) {
    activeChannelId.value = item.id;
    router.push(`/study/${currentStudyId.value}/channel/${item.id}`);
}
</script>
