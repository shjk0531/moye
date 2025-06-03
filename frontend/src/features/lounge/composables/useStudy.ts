import { ref, computed } from 'vue';
import { fetchLounges, fetchLoungeById } from '@/entities/lounge';
import type { Lounge } from '@/entities/lounge';
import { useLoungeStore } from '@/store';

export function useLounge() {
    const lounges = ref<Lounge[]>([]);
    const currentLounge = ref<Lounge | null>(null);
    const loading = ref(false);
    const error = ref<string | null>(null);

    // 스터디 목록 로드
    const loadLounges = async () => {
        try {
            loading.value = true;
            error.value = null;
            lounges.value = await fetchLounges();
        } catch (e) {
            console.error('스터디 목록을 불러오는 중 오류 발생:', e);
            error.value = '스터디 목록을 불러올 수 없습니다.';
            lounges.value = [];
        } finally {
            loading.value = false;
        }
    };

    // 특정 스터디 정보 로드
    const loadLoungeById = async (loungeId: number) => {
        try {
            loading.value = true;
            error.value = null;
            currentLounge.value = await fetchLoungeById(loungeId);

            // 스토어에 스터디 정보 업데이트
            const loungeStore = useLoungeStore();
            loungeStore.setLoungeName(currentLounge.value.name);
            loungeStore.setLoungeIcon(currentLounge.value.icon);
        } catch (e) {
            console.error(
                `스터디 ID ${loungeId}의 정보를 불러오는 중 오류 발생:`,
                e,
            );
            error.value = '스터디 정보를 불러올 수 없습니다.';
            currentLounge.value = null;
        } finally {
            loading.value = false;
        }
    };

    return {
        lounges,
        currentLounge,
        loading,
        error,
        loadLounges,
        loadLoungeById,
    };
}
