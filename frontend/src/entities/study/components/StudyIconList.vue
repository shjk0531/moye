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

<script>
import { fetchStudies } from '@/features/study';
import ScrollPanel from 'primevue/scrollpanel';

export default {
    name: 'StudyIconList',
    components: { ScrollPanel },
    data() {
        return {
            studies: [],
        };
    },
    async created() {
        // 스터디 데이터를 API를 통해 불러옵니다.
        this.studies = await fetchStudies();
    },
    computed: {
        // URL의 studyId는 문자열입니다.
        currentStudyId() {
            return this.$route.params.studyId;
        },
    },
    methods: {
        handleStudyClick(study) {
            // Vuex 스토어에 스터디 관련 정보 업데이트
            this.$store.commit('setStudyName', study.name);
            this.$store.commit('setStudyIcon', study.icon);

            // study.id를 문자열로 변환하여 사용
            const key = String(study.id);
            // 기본 목록 타입을 "channel"로 가정하여 activeItems에서 해당 값을 조회
            const activeChannelId = this.$store.state.activeItems[key]?.channel;

            if (activeChannelId !== null && activeChannelId !== undefined) {
                // 만약 해당 스터디에 저장된 마지막 접속 채널이 있다면, 해당 채널 URL로 이동합니다.
                this.$router.push(
                    `/study/${study.id}/channel/${activeChannelId}`,
                );
            } else {
                // 없다면 기본 스터디 URL로 이동합니다.
                this.$router.push(`/study/${study.id}`);
            }
        },
        // 필요 시 타이틀 아이콘 클릭 이벤트 처리
        handleTitleIconClick() {
            this.$router.push('/me');
            this.$store.commit('setStudyName', 'Moye');
            this.$store.commit(
                'setStudyIcon',
                'https://picsum.photos/200/300?random=1',
            );
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
