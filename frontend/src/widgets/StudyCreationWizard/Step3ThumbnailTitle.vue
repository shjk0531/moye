<template>
    <div class="space-y-4">
        <div>
            <label class="block mb-1 font-medium">제목</label>
            <input
                type="text"
                v-model="title"
                class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                placeholder="제목을 입력하세요"
            />
        </div>
        <div>
            <label class="block mb-1 font-medium">썸네일 URL</label>
            <input
                type="text"
                v-model="thumbnailUrl"
                class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                placeholder="이미지 URL을 입력하세요"
            />
            <div v-if="thumbnailUrl" class="mt-2">
                <img
                    :src="thumbnailUrl"
                    alt="썸네일 미리보기"
                    class="max-w-full h-auto rounded"
                />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';

const props = defineProps<{ title: string; thumbnailUrl: string }>();
const emit = defineEmits<{
    (e: 'update:title', val: string): void;
    (e: 'update:thumbnailUrl', val: string): void;
}>();

const title = computed<string>({
    get: () => props.title,
    set: (val: string) => emit('update:title', val),
});

const thumbnailUrl = computed<string>({
    get: () => props.thumbnailUrl,
    set: (val: string) => emit('update:thumbnailUrl', val),
});
</script>

<style lang="scss">
/* 추가 스타일 필요한 경우 작성 */
</style>
