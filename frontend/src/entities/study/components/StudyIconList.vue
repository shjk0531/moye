<!-- src/entities/study/components/StudyIconList.vue -->
<template>
    <div
        v-for="study in studies"
        :key="study.id"
        class="relative flex items-center justify-center w-full group cursor-pointer"
        @click="handleStudyClick(study)"
    >
        <!-- indicator -->
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
import type { StudyIcon } from '@/entities/study';
import { useStudyStore } from '@/store/study'; // <-- 경로 주의!
import { PATHS } from '@/router/paths';
import { fetchStudyIcons } from '@/entities';

// router, store 세팅
const router = useRouter();
const route = useRoute();
const studyStore = useStudyStore();

// 스터디 리스트
const studies = ref<StudyIcon[]>([]);

// 현재 선택된 studyId (URL 파라미터)
const currentStudyId = computed(() => {
    const raw = route.params[PATHS.STUDY_PARAM];
    return Array.isArray(raw) ? raw[0] : raw ?? '';
});

// 스터디 아이콘 로드
async function loadStudyData() {
    try {
        const { icons } = await fetchStudyIcons();
        studies.value = icons;
    } catch {
        studies.value = [];
    }
}

// Helpers: 경로 빌더
function studyBasePath(studyId: string) {
    return `${PATHS.STUDY_BASE}/${studyId}`;
}
function channelPath(studyId: string, channelId: string) {
    return `${studyBasePath(studyId)}/${PATHS.STUDY_CHANNEL}/${channelId}`;
}

// 클릭 핸들러
async function handleStudyClick(study: StudyIcon) {
    const studyId = String(study.id);

    // 1) 스토어에 스터디 정보 저장
    studyStore.setCurrentStudyInfo({
        id: study.id,
        name: study.name,
        icon: study.profile_url,
        type: 'channel',
        channels: {
            items: [],
        },
    });

    // 2) 이미 활성화된 채널이 있으면 그걸 쓰고, 없으면 첫 번째 채널을 API 호출로 가져와서
    const activeChannel = studyStore.getActiveItemByStudyAndType(
        studyId,
        'channel',
    );
    const channelId =
        activeChannel || (await studyStore.getFirstChannelId(studyId)) || ''; // 채널이 하나도 없으면 빈문자열

    console.log('channelId:', channelId);
    // 3) 스토어에도 저장하고, 라우팅
    studyStore.setActiveItem({
        studyId,
        listType: 'channel',
        itemId: channelId,
    });
    if (channelId) {
        router.push(channelPath(studyId, channelId));
    } else {
        // 혹시 채널이 하나도 없는 스터디라면 기본 페이지로
        router.push(studyBasePath(studyId));
    }
}

// 마운트 시 아이콘 로드
onMounted(loadStudyData);
</script>

<style scoped></style>
