import { ref, computed } from 'vue';
import {
    fetchChannelsGrouped,
    fetchChannelsUngrouped,
    findFirstChannel,
} from '@/entities/channel';
import type { ChannelGroup, Channel } from '@/entities/channel';
import { useStudyStore } from '@/store';

export function useChannel() {
    const channelGroups = ref<ChannelGroup[]>([]);
    const channels = ref<Channel[]>([]);
    const currentChannelId = ref<number | null>(null);
    const loading = ref(false);
    const error = ref<string | null>(null);

    // 채널 그룹 및 채널 목록 로드
    const loadChannels = async () => {
        try {
            loading.value = true;
            error.value = null;
            const [groups, channelList] = await Promise.all([
                fetchChannelsGrouped(),
                fetchChannelsUngrouped(),
            ]);
            channelGroups.value = groups;
            channels.value = channelList;
        } catch (e) {
            console.error('채널 데이터를 불러오는 중 오류 발생:', e);
            error.value = '채널 데이터를 불러올 수 없습니다.';
            channelGroups.value = [];
            channels.value = [];
        } finally {
            loading.value = false;
        }
    };

    // 첫 번째 채널 ID 찾기
    const findFirstChannelId = computed(() => {
        return findFirstChannel(channelGroups.value, channels.value);
    });

    // 현재 채널 설정
    const setCurrentChannel = (channelId: number, studyId: string) => {
        currentChannelId.value = channelId;

        // 스토어에 현재 채널 ID 저장
        const studyStore = useStudyStore();
        studyStore.setActiveItem({
            studyId,
            listType: 'channel',
            itemId: String(channelId),
        });
    };

    // 특정 그룹에 속한 채널 필터링
    const getChannelsByGroupId = (groupId: number | null) => {
        return channels.value.filter((channel) => channel.groupId === groupId);
    };

    return {
        channelGroups,
        channels,
        currentChannelId,
        loading,
        error,
        loadChannels,
        findFirstChannelId,
        setCurrentChannel,
        getChannelsByGroupId,
    };
}
