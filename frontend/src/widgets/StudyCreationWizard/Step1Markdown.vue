<template>
    <div class="flex flex-col h-full">
        <label class="block mb-2 font-semibold text-gray-800 dark:text-gray-50"
            >본문 작성</label
        >
        <div class="flex-1 overflow-auto">
            <div
                class="markdown-editor flex h-full border dark:border-gray-800 rounded"
            >
                <div class="w-1/2 overflow-hidden bg-gray-100 dark:bg-gray-800">
                    <MarkdownEditor v-model="localValue" />
                </div>
                <!-- 구분선  -->
                <div class="border-1 dark:border-gray-700"></div>
                <div class="w-1/2 overflow-auto bg-gray-100 dark:bg-gray-800">
                    <MarkdownPreview :content="localValue" />
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { MarkdownEditor, MarkdownPreview } from '@/shared/ui';

const props = defineProps<{ modelValue: string }>();
const emit = defineEmits<{
    (e: 'update:modelValue', val: string): void;
}>();

// prop을 업데이트할 수 있는 computed 래퍼
const localValue = computed<string>({
    get: () => props.modelValue,
    set: (val: string) => emit('update:modelValue', val),
});
</script>

<style lang="scss"></style>
