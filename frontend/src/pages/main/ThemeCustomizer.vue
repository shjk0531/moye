<!-- ThemeCustomizer.vue -->
<template>
    <div class="p-4 space-y-6">
        <h1 class="text-2xl font-bold">테마 색상 세부 설정</h1>

        <div class="flex items-center space-x-4">
            <label class="font-semibold">Primary</label>
            <ColorPicker
                v-model="colors.primary"
                @input="updateTheme('primary', colors.primary)"
            />
        </div>
        <div class="flex items-center space-x-4">
            <label class="font-semibold">Secondary</label>
            <ColorPicker
                v-model="colors.secondary"
                @input="updateTheme('secondary', colors.secondary)"
            />
        </div>
        <div class="flex items-center space-x-4">
            <label class="font-semibold">Accent</label>
            <ColorPicker
                v-model="colors.accent"
                @input="updateTheme('accent', colors.accent)"
            />
        </div>
        <div class="flex items-center space-x-4">
            <label class="font-semibold">Text</label>
            <ColorPicker
                v-model="colors.text"
                @input="updateTheme('text', colors.text)"
            />
        </div>
        <div class="flex items-center space-x-4">
            <label class="font-semibold">Background</label>
            <ColorPicker
                v-model="colors.background"
                @input="updateTheme('background', colors.background)"
            />
        </div>

        <div
            class="mt-8 p-4 rounded"
            :style="{ backgroundColor: colors.background, color: colors.text }"
        >
            이 영역은 선택한 배경색과 텍스트 색상이 적용됩니다.
        </div>
    </div>
</template>

<script setup>
import { reactive } from 'vue';
import ColorPicker from 'primevue/colorpicker';

// 색상 상태를 reactive 객체로 관리
const colors = reactive({
    primary: '#3490dc',
    secondary: '#ffed4a',
    accent: '#e3342f',
    text: '#333333',
    background: '#ffffff',
});

// 각 항목에 해당하는 CSS 변수를 업데이트하는 함수
const updateTheme = (key, value) => {
    document.documentElement.style.setProperty(`--color-${key}`, value);
};
</script>
