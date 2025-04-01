<!-- src/widgets/titlebar/components/Titlebar.vue -->
<template>
    <div class="titlebar flex flex-row items-center justify-center">
        <div class="flex flex-row items-center justify-center">
            <div class="titleIcon bg-gray-400 w-2 h-2">
                <!-- 스터디 아이콘이 존재하면 이미지를 렌더링 -->
                <img
                    v-if="studyIcon"
                    :src="studyIcon"
                    alt="study icon"
                    class="w-full h-full"
                />
            </div>
            <div class="text-gray-200 text-2xs">
                {{ studyName || '타이틀' }}
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'Titlebar',
    computed: {
        studyName() {
            return this.$globalState.studyName;
        },
        studyIcon() {
            return this.$globalState.studyIcon;
        },
    },
};
</script>

<style scoped>
.titlebar {
    grid-area: titlebar;
}
</style>
