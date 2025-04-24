<template>
    <div class="stepper" :class="{ backward: backward }">
        <div
            v-for="(label, idx) in steps"
            :key="idx"
            class="step flex items-center flex-1"
        >
            <div
                class="step-circle cursor-pointer"
                :class="{ active: modelValue >= idx }"
                @click="handleClick(idx)"
            >
                {{ idx + 1 }}
            </div>
            <div
                class="step-label ml-2 text-sm text-gray-700 cursor-pointer"
                @click="handleClick(idx)"
            >
                {{ label }}
            </div>
            <div v-if="idx < steps.length - 1" class="step-line flex-1 mx-2">
                <div
                    :class="['step-line-fill', { active: modelValue > idx }]"
                ></div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { defineProps, defineEmits } from 'vue';
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

<style scoped lang="scss">
.stepper {
    display: flex;
    align-items: center;
}
.step {
    display: flex;
    align-items: center;
    flex: 1;
}
.step-circle {
    width: 32px;
    height: 32px;
    border: 2px solid #d1d5db;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    color: #6b7280;
    background-color: #ffffff;
    transition: background-color 0.3s ease 0.3s, border-color 0.3s ease 0.3s,
        color 0.3s ease 0.3s;
}
.step-circle.active {
    background-color: #3b82f6;
    border-color: #3b82f6;
    color: #ffffff;
    transition-delay: 0.3s;
}
.step-line {
    position: relative;
    flex: 1;
    height: 2px;
    background-color: #d1d5db;
    overflow: hidden;
}
.step-line-fill {
    position: absolute;
    top: 0;
    left: 0;
    height: 100%;
    background-color: #3b82f6;
    width: 0;
    transition: width 0.3s ease;
}
.step-line-fill.active {
    width: 100%;
}
.step-label {
    font-size: 0.875rem;
}
.stepper.backward .step-circle {
    transition-delay: 0s !important;
}
.stepper.backward .step-line-fill {
    transition: width 0.3s ease 0s;
}
</style>
