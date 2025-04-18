<template>
    <!-- 타이틀바 -->
    <!-- <TitleBar /> -->
    <div class="relative flex flex-row h-screen w-screen">
        <div class="left-sidebar flex flex-col dark:bg-gray-950">
            <div class="flex flex-row flex-grow-1">
                <StudyListSidebar />
                <router-view name="sidebar" />
            </div>
            <div class="flex flex-col">
                <UserPanel />
            </div>
        </div>
        <div class="page flex-grow-1">
            <router-view name="page" />
        </div>
    </div>
</template>

<script setup lang="ts">
import { StudyListSidebar, UserPanel } from '@/widgets';
</script>

<style scoped lang="scss"></style>
