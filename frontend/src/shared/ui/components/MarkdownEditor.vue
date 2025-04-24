<template>
    <div class="markdown-editor flex h-full">
        <div class="markdown-editor-input w-1/2 p-4 bg-gray-50">
            <textarea
                v-model="value"
                @input="onInput"
                @keydown="onKeydown"
                class="w-full h-full resize-none p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                placeholder="Markdown 입력..."
            ></textarea>
        </div>
        <div
            class="markdown-editor-preview w-1/2 p-4 bg-white overflow-auto markdown-body"
        >
            <MarkdownPreview :content="value" />
        </div>
    </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import MarkdownPreview from './MarkdownPreview.vue';

const props = defineProps({
    modelValue: {
        type: String,
        default: '',
    },
});
const emit = defineEmits<{
    (e: 'update:modelValue', value: string): void;
}>();

const value = computed<string>({
    get: () => props.modelValue,
    set: (val: string) => emit('update:modelValue', val),
});

function onInput(event: Event) {
    const target = event.target as HTMLTextAreaElement;
    value.value = target.value;
}

function onKeydown(event: KeyboardEvent) {
    if (event.key === 'Tab') {
        event.preventDefault();
        const target = event.target as HTMLTextAreaElement;
        const start = target.selectionStart;
        const end = target.selectionEnd;
        target.setRangeText('\t', start, end, 'end');
        onInput(event);
    }
}
</script>

<style lang="scss">
.markdown-editor {
    height: 100%;
}
</style>
