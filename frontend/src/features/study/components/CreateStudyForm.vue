<template>
    <div class="create-study-form p-4">
        <h2 class="text-2xl font-bold mb-6 dark:text-gray-100">
            새 스터디 생성
        </h2>

        <form @submit.prevent="onSubmit" class="space-y-6">
            <!-- 스터디 이름 -->
            <div class="form-group">
                <label class="block mb-2 font-medium dark:text-gray-200"
                    >스터디 이름</label
                >
                <input
                    v-model="form.name"
                    type="text"
                    class="w-full p-2 border border-gray-300 rounded focus:outline-none dark:bg-gray-700 dark:border-gray-600 dark:text-gray-200"
                    placeholder="스터디 이름을 입력하세요"
                    required
                />
            </div>

            <!-- 프로필 이미지 URL -->
            <div class="form-group">
                <label class="block mb-2 font-medium dark:text-gray-200"
                    >프로필 이미지 URL</label
                >
                <input
                    v-model="form.profileUrl"
                    type="text"
                    class="w-full p-2 border border-gray-300 rounded focus:outline-none dark:bg-gray-700 dark:border-gray-600 dark:text-gray-200"
                    placeholder="이미지 URL을 입력하세요"
                    required
                />
                <div v-if="form.profileUrl" class="mt-2">
                    <img
                        :src="form.profileUrl"
                        alt="프로필 이미지 미리보기"
                        class="max-w-full h-32 object-cover rounded"
                    />
                </div>
            </div>

            <!-- 소개글 -->
            <div class="form-group">
                <label class="block mb-2 font-medium dark:text-gray-200"
                    >소개글</label
                >
                <textarea
                    v-model="form.description"
                    class="w-full p-2 border border-gray-300 rounded h-24 resize-none dark:bg-gray-700 dark:border-gray-600 dark:text-gray-200"
                    placeholder="스터디에 대한 소개글을 입력하세요"
                    required
                ></textarea>
            </div>

            <!-- 본문 내용 -->
            <div class="form-group">
                <label class="block mb-2 font-medium dark:text-gray-200"
                    >본문 내용</label
                >
                <textarea
                    v-model="form.content"
                    class="w-full p-2 border border-gray-300 rounded h-32 resize-none dark:bg-gray-700 dark:border-gray-600 dark:text-gray-200"
                    placeholder="스터디에 대한 상세 내용을 입력하세요"
                    required
                ></textarea>
            </div>

            <!-- 태그 -->
            <div class="form-group">
                <label class="block mb-2 font-medium dark:text-gray-200"
                    >태그 (쉼표로 구분)</label
                >
                <input
                    v-model="tagsInput"
                    type="text"
                    class="w-full p-2 border border-gray-300 rounded focus:outline-none dark:bg-gray-700 dark:border-gray-600 dark:text-gray-200"
                    placeholder="예: javascript, vue, react"
                />
                <div
                    v-if="form.tags.length > 0"
                    class="mt-2 flex flex-wrap gap-2"
                >
                    <span
                        v-for="(tag, index) in form.tags"
                        :key="index"
                        class="px-2 py-1 bg-blue-100 text-blue-800 rounded dark:bg-blue-900 dark:text-blue-200"
                    >
                        {{ tag }}
                    </span>
                </div>
            </div>

            <!-- 버튼 -->
            <div class="flex justify-end space-x-4">
                <button
                    type="button"
                    @click="onCancel"
                    class="px-4 py-2 bg-gray-200 rounded hover:bg-gray-300 dark:bg-gray-700 dark:hover:bg-gray-600 dark:text-gray-200"
                >
                    취소
                </button>
                <button
                    type="submit"
                    class="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 dark:bg-blue-600 dark:hover:bg-blue-700"
                    :disabled="isSubmitting"
                >
                    {{ isSubmitting ? '생성 중...' : '스터디 생성' }}
                </button>
            </div>
        </form>
    </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, watch } from 'vue';
import type { StudyCreatePayload } from '@/entities/study/models/types';

const emit = defineEmits<{
    (e: 'cancel'): void;
    (e: 'submit', payload: StudyCreatePayload): void;
}>();

const form = reactive<StudyCreatePayload>({
    name: '',
    profileUrl: '',
    description: '',
    content: '',
    tags: [],
});

const tagsInput = ref('');
const isSubmitting = ref(false);

// 태그 입력 감시
watch(tagsInput, (newValue) => {
    if (newValue) {
        form.tags = newValue
            .split(',')
            .map((tag) => tag.trim())
            .filter((tag) => tag !== '');
    } else {
        form.tags = [];
    }
});

function onCancel() {
    emit('cancel');
}

function onSubmit() {
    isSubmitting.value = true;

    try {
        emit('submit', { ...form });
    } finally {
        isSubmitting.value = false;
    }
}
</script>

<style lang="scss" scoped>
.create-study-form {
    max-width: 800px;
    margin: 0 auto;
}
</style>
