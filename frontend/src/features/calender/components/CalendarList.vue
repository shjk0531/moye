<!-- src/features/calendar/components/CalendarList.vue -->

<template>
    <PanelList
        :groups="calendarsGrouped"
        :items="calendarsUngrouped"
        :activeItemId="activeCalendarId"
        @item-click="handleCalendarClick"
    />
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import PanelList from '@/components/panel/PanelList.vue';
import {
    fetchCalendarsGrouped,
    fetchCalendarsUngrouped,
} from '@/features/calendar/api/calendarApi'; // 달력 API

const router = useRouter();

const calendarsGrouped = ref([]);
const calendarsUngrouped = ref([]);

onMounted(async () => {
    calendarsGrouped.value = await fetchCalendarsGrouped();
    calendarsUngrouped.value = await fetchCalendarsUngrouped();
});

// activeCalendarId 관리는 필요에 따라 구현 (예: Vuex 또는 local state)
const activeCalendarId = ref(null);

function handleCalendarClick(item) {
    activeCalendarId.value = item.id;
    router.push(`/calendar/${item.id}`); // 라우터 경로는 요구사항에 맞게 수정
}
</script>
