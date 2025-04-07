// 공통으로 사용되는 훅 함수들

import { ref, onMounted, onUnmounted } from 'vue';
import type { Ref } from 'vue';

/**
 * 윈도우 크기를 추적하는 훅
 */
export function useWindowSize() {
    const width = ref(window.innerWidth);
    const height = ref(window.innerHeight);

    function updateSize() {
        width.value = window.innerWidth;
        height.value = window.innerHeight;
    }

    onMounted(() => {
        window.addEventListener('resize', updateSize);
    });

    onUnmounted(() => {
        window.removeEventListener('resize', updateSize);
    });

    return { width, height };
}

/**
 * 로컬 스토리지 상태를 관리하는 훅
 */
export function useLocalStorage<T>(key: string, initialValue: T) {
    const storedValue = ref<T>(initialValue);

    // 초기값 로드
    onMounted(() => {
        try {
            const item = window.localStorage.getItem(key);
            if (item) {
                storedValue.value = JSON.parse(item);
            }
        } catch (error) {
            console.error(error);
        }
    });

    const setValue = (value: T) => {
        try {
            storedValue.value = value;
            window.localStorage.setItem(key, JSON.stringify(value));
        } catch (error) {
            console.error(error);
        }
    };

    return [storedValue, setValue] as const;
}
