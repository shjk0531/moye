import { ref, computed } from 'vue';
import { fetchStudies, fetchStudyById } from '@/entities/study';
import type { Study } from '@/entities/study';
import { useStudyStore } from '@/store';

export function useStudy() {
    const studies = ref<Study[]>([]);
    const currentStudy = ref<Study | null>(null);
    const loading = ref(false);
    const error = ref<string | null>(null);

    // 스터디 목록 로드
    const loadStudies = async () => {
        try {
            loading.value = true;
            error.value = null;
            studies.value = await fetchStudies();
        } catch (e) {
            console.error('스터디 목록을 불러오는 중 오류 발생:', e);
            error.value = '스터디 목록을 불러올 수 없습니다.';
            studies.value = [];
        } finally {
            loading.value = false;
        }
    };

    // 특정 스터디 정보 로드
    const loadStudyById = async (studyId: number) => {
        try {
            loading.value = true;
            error.value = null;
            currentStudy.value = await fetchStudyById(studyId);

            // 스토어에 스터디 정보 업데이트
            const studyStore = useStudyStore();
            studyStore.setStudyName(currentStudy.value.name);
            studyStore.setStudyIcon(currentStudy.value.icon);
        } catch (e) {
            console.error(
                `스터디 ID ${studyId}의 정보를 불러오는 중 오류 발생:`,
                e,
            );
            error.value = '스터디 정보를 불러올 수 없습니다.';
            currentStudy.value = null;
        } finally {
            loading.value = false;
        }
    };

    return {
        studies,
        currentStudy,
        loading,
        error,
        loadStudies,
        loadStudyById,
    };
}
