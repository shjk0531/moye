<template>
    <div class="study-creation-wizard flex flex-col p-6 h-screen w-full">
        <div class="flex items-center mb-6">
            <div class="flex-1 flex items-center space-x-2 gap-x-2">
                <button
                    @click="handleCancel"
                    class="px-4 py-1 bg-gray-200 rounded hover:bg-gray-300 dark:bg-red-500 dark:hover:bg-red-600 dark:text-gray-50 cursor-pointer"
                >
                    나가기
                </button>
                <Stepper
                    :steps="stepLabels"
                    :modelValue="currentStep"
                    :backward="isBackward"
                    @update:modelValue="goToStep"
                    class="flex-1"
                />
            </div>
            <div class="flex space-x-2">
                <button
                    :disabled="currentStep === 0"
                    @click="prevStep"
                    class="px-4 py-1 bg-gray-200 rounded hover:bg-gray-300 dark:bg-gray-500 dark:hover:bg-gray-700 dark:text-gray-50 cursor-pointer disabled:opacity-100 disabled:cursor-not-allowed"
                >
                    이전
                </button>
                <button
                    v-if="currentStep === stepLabels.length - 1"
                    @click="handleCreate"
                    :disabled="!isCreateButtonEnabled"
                    class="px-4 py-1 bg-green-500 text-white rounded hover:bg-green-600 dark:bg-green-600 dark:hover:bg-green-700 cursor-pointer disabled:opacity-50 disabled:cursor-not-allowed"
                >
                    생성
                </button>
                <button
                    v-else
                    @click="nextStep"
                    :disabled="!canGoToNextStep"
                    class="px-4 py-1 bg-gray-200 text-white rounded hover:bg-gray-300 dark:bg-gray-50 dark:hover:bg-gray-150 dark:text-gray-950 cursor-pointer disabled:opacity-50 disabled:cursor-not-allowed"
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
        <!-- 유효성 검사 오류 메시지 -->
        <div v-if="validationError" class="mt-4 text-red-500 text-sm">
            {{ validationError }}
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue';
import { Stepper } from '@/shared/ui';
import Step1Markdown from './Step1Markdown.vue';
import Step2Tags from './Step2Tags.vue';
import Step3ThumbnailTitle from './Step3ThumbnailTitle.vue';

const emit = defineEmits<{
    (e: 'cancel'): void;
    (
        e: 'create',
        payload: {
            content: string;
            tags: string[];
            intro: string;
            title: string;
            thumbnailUrl: string;
        },
    ): void;
}>();

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
const validationError = ref('');

// 유효성 검사 함수
function validateCurrentStep(): boolean {
    validationError.value = '';

    switch (currentStep.value) {
        case 0: // 본문 작성
            if (!form.content.trim()) {
                validationError.value = '본문 내용을 입력해주세요.';
                return false;
            }
            return true;
        case 1: // 해시태그, 소개
            if (!form.intro.trim()) {
                validationError.value = '소개글을 입력해주세요.';
                return false;
            }
            return true;
        case 2: // 제목, 썸네일
            if (!form.title.trim()) {
                validationError.value = '제목을 입력해주세요.';
                return false;
            }
            if (!form.thumbnailUrl.trim()) {
                validationError.value = '썸네일 URL을 입력해주세요.';
                return false;
            }
            return true;
        default:
            return true;
    }
}

// 다음 단계로 이동할 수 있는지 확인
const canGoToNextStep = computed(() => {
    return validateCurrentStep();
});

// 생성 버튼 활성화 여부
const isCreateButtonEnabled = computed(() => {
    return (
        form.content.trim() !== '' &&
        form.intro.trim() !== '' &&
        form.title.trim() !== '' &&
        form.thumbnailUrl.trim() !== ''
    );
});

function handleCancel() {
    emit('cancel');
}

function handleCreate() {
    if (!validateCurrentStep()) return;

    emit('create', {
        content: form.content,
        tags: form.hashes,
        intro: form.intro,
        title: form.title,
        thumbnailUrl: form.thumbnailUrl,
    });
}

function goToStep(idx: number) {
    if (isAnimating.value || idx === currentStep.value) return;

    // 다음 스텝으로 이동할 때만 유효성 검사
    if (idx > currentStep.value && !validateCurrentStep()) {
        return;
    }

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
    if (currentStep.value < stepLabels.length - 1 && validateCurrentStep()) {
        goToStep(currentStep.value + 1);
    }
}
</script>

<style scoped lang="scss"></style>
