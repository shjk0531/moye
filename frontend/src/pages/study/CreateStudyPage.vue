<template>
    <div
        class="create-study-page flex justify-center bg-gray-100 dark:bg-gray-950 min-h-screen w-full overflow-hidden"
    >
        <CreateStudyProcess />
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAppStore } from '@/store';
import { useToast } from 'primevue/usetoast';
import { createStudy } from '@/entities/study/api/studyApi';
import CreateStudyProcess from '@/processes/study/CreateStudyProcess.vue';
import type { StudyCreatePayload } from '@/entities/study/models/types';

// 위자드 또는 기본 폼 중 선택 (라우터 쿼리 파라미터로 제어 가능)
const useWizard = ref(true);

const router = useRouter();
const appStore = useAppStore();
const toast = useToast();

// 취소 시 마지막으로 저장된 라우트로 이동
function onCancel() {
    const lastRoute = appStore.getLastRoute || '/';
    router.push(lastRoute);
}

// 폼 제출 시 API 호출
async function onSubmit(payload: StudyCreatePayload) {
    try {
        await createStudy(payload);

        toast.add({
            severity: 'success',
            summary: '스터디 생성',
            detail: '스터디가 성공적으로 생성되었습니다.',
            life: 3000,
        });

        // 생성 후 스터디 목록 페이지로 이동
        router.push('/study');
    } catch (error) {
        console.error('스터디 생성 실패:', error);
        toast.add({
            severity: 'error',
            summary: '생성 실패',
            detail: '스터디 생성에 실패했습니다.',
            life: 5000,
        });
    }
}
</script>

<style lang="scss">
/* 필요 시 추가 스타일을 작성하세요 */
</style>
