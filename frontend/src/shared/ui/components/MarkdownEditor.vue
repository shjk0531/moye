<template>
    <textarea
        v-model="value"
        @input="onInput"
        @keydown="onKeydown"
        class="w-full h-full p-3 resize-none focus:outline-none text-gray-800 dark:text-gray-200"
        placeholder="소개글 입력..."
    ></textarea>
</template>

<script setup lang="ts">
import { computed } from 'vue';

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

<style lang="scss"></style>
