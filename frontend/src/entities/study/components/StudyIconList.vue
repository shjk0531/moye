<!-- src/entities/study/components/StudyIconList.vue -->
<template>
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
            <img :src="study.icon" alt="study icon" class="w-full h-full" />
        </div>
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
import { useStudyStore } from '@/store';
import { PATHS } from '@/router/paths'; // PATHS import

const router = useRouter();
const route = useRoute();
const studies = ref<Study[]>([]);
const studyStore = useStudyStore();

// URL의 studyId는 문자열입니다.
const currentStudyId = computed(() => {
    const raw = route.params[PATHS.STUDY_PARAM];
    return Array.isArray(raw) ? raw[0] : raw ?? '';
});

// Helpers to build paths
function studyBasePath(studyId: string) {
    return `${PATHS.STUDY_BASE}/${studyId}`;
}
function channelPath(studyId: string, channelId: string) {
    return `${studyBasePath(studyId)}/${PATHS.STUDY_CHANNEL}/${channelId}`;
}

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
    router.push(channelPath(String(studyId), channelId));
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
        const groups = await fetchChannelsGrouped();
        const channels = await fetchChannelsUngrouped();
        const firstChannelId = findFirstChannel(groups, channels);

        if (firstChannelId !== null) {
            router.push(channelPath(String(studyId), String(firstChannelId)));
            saveActiveChannel(studyKey, String(firstChannelId));
        } else {
            router.push(studyBasePath(String(studyId)));
        }
    } catch (error) {
        console.error('첫 번째 채널을 찾는 중 오류 발생:', error);
        router.push(studyBasePath(String(studyId)));
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

    if (activeChannelId) {
        navigateToExistingChannel(study.id, activeChannelId);
    } else {
        await navigateToFirstChannel(study.id, studyId);
    }
}

// 컴포넌트 생성 시 데이터 로드
onMounted(loadStudyData);
</script>

<style scoped></style>
