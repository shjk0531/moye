<!-- src/features/calender/components/CalendarList.vue -->

<template>
    <PanelList
        :groups="calendarsGrouped"
        :items="calendarsUngrouped"
        :activeItemId="activeCalendarId"
        listType="calendar"
        @item-click="handleCalendarClick"
    />
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useActiveItem } from '@/shared/composables';
import { PanelList } from '@/shared/panel';
import {
    fetchCalendarsGrouped,
    fetchCalendarsUngrouped,
    findFirstCalendar,
} from '@/entities/calendar';

const route = useRoute();
const router = useRouter();

const calendarsGrouped = ref([]);
const calendarsUngrouped = ref([]);

// 'calendar' 리스트 타입을 지정하여 composable을 사용합니다.
const {
    activeItemId: activeCalendarId,
    checkAndRedirect,
    currentStudyId,
} = useActiveItem('calendar');

onMounted(async () => {
    await loadCalendarData();
    await redirectToActiveCalendar();
});

// 캘린더 데이터 로드 함수
async function loadCalendarData() {
    try {
        calendarsGrouped.value = await fetchCalendarsGrouped();
        calendarsUngrouped.value = await fetchCalendarsUngrouped();
    } catch (error) {
        console.error('캘린더 데이터 로드 중 오류 발생:', error);
    }
}

// 활성 캘린더로 리다이렉트 함수
async function redirectToActiveCalendar() {
    try {
        await checkAndRedirect(`/study/${currentStudyId.value}/calendar`);
    } catch (error) {
        console.error('캘린더 리다이렉트 중 오류 발생:', error);
    }
}

// 캘린더 아이템 클릭 핸들러
function handleCalendarClick(item) {
    // activeCalendarId 업데이트 및 해당 캘린더 URL로 이동
    activeCalendarId.value = item.id;
    navigateToCalendar(item.id);
}

// 캘린더 페이지로 이동하는 함수
function navigateToCalendar(calendarId) {
    router.push(`/study/${currentStudyId.value}/calendar/${calendarId}`);
}
</script>
