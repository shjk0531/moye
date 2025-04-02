<!-- src/entities/study/components/StudyIconList.vue -->
<template>
    <!-- Study Icons List: 스크롤 가능한 영역으로 변경 -->
    <div class="study-list">
        <ScrollPanel style="height: 100%">
            <div class="flex flex-col items-center justify-start w-full gap-2">
                <div
                    v-for="study in studies"
                    :key="study.id"
                    class="relative flex items-center justify-center w-full group cursor-pointer"
                    @click="handleStudyClick(study)"
                >
                    <!-- indicator: 현재 스터디면 상시 활성 -->
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

<script>
import { fetchStudies } from '@/features/study/api/studiesApi';
import ScrollPanel from 'primevue/scrollpanel';

export default {
    name: 'StudyIconList',
    components: {
        ScrollPanel,
    },
    data() {
        return {
            studies: [],
        };
    },
    async created() {
        this.studies = await fetchStudies();
    },
    computed: {
        // URL의 studyId는 문자열입니다.
        currentStudyId() {
            return this.$route.params.studyId;
        },
        isTitleActive() {
            return this.$route.path === '/me';
        },
    },
    methods: {
        handleStudyClick(study) {
            // 전역 상태 업데이트
            this.$globalState.studyName = study.name;
            this.$globalState.studyIcon = study.icon;
            // study.id를 문자열로 변환하여 사용
            const key = String(study.id);
            // 해당 스터디에 저장된 activeChannelId가 있는지 확인
            const activeChannelId = this.$globalState.activeChannelMap[key];
            if (activeChannelId) {
                // activeChannelId가 있으면 해당 채널로 이동
                this.$router.push(
                    `/study/${study.id}/channel/${activeChannelId}`,
                );
            } else {
                // activeChannelId가 없으면 스터디의 기본 URL로 이동
                this.$router.push(`/study/${study.id}`);
            }
        },
        handleTitleIconClick() {
            this.$router.push('/me');
            this.$globalState.studyName = 'Moye';
            this.$globalState.studyIcon =
                'https://picsum.photos/200/300?random=1';
        },
    },
};
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
