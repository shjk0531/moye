<!-- src/entities/lounge/components/LoungeCard.vue -->
<template>
    <div
        class="flex flex-col bg-gray-800 text-gray-100 border border-gray-700 rounded-lg overflow-hidden h-full transition-transform duration-200 ease-in-out transform hover:-translate-y-1 hover:shadow-lg hover:shadow-black/40 cursor-pointer"
        @click="handleClick"
    >
        <!-- 이미지 (고정 높이, 전체 너비) -->
        <img
            :src="lounge.profile_url"
            alt="lounge-profile"
            class="w-full h-40 object-cover"
        />
        <!-- 텍스트 영역 -->
        <div class="p-4 flex-1 flex flex-col">
            <h3 class="text-lg font-semibold mb-2">{{ lounge.name }}</h3>
            <p class="text-sm text-gray-300 line-clamp-2 mb-2">
                {{ lounge.description }}
            </p>
            <div class="flex flex-wrap gap-2 mb-2">
                <span
                    v-for="(tag, idx) in lounge.tags"
                    :key="idx"
                    class="text-xs bg-gray-700 px-2 py-1 rounded-full"
                >
                    {{ tag }}
                </span>
            </div>
            <p class="text-xs text-gray-400 mt-auto">{{ formattedDate }}</p>
        </div>
    </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { type LoungeInfo } from '@/entities';

const props = defineProps<{
    lounge: LoungeInfo;
}>();

// "select" 이벤트를 상위로 올리기 위한 정의
const emit = defineEmits<{
    (e: 'select', payload: LoungeInfo): void;
}>();

// 클릭 시 호출
const handleClick = () => {
    emit('select', props.lounge);
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
