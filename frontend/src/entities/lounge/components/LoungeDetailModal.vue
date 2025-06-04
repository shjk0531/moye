<!-- src/entities/lounge/components/LoungeDetailModal.vue -->
<template>
    <teleport to="body">
        <!-- 모달 오버레이 -->
        <div
            class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
            @click.self="close"
        >
            <!-- 모달 박스: 최대 높이를 90vh로 제한, flex-col으로 내부 레이아웃 제어 -->
            <div
                class="relative bg-gray-50 dark:bg-gray-800 text-gray-900 dark:text-gray-100 rounded-lg w-11/12 sm:w-3/4 md:w-2/3 lg:w-1/2 xl:w-2/5 max-h-[90vh] flex flex-col shadow-2xl"
            >
                <!-- 닫기 버튼(우상단) -->
                <button
                    class="absolute top-3 right-3 text-gray-500 hover:text-gray-200 z-10 cursor-pointer"
                    @click="close"
                >
                    ✕
                </button>

                <!-- 모달 콘텐츠: flex-1으로 남은 공간 채우기, min-h-0 꼭 필요 -->
                <div class="flex flex-1 min-h-0">
                    <!-- 왼쪽: 라운지 상세 정보 (스크롤 가능) -->
                    <div class="w-2/3 p-6 overflow-y-auto">
                        <!-- 라운지 이름 -->
                        <h2
                            class="text-2xl font-bold text-gray-900 dark:text-gray-100 mb-4"
                        >
                            {{ lounge.name }}
                        </h2>

                        <!-- 라운지 프로필 이미지 (상단에 크게 노출) -->
                        <img
                            :src="lounge.profile_url"
                            alt="lounge-profile"
                            class="w-full h-48 object-cover rounded-md mb-4"
                        />

                        <!-- 상세 설명 -->
                        <div class="prose prose-sm dark:prose-invert">
                            <p>{{ lounge.description }}</p>
                            <!-- 예시로 설명이 길다고 가정할 때, 아래에 더미 텍스트를 추가해보았습니다 -->
                            <p v-if="!lounge.description">
                                이 라운지는 여러분의 열정과 아이디어를 연결하는
                                공간입니다. 다양한 주제의 토론, 경험 공유,
                                프로젝트 협업 등을 통해 서로에게 영감을 주며
                                성장할 수 있습니다. 라운지 내부에서는 자유롭게
                                의견을 나눌 수 있으며, 정기적인 온라인 모임과
                                오프라인 워크샵을 통해 실제 네트워킹까지
                                확장됩니다.
                                <br /><br />
                                “함께 도전하고, 함께 성장한다”는 슬로건 아래,
                                새로운 사람들과의 만남을 기대해 보세요. 멤버들과
                                지식을 나누며 시너지를 만들어가는 순간, 여러분이
                                속한 라운지는 단순한 온라인 커뮤니티를 넘어
                                진정한 ‘성장의 장’이 될 것입니다.
                            </p>
                        </div>

                        <!-- 업데이트 날짜 (스크롤 영역 끝) -->
                        <p
                            class="text-xs text-gray-500 dark:text-gray-400 mt-6"
                        >
                            업데이트: {{ formattedDate }}
                        </p>
                    </div>

                    <!-- 오른쪽: 리더 정보, 멤버 수, 태그, 지원 버튼 (독립적으로 고정) -->
                    <div
                        class="w-1/3 border-l border-gray-200 dark:border-gray-700 p-6 flex flex-col"
                    >
                        <!-- 리더 정보 -->
                        <div class="flex items-center mb-6">
                            <!-- 프로필 이미지(하드코딩 예시) -->
                            <img
                                src="https://via.placeholder.com/80"
                                alt="leader-avatar"
                                class="w-16 h-16 rounded-full object-cover mr-4"
                            />
                            <div>
                                <p
                                    class="text-lg font-semibold text-gray-900 dark:text-gray-100"
                                >
                                    LeaderName
                                </p>
                                <p
                                    class="text-sm text-gray-500 dark:text-gray-400"
                                >
                                    닉네임
                                </p>
                            </div>
                        </div>

                        <!-- 멤버 수 (하드코딩 예시) -->
                        <div class="mb-6">
                            <p class="text-sm text-gray-600 dark:text-gray-300">
                                멤버 수:
                                <span
                                    class="font-medium text-gray-900 dark:text-gray-100"
                                >
                                    10명
                                </span>
                            </p>
                        </div>

                        <!-- 태그 -->
                        <div class="mb-6">
                            <p
                                class="text-sm font-medium text-gray-700 dark:text-gray-300 mb-2"
                            >
                                태그
                            </p>
                            <div class="flex flex-wrap gap-2">
                                <span
                                    v-for="(tag, idx) in lounge.tags"
                                    :key="idx"
                                    class="text-xs bg-gray-200 text-gray-800 dark:bg-gray-700 dark:text-gray-200 px-3 py-1 rounded-full"
                                >
                                    {{ tag }}
                                </span>
                            </div>
                        </div>

                        <!-- 지원하기 버튼 -->
                        <button
                            class="mt-auto w-full bg-indigo-600 hover:bg-indigo-500 text-white font-semibold py-2 px-4 rounded-lg transition-colors duration-200 cursor-pointer"
                        >
                            지원하기
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </teleport>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { type LoungeInfo } from '@/entities';

const props = defineProps<{
    lounge: LoungeInfo;
}>();

const emit = defineEmits<{
    (e: 'close'): void;
}>();

const close = () => {
    emit('close');
};

const formattedDate = computed(() => {
    if (!props.lounge.updated_at) return '';
    const d = new Date(props.lounge.updated_at);
    return `${d.getFullYear()}-${String(d.getMonth() + 1).padStart(
        2,
        '0',
    )}-${String(d.getDate()).padStart(2, '0')}`;
});
</script>
