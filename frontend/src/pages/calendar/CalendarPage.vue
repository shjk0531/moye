<template>
    <PageLayout>
        <template #notice>
            <CalendarNotice
                @handleDateChange="handleDateChange"
                :selectedDate="selectedDate"
            />
        </template>
        <template #content>
            <!-- 캘린더 내용 -->
            <div class="h-full">
                <CalendarForm :selectedDate="selectedDate" />
            </div>
        </template>
        <template #sidebar>
            <MemberList />
        </template>
    </PageLayout>
</template>

<script setup lang="ts">
import { MemberList } from '@/widgets/sidebar';
import { CalendarNotice } from '@/widgets/notice';
import { PageLayout } from '@/shared/layout';
import { CalendarForm } from '@/entities';
import { ref } from 'vue';

const selectedDate = ref(new Date());
const handleDateChange = (date: Date) => {
    selectedDate.value = date;
};
</script>

<style scoped></style>
