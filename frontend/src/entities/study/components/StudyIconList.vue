<!-- src/entities/study/components/StudyIconList.vue -->
<template>
    <div class="study-list">
        <ScrollPanel style="height: 100%">
            <div class="flex flex-col items-center justify-start w-full gap-2">
                <div
                    v-for="study in studies"
                    :key="study.id"
                    class="relative flex items-center justify-center w-full group cursor-pointer"
                    @click="handleStudyClick(study)"
                >
                    <!-- indicator: 현재 선택된 스터디면 왼쪽에 표시 -->
                    <div
                        :class="[
                            'absolute left-0 top-1/2 -translate-y-1/2 rounded transition-all duration-300',
                            String(study.id) === currentStudyId
                                ? 'w-1 h-[70%] bg-gray-50'
                                : 'w-0 group-hover:w-1 h-[50%] bg-gray-150',
                        ]"
                    ></div>
                    <div class="w-10 h-10 overflow-hidden border-2 rounded-lg">
                        <img
                            :src="study.icon"
                            alt="study icon"
                            class="w-full h-full"
                        />
                    </div>
                </div>
            </div>
        </ScrollPanel>
    </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { fetchStudies, type Study } from '@/entities/study';
import {
    fetchChannelsGrouped,
    fetchChannelsUngrouped,
    findFirstChannel,
} from '@/features/channel';
import ScrollPanel from 'primevue/scrollpanel';
import { useStudyStore } from '@/store';

const router = useRouter();
const route = useRoute();
const studies = ref<Study[]>([]);
const studyStore = useStudyStore();

// URL의 studyId는 문자열입니다.
const currentStudyId = computed(() => {
    return route.params.studyId;
});

// 스터디 데이터 로드
async function loadStudyData() {
    try {
        studies.value = await fetchStudies();
    } catch (error) {
        console.error('스터디 데이터를 불러오는 중 오류 발생:', error);
        studies.value = [];
    }
}

// 저장된 활성 채널 ID 가져오기
function getActiveChannelId(studyId: string) {
    return studyStore.activeItems[studyId]?.channel;
}

// 기존 활성 채널로 이동
function navigateToExistingChannel(
    studyId: number | string,
    channelId: string,
) {
    router.push(`/study/${studyId}/channel/${channelId}`);
}

// 활성 채널 ID 저장
function saveActiveChannel(studyId: string, channelId: string) {
    studyStore.setActiveItem({
        studyId,
        listType: 'channel',
        itemId: channelId,
    });
}

// 첫 번째 채널 찾아서 이동
async function navigateToFirstChannel(
    studyId: number | string,
    studyKey: string,
) {
    try {
        // 채널 및 그룹 데이터 로드
        const groups = await fetchChannelsGrouped();
        const channels = await fetchChannelsUngrouped();

        // 첫 번째 채널 ID 가져오기
        const firstChannelId = findFirstChannel(groups, channels);

        if (firstChannelId !== null) {
            // 첫 번째 채널 ID로 이동
            router.push(`/study/${studyId}/channel/${firstChannelId}`);

            // 첫 번째 채널 ID를 activeItemId로 저장
            saveActiveChannel(studyKey, String(firstChannelId));
        } else {
            // 채널이 없는 경우 기본 스터디 URL로 이동
            router.push(`/study/${studyId}`);
        }
    } catch (error) {
        console.error('첫 번째 채널을 찾는 중 오류 발생:', error);
        // 오류 발생시 기본 스터디 URL로 이동
        router.push(`/study/${studyId}`);
    }
}

// Pinia 스토어에 스터디 정보 업데이트
function updateStudyInfoInStore(study: any) {
    studyStore.setStudyName(study.name);
    studyStore.setStudyIcon(study.icon);
}

// 스터디 클릭 이벤트 핸들러
async function handleStudyClick(study: any) {
    updateStudyInfoInStore(study);

    const studyId = String(study.id);
    const activeChannelId = getActiveChannelId(studyId);

    if (activeChannelId !== null && activeChannelId !== undefined) {
        navigateToExistingChannel(study.id, activeChannelId);
    } else {
        await navigateToFirstChannel(study.id, studyId);
    }
}

// 타이틀 아이콘 클릭 이벤트 처리
function handleTitleIconClick() {
    router.push('/me');
    studyStore.setStudyName('Moye');
    studyStore.setStudyIcon('https://picsum.photos/200/300?random=1');
}

// 컴포넌트 생성 시 데이터 로드
onMounted(async () => {
    await loadStudyData();
});
</script>

<style scoped>
.study-list {
    flex: 1;
    overflow-y: auto;
    overflow-x: hidden;
}

.study-list :deep(.p-scrollpanel-content) {
    height: 100%;
}

.study-list :deep(.p-scrollpanel-bar) {
    background-color: var(--custom-scrollbar-color);
    border-radius: 4px;
    transition: background-color 0.3s ease;
}

.study-list :deep(.p-scrollpanel-bar:hover) {
    background-color: var(--custom-scrollbar-color-hover);
}

.study-list :deep(.p-scrollpanel-bar:active) {
    background-color: var(--custom-scrollbar-color-active);
}

.study-list :deep(.p-scrollpanel-bar-y) {
    width: calc(var(--spacing)) !important;
}
</style>
