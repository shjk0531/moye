<template>
    <div class="space-y-4">
        <div>
            <label class="block mb-1 font-medium dark:text-gray-200"
                >해시태그 (콤마로 구분)</label
            >
            <input
                type="text"
                v-model="hashtagsStr"
                class="w-full p-2 border border-gray-300 rounded dark:bg-gray-700 dark:border-gray-600 dark:text-gray-200 dark:placeholder-gray-400 focus:outline-none dark:focus:ring-blue-500"
                placeholder="예: vue,typescript,tailwind"
            />
        </div>
        <div>
            <label class="block mb-1 font-medium dark:text-gray-200"
                >소개글</label
            >
            <textarea
                v-model="introStr"
                class="w-full p-2 border border-gray-300 rounded h-24 resize-none dark:bg-gray-700 dark:border-gray-600 dark:text-gray-200 dark:placeholder-gray-400 focus:outline-none dark:focus:ring-blue-500"
                placeholder="소개글을 입력하세요"
            ></textarea>
        </div>
    </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';

const props = defineProps<{ hashes: string[]; intro: string }>();
const emit = defineEmits<{
    (e: 'update:hashes', val: string[]): void;
    (e: 'update:intro', val: string): void;
}>();

// computed wrapper for v-model:hashes
const hashtagsStr = computed<string>({
    get: () => props.hashes.join(', '),
    set: (val: string) => {
        const arr = val
            .split(',')
            .map((s) => s.trim())
            .filter(Boolean);
        emit('update:hashes', arr);
    },
});

// computed wrapper for v-model:intro
const introStr = computed<string>({
    get: () => props.intro,
    set: (val: string) => emit('update:intro', val),
});
</script>

<style lang="scss">
/* 추가 스타일이 필요하면 작성 */
</style>
