<template>
    <div class="calendar-notice text-white p-2">
        <p>공지사항</p>
    </div>

    <CalendarList />
</template>

<script lang="ts">
import { CalendarList } from '@/features/calendar';

export default {
    name: 'CalendarListSidebar',
    components: {
        CalendarList,
    },
};
</script>

<style scoped lang="scss"></style>
