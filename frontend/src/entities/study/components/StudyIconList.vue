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
            <img
                :src="study.profile_url"
                alt="study icon"
                class="w-full h-full"
            />
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { type StudyIcon } from '@/entities/study';
import { useStudyStore } from '@/store';
import { PATHS } from '@/router/paths'; // PATHS import
import { fetchStudyIcons } from '@/entities';

const router = useRouter();
const route = useRoute();
const studies = ref<StudyIcon[]>([]);
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
        const response = await fetchStudyIcons();
        studies.value = response.icons;
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
function navigateToExistingChannel(studyId: string, channelId: string) {
    router.push(channelPath(studyId, channelId));
}

// 활성 채널 ID 저장
function saveActiveChannel(studyId: string, channelId: string) {
    studyStore.setActiveItem({
        studyId,
        listType: 'channel',
        itemId: channelId,
    });
}

// Pinia 스토어에 스터디 정보 업데이트
function updateStudyInfoInStore(study: StudyIcon) {
    studyStore.setStudyName(study.name);
    studyStore.setStudyIcon(study.profile_url);
}

// 스터디 클릭 이벤트 핸들러
async function handleStudyClick(study: StudyIcon) {
    updateStudyInfoInStore(study);
    const studyId = String(study.id);
    const activeChannelId = getActiveChannelId(studyId);

    if (activeChannelId) {
        navigateToExistingChannel(studyId, activeChannelId);
    } else {
        router.push(studyBasePath(studyId));
    }
}

// 컴포넌트 생성 시 데이터 로드
onMounted(loadStudyData);
</script>

<style scoped></style>
