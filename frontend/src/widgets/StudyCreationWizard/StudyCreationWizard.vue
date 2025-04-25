<template>
    <div class="study-creation-wizard p-6 h-full w-full flex flex-col">
        <div class="flex justify-between items-center mb-6">
            <!-- 스텝 표시 -->
            <Stepper
                :steps="stepLabels"
                :modelValue="currentStep"
                :backward="isBackward"
                @update:modelValue="goToStep"
                class="flex-1"
            />
            <!-- 이전/다음 버튼 -->
            <div class="space-x-2">
                <button
                    @click="prevStep"
                    :disabled="currentStep === 0"
                    class="px-4 py-1 bg-gray-200 rounded hover:bg-gray-300 disabled:opacity-50"
                >
                    이전
                </button>
                <button
                    @click="nextStep"
                    :disabled="currentStep === stepLabels.length - 1"
                    class="px-4 py-1 bg-blue-500 text-white rounded hover:bg-blue-600 disabled:opacity-50"
                >
                    다음
                </button>
            </div>
        </div>
        <!-- 각 스텝 폼 -->
        <div class="flex-1 overflow-auto">
            <Step1Markdown v-if="currentStep === 0" v-model="form.content" />
            <Step2Tags
                v-else-if="currentStep === 1"
                v-model:hashes="form.hashes"
                v-model:intro="form.intro"
            />
            <Step3ThumbnailTitle
                v-else
                v-model:title="form.title"
                v-model:thumbnailUrl="form.thumbnailUrl"
            />
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue';
import { Stepper } from '@/shared/ui';
import Step1Markdown from './Step1Markdown.vue';
import Step2Tags from './Step2Tags.vue';
import Step3ThumbnailTitle from './Step3ThumbnailTitle.vue';

const stepLabels = ['본문 작성', '해시태그・소개', '제목・썸네일'];
const currentStep = ref(0);
const form = reactive({
    content: '',
    hashes: [] as string[],
    intro: '',
    title: '',
    thumbnailUrl: '',
});
const isAnimating = ref(false);
const isBackward = ref(false);

function goToStep(idx: number) {
    if (isAnimating.value || idx === currentStep.value) return;
    const start = currentStep.value;
    const stepCount = Math.abs(idx - start);
    const direction = idx > start ? 1 : -1;
    if (stepCount > 1) {
        isBackward.value = direction < 0;
        isAnimating.value = true;
        let count = 0;
        const animateStep = () => {
            count++;
            currentStep.value = start + direction * count;
            if (count < stepCount) {
                setTimeout(animateStep, 600);
            } else {
                isAnimating.value = false;
                isBackward.value = false;
            }
        };
        animateStep();
    } else {
        if (direction < 0) {
            // 단일 스텝 backward 이동: 원 먼저 회색, 0.3s 후 선 애니메이션, 600ms 후 backward 플래그 해제
            isBackward.value = true;
            currentStep.value = idx;
            setTimeout(() => {
                isBackward.value = false;
            }, 600);
        } else {
            currentStep.value = idx;
        }
    }
}
function prevStep() {
    if (currentStep.value > 0) {
        goToStep(currentStep.value - 1);
    }
}
function nextStep() {
    if (currentStep.value < stepLabels.length - 1) {
        goToStep(currentStep.value + 1);
    }
}
</script>

<style scoped lang="scss"></style>
