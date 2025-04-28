<template>
    <div class="flex items-center">
        <div
            v-for="(label, idx) in steps"
            :key="idx"
            class="flex items-center flex-1"
        >
            <div
                :class="[
                    'w-8 h-8 rounded-full flex items-center justify-center transition-colors duration-300 cursor-pointer',
                    backward ? 'delay-0' : 'delay-300',
                    modelValue >= idx
                        ? 'bg-green-500 text-white dark:bg-green-600'
                        : 'bg-white text-gray-500 dark:bg-gray-700 dark:text-gray-400',
                ]"
                @click="handleClick(idx)"
            >
                {{ idx + 1 }}
            </div>
            <div
                class="ml-2 text-sm text-gray-700 cursor-pointer dark:text-gray-50"
                @click="handleClick(idx)"
            >
                {{ label }}
            </div>
            <div
                v-if="idx < steps.length - 1"
                class="relative flex-1 mx-2 h-0.5 bg-gray-300 dark:bg-gray-700 overflow-hidden"
            >
                <div
                    :class="[
                        'absolute top-0 left-0 h-full bg-green-500 transition-all duration-300 ease-in-out dark:bg-green-600',
                        modelValue > idx ? 'w-full' : 'w-0',
                    ]"
                ></div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import type { PropType } from 'vue';

const props = defineProps({
    steps: { type: Array as PropType<string[]>, required: true },
    modelValue: { type: Number, required: true },
    backward: { type: Boolean, default: false },
});

const emit = defineEmits<{ (e: 'update:modelValue', value: number): void }>();

function handleClick(idx: number) {
    emit('update:modelValue', idx);
}
</script>
