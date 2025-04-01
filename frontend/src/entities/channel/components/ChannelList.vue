<!-- /src/entities/channel/components/ChannelList.vue -->

<template>
    <div class="channel-list">
        <div class="flex flex-col w-full">
            <div v-for="channel in mergedChannels" :key="channel.id">
                <!-- 그룹 채널: 내부 PanelSection에서 channel-click 이벤트를 emit합니다. -->
                <PanelSection
                    v-if="channel.type === 'group'"
                    :label="channel.label"
                    :items="channel.items"
                    :activeChannelId="activeChannelId"
                    @channel-click="handleChannelClick"
                />
                <!-- 비그룹 채널: 클릭 시 바로 handleChannelClick 호출 -->
                <PanelItem
                    v-else
                    :item="channel"
                    :isActive="channel.id === activeChannelId"
                    @click.native="handleChannelClick(channel)"
                />
            </div>
        </div>
    </div>
</template>

<script>
import {
    fetchChannelsGrouped,
    fetchChannelsUngrouped,
} from '@/features/channel';
import { PanelSection, PanelItem } from '@/shared/panel';

export default {
    name: 'ChannelList',
    components: {
        PanelSection,
        PanelItem,
    },
    data() {
        return {
            channelsGrouped: [],
            channelsUngrouped: [],
        };
    },
    computed: {
        mergedChannels() {
            const groups = this.channelsGrouped.map((group) => ({
                ...group,
                type: 'group',
            }));
            const ungrouped = this.channelsUngrouped.map((channel) => ({
                ...channel,
                type: 'ungrouped',
            }));
            return [...groups, ...ungrouped].sort(
                (a, b) => (a.order || 9999) - (b.order || 9999),
            );
        },
        // URL의 studyId는 이미 문자열입니다.
        currentStudyId() {
            return this.$route.params.studyId;
        },
        // 전역 저장소에서 studyId별 activeChannelId 관리
        activeChannelId: {
            get() {
                return (
                    this.$globalState.activeChannelMap[this.currentStudyId] ||
                    null
                );
            },
            set(val) {
                this.$globalState.activeChannelMap[this.currentStudyId] = val;
            },
        },
    },
    watch: {
        // 채널 목록이 로드되고 아직 activeChannelId가 없다면 첫 번째 채널을 기본값으로 설정
        mergedChannels(newVal) {
            if (newVal.length > 0 && !this.activeChannelId) {
                this.activeChannelId = newVal[0].id;
                if (this.$route.params.channelId !== String(newVal[0].id)) {
                    this.$router.replace(
                        `/study/${this.currentStudyId}/${newVal[0].id}`,
                    );
                }
            }
        },
    },
    methods: {
        handleChannelClick(channel) {
            // 선택한 채널을 전역 state에 업데이트 후 라우터 이동
            this.activeChannelId = channel.id;
            this.$router.push(`/study/${this.currentStudyId}/${channel.id}`);
        },
    },
    async created() {
        this.channelsGrouped = await fetchChannelsGrouped();
        this.channelsUngrouped = await fetchChannelsUngrouped();
    },
};
</script>
