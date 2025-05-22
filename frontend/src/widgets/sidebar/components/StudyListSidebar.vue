<template>
    <div
        class="flex flex-col h-full"
        :style="{ width: 'var(--custom-study-list-width)' }"
    >
        <!-- 최상단 홈 아이콘 바 -->
        <div
            class="flex items-center justify-center w-full h-[var(--custom-notice-bar-height)] gap-2"
        >
            <div
                class="relative flex items-center justify-center w-full group cursor-pointer"
                @click="handleTitleIconClick"
            >
                <div
                    :class="[
                        'absolute left-0 top-1/2 -translate-y-1/2 rounded transition-all duration-300',
                        isTitleActive()
                            ? 'w-1 h-[70%] bg-gray-50'
                            : 'w-0 group-hover:w-1 h-[50%] bg-gray-150',
                    ]"
                ></div>
                <div
                    class="w-10 h-10 mx-2 my-1 overflow-hidden border-2 rounded-lg"
                >
                    <img
                        src="/logo.svg"
                        alt="logo"
                        class="w-full h-full bg-gray-50"
                    />
                </div>
            </div>
        </div>

        <!-- 스크롤 영역 wrappers -->
        <div class="scroll-wrapper relative flex-1 min-h-0">
            <!-- 커스텀 트랙 (아이콘 뒤) -->
            <div class="custom-scrollbar-track"></div>

            <!-- 실제 스크롤 컨테이너 (native scrollbar 숨김) -->
            <div ref="scrollContainer" class="study-list-scroll">
                <div class="flex flex-col items-center gap-2 w-full">
                    <StudyIconList />
                    <div
                        class="relative flex items-center justify-center w-full group cursor-pointer"
                        @click="handleNewStudyClick"
                    >
                        <div
                            class="w-10 h-10 overflow-hidden border-2 rounded-lg flex items-center justify-center p-1 dark:bg-gray-50 hover:bg-gray-200"
                        >
                            <span class="mdi mdi-plus-circle text-2xl"></span>
                        </div>
                    </div>
                </div>
            </div>

            <!-- 커스텀 썸 (아이콘 위) -->
            <div ref="thumb" class="custom-scrollbar-thumb"></div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { PATHS } from '@/router/paths';
import { useAppStore } from '@/store';
import { StudyIconList } from '@/entities';

const router = useRouter();
const route = useRoute();
const appStore = useAppStore();

const scrollContainer = ref<HTMLElement | null>(null);
const thumb = ref<HTMLElement | null>(null);

function handleTitleIconClick() {
    router.push({ path: PATHS.ROOT });
}
const isTitleActive = () => route.path === `/${PATHS.ME}`;
function handleNewStudyClick() {
    appStore.saveLastRoute(route.path as string);
    router.push(`/${PATHS.CREATE}`);
}

// 썸 크기/위치 업데이트
function updateThumb() {
    if (!scrollContainer.value || !thumb.value) return;

    const sc = scrollContainer.value;
    const th = thumb.value;
    const viewportH = sc.clientHeight;
    const contentH = sc.scrollHeight;

    const ratio = viewportH / contentH;
    const thumbH = Math.max(viewportH * ratio, 20);
    th.style.height = `${thumbH}px`;

    const maxScroll = contentH - viewportH;
    const trackH = viewportH - thumbH;
    const top = (sc.scrollTop / maxScroll) * trackH;
    th.style.transform = `translateY(${top}px)`;
}

onMounted(() => {
    if (scrollContainer.value) {
        scrollContainer.value.addEventListener('scroll', updateThumb);
        updateThumb();
    }
});
onBeforeUnmount(() => {
    if (scrollContainer.value) {
        scrollContainer.value.removeEventListener('scroll', updateThumb);
    }
});
</script>

<style scoped lang="scss">
/* 스크롤바는 hover 시에만 보여줌 */
.scroll-wrapper .custom-scrollbar-track,
.scroll-wrapper .custom-scrollbar-thumb {
    opacity: 0;
    transition: opacity 0.3s ease;
}
.scroll-wrapper:hover .custom-scrollbar-track,
.scroll-wrapper:hover .custom-scrollbar-thumb {
    opacity: 1;
}

/* 실제 스크롤 컨테이너 */
.study-list-scroll {
    height: 100%;
    overflow-y: auto;
    scrollbar-width: none;
    -ms-overflow-style: none;
    padding-right: 4px; /* 썸/트랙 두께 대비 여유 확보 */
}
.study-list-scroll::-webkit-scrollbar {
    display: none;
}

/* 커스텀 트랙: 배경 제거 */
.custom-scrollbar-track {
    position: absolute;
    top: var(--custom-notice-bar-height);
    bottom: 0;
    right: 0;
    width: 4px;
    background-color: transparent;
    pointer-events: none;
    z-index: 0;
}

/* 커스텀 썸 */
.custom-scrollbar-thumb {
    position: absolute;
    top: 0;
    right: 0;
    width: 4px;
    background-color: var(--custom-scrollbar-color, rgba(255, 255, 255, 0.4));
    border-radius: 4px;
    pointer-events: none;
    z-index: 1;
}
</style>
