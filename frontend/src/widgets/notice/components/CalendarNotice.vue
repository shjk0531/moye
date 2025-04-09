<!-- src/widgets/notice/components/CalendarNotice.vue -->

<template>
    <div class="flex flex-row justify-between items-center">
        <div class="flex flex-row items-center gap-2">
            <!-- 캘린더 타이틀 -->
            <div class="flex flex-row items-center gap-2">
                <!-- 이전 달 버튼 -->
                <span
                    class="mdi mdi-chevron-left text-gray-150 hover:text-gray-200 cursor-pointer"
                    @click="handlePrevMonth"
                ></span>
                <!-- 선택 날짜(default: 오늘)(0000년 00월) -->
                <span class="text-gray-50">
                    {{ selectedDate.getFullYear() }}년
                    {{ selectedDate.getMonth() + 1 }}월
                </span>
                <!-- 다음 달 버튼 -->
                <span
                    class="mdi mdi-chevron-right text-gray-150 hover:text-gray-200 cursor-pointer"
                    @click="handleNextMonth"
                ></span>
            </div>
            <!-- 캘린더 월 선택 아이콘-->
            <span
                class="mdi mdi-calendar-month text-gray-150 hover:text-gray-200 cursor-pointer"
                @click="handleDateSelect"
            ></span>
        </div>
        <div class="notice-icons flex flex-row gap-4 text-2xl">
            <ChannelButton />
            <CalendarButton />
            <PostButton />
            <NoteButton />
            <MemberButton />
        </div>
    </div>
</template>

<script>
import {
    ChannelButton,
    CalendarButton,
    PostButton,
    NoteButton,
    MemberButton,
} from '@/shared/ui/Button';
import { ref } from 'vue';

export default {
    name: 'CalendarNotice',
    components: {
        ChannelButton,
        CalendarButton,
        PostButton,
        NoteButton,
        MemberButton,
    },
    props: {
        selectedDate: {
            type: Date,
            required: true,
        },
    },
    methods: {
        handlePrevMonth() {
            const newDate = new Date(this.selectedDate);
            newDate.setMonth(newDate.getMonth() - 1);
            this.$emit('handle-date-change', newDate);
        },
        handleNextMonth() {
            const newDate = new Date(this.selectedDate);
            newDate.setMonth(newDate.getMonth() + 1);
            this.$emit('handle-date-change', newDate);
        },
    },
};
</script>
