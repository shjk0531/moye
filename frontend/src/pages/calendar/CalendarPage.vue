<template>
    <div class="notice bg-gray-850 px-4 py-2">
        <CalendarNotice
            @handle-date-change="handleDateChange"
            :selected-date="selectedDate"
        />
    </div>
    <div class="calendar-form">
        <CalendarForm :selected-date="selectedDate" />
    </div>
</template>

<script lang="ts">
import { CalendarForm } from '@/entities/calendar';
import { CalendarNotice } from '@/widgets/notice';
import { ref } from 'vue';

export default {
    name: 'CalendarPage',
    components: {
        CalendarForm,
        CalendarNotice,
    },
    setup() {
        const selectedDate = ref(new Date());

        const handleDateChange = (date: Date) => {
            selectedDate.value = new Date(date);
        };

        return {
            selectedDate,
            handleDateChange,
        };
    },
};
</script>

<style scoped>
.notice {
    grid-row: titlebarEnd / noticeEnd;
    grid-column: channelEnd / end;
}

.calendar-form {
    grid-row: noticeEnd / end;
    grid-column: channelEnd / pageEnd;
}
</style>
