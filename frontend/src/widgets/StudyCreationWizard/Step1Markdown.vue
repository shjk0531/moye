<template>
    <div class="flex flex-col h-full">
        <label class="block mb-2 font-semibold">본문 작성</label>
        <div class="flex-1 overflow-auto">
            <MarkdownEditor v-model="localValue" />
        </div>
    </div>
</template>

<script setup lang="ts">
import { defineProps, defineEmits, computed } from 'vue';
import { MarkdownEditor } from '@/shared/ui';

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

<style lang="scss">
/* 추가적인 스타일이 필요하면 작성 */
</style>
