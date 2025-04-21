<template>
    <div class="relative flex flex-row h-screen w-screen dark:bg-gray-800">
        <div class="flex w-(--custom-side-bar-left-width) dark:bg-gray-950">
            <SettingListSidebar />
        </div>
        <div class="flex flex-1">
            <router-view />
        </div>
        <div class="flex w-fix p-10">
            <span
                class="mdi mdi-close-circle text-4xl text-gray-150 hover:text-gray-200 cursor-pointer"
                @click="handleClose"
            ></span>
        </div>
    </div>
</template>

<script setup lang="ts">
import { SettingListSidebar } from '@/widgets/sidebar';
import { useAppStore } from '@/store';
import { useRouter } from 'vue-router';

const router = useRouter();

const appStore = useAppStore();

const handleClose = () => {
    router.push(appStore.getLastRoute);
};
</script>

<style scoped lang="scss"></style>
