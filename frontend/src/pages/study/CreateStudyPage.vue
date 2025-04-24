<template>
    <div
        class="create-study-page flex justify-center bg-gray-100 h-screen w-screen overflow-hidden"
    >
        <StudyCreationWizard />
    </div>
</template>

<script setup lang="ts">
import { StudyCreationWizard } from '@/widgets/StudyCreationWizard';
</script>

<style lang="scss">
/* 필요 시 추가 스타일을 작성하세요 */
</style>
