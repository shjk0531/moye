<!-- src/widgets/notice/components/CalendarNotice.vue -->

<template>
    <div class="flex flex-row w-full justify-between items-center">
        <div class="flex flex-row flex-1 items-center gap-2">
            <!-- 캘린더 타이틀 -->
            <div class="flex flex-row items-center gap-2">
                <span
                    class="mdi mdi-chevron-left text-gray-150 hover:text-gray-200 cursor-pointer"
                    @click="handlePrevMonth"
                ></span>
                <span class="text-gray-50">
                    {{ selectedDate.getFullYear() }}년
                    {{ selectedDate.getMonth() + 1 }}월
                </span>
                <span
                    class="mdi mdi-chevron-right text-gray-150 hover:text-gray-200 cursor-pointer"
                    @click="handleNextMonth"
                ></span>
            </div>
            <!-- 캘린더 버튼 -->
        </div>
        <div class="flex flex-row gap-4 text-2xl">
            <ChannelButton />
            <CalendarButton />
            <PostButton />
            <NoteButton />
            <MemberButton />
        </div>
    </div>
</template>

<script setup lang="ts">
import {
    ChannelButton,
    CalendarButton,
    PostButton,
    NoteButton,
    MemberButton,
} from '@/shared/ui/Button';

const props = defineProps<{
    selectedDate: Date;
}>();

const emit = defineEmits<{
    (e: 'handleDateChange', date: Date): void;
}>();

const handlePrevMonth = () => {
    const newDate = new Date(props.selectedDate);
    newDate.setMonth(newDate.getMonth() - 1);
    emit('handleDateChange', newDate);
};

const handleNextMonth = () => {
    const newDate = new Date(props.selectedDate);
    newDate.setMonth(newDate.getMonth() + 1);
    emit('handleDateChange', newDate);
};
</script>

<style scoped></style>
