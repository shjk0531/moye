<template>
    <div class="space-y-4">
        <div class="flex flex-col gap-2">
            <label class="block mb-1 font-medium dark:text-gray-200"
                >해시태그 추가 (엔터키로 추가)</label
            >
            <input
                type="text"
                v-model="inputTag"
                @keyup.enter.prevent="addTag"
                class="w-full p-2 border border-gray-300 rounded dark:bg-gray-700 dark:border-gray-600 dark:text-gray-200 focus:outline-none dark:focus:ring-blue-500"
                placeholder="해시태그를 입력하고 Enter를 누르세요"
            />
            <div class="flex flex-wrap gap-2 mb-2">
                <span
                    v-for="(tag, index) in props.hashes"
                    :key="index"
                    @click="removeTag(tag)"
                    class="inline-flex items-center bg-blue-500 text-white text-sm rounded-full px-3 py-1 hover:bg-blue-600 cursor-pointer"
                >
                    {{ tag }}
                </span>
            </div>
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
import { computed, ref } from 'vue';

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

const inputTag = ref('');

const addTag = () => {
    if (inputTag.value.trim()) {
        const newTag = inputTag.value.trim();
        const newHashes = [...props.hashes, newTag];
        emit('update:hashes', newHashes);
        inputTag.value = '';
    }
};

const removeTag = (tag: string) => {
    const newHashes = props.hashes.filter((t) => t !== tag);
    emit('update:hashes', newHashes);
};
</script>

<style lang="scss">
/* 추가 스타일이 필요하면 작성 */
</style>
