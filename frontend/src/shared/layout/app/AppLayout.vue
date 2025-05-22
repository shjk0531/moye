<template>
    <div class="flex h-screen w-screen">
        <!-- 사이드바 -->
        <div class="left-sidebar flex flex-col h-full dark:bg-gray-950">
            <!-- 상단: 스터디 리스트 + 서브 사이드바 -->
            <div class="flex flex-1 overflow-hidden min-h-0">
                <StudyListSidebar />
                <router-view name="sidebar" />
            </div>
            <!-- 하단: 유저 패널 -->
            <UserPanel />
        </div>
        <!-- 메인 페이지 영역 -->
        <div class="page flex-1 overflow-auto">
            <router-view name="page" />
        </div>
    </div>
</template>

<script setup lang="ts">
import { StudyListSidebar, UserPanel } from '@/widgets';
</script>
