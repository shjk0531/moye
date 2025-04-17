<template>
    <div class="calendar-view">
        <div class="calendar-header">
            <div class="flex items-center justify-between mb-4">
                <div class="flex items-center space-x-2">
                    <Button
                        icon="pi pi-arrow-left"
                        outlined
                        size="small"
                        @click="prevMonth"
                    />
                    <Button
                        icon="pi pi-arrow-right"
                        outlined
                        size="small"
                        @click="nextMonth"
                    />
                    <Button
                        label="오늘"
                        outlined
                        size="small"
                        class="ml-2"
                        @click="goToToday"
                    />
                </div>
                <div class="text-xl font-semibold">
                    {{ currentYear }}년 {{ currentMonth + 1 }}월
                </div>
            </div>
        </div>
        <div class="calendar-grid">
            <!-- 요일 헤더 -->
            <div
                v-for="day in weekdays"
                :key="day"
                class="calendar-day-header"
                :class="{ weekend: day === '일' || day === '토' }"
            >
                {{ day }}
            </div>

            <!-- 날짜 그리드 -->
            <div
                v-for="(day, index) in days"
                :key="index"
                class="calendar-day"
                :class="{
                    'not-current-month': !day.isCurrentMonth,
                    today: day.isToday,
                    weekend: day.isWeekend,
                    holiday: day.isHoliday,
                }"
            >
                <div class="day-number">{{ day.date.getDate() }}</div>
                <div class="day-events" v-if="day.events.length > 0">
                    <div
                        v-for="event in day.events.slice(0, 3)"
                        :key="event.id"
                        class="day-event"
                        :style="{ backgroundColor: event.color }"
                    >
                        {{ event.title }}
                    </div>
                    <div class="more-events" v-if="day.events.length > 3">
                        +{{ day.events.length - 3 }} 더보기
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useCalendar } from '../composables/useCalendar';
import Button from 'primevue/button';

// 요일 배열
const weekdays = ['일', '월', '화', '수', '목', '금', '토'];

// useCalendar 컴포저블 사용
const {
    currentDate,
    currentYear,
    currentMonth,
    getDaysInMonth,
    prevMonth,
    nextMonth,
    goToToday,
    loadEvents,
} = useCalendar();

const days = computed(() => getDaysInMonth.value);

// 컴포넌트 마운트 시 이벤트 로드
onMounted(() => {
    loadEvents();
});
</script>

<style scoped>
.calendar-view {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    background-color: white;
    border-radius: 8px;
    padding: 16px;
}

.calendar-grid {
    display: grid;
    grid-template-columns: repeat(7, 1fr);
    gap: 1px;
    border: 1px solid #e5e7eb;
    border-radius: 8px;
    overflow: hidden;
    flex-grow: 1;
}

.calendar-day-header {
    padding: 8px;
    text-align: center;
    font-weight: 600;
    background-color: #f9fafb;
    border-bottom: 1px solid #e5e7eb;
}

.calendar-day {
    position: relative;
    min-height: 100px;
    padding: 4px;
    border-right: 1px solid #e5e7eb;
    border-bottom: 1px solid #e5e7eb;
    background-color: white;
}

.calendar-day:nth-child(7n) {
    border-right: none;
}

.day-number {
    font-weight: 500;
    margin-bottom: 4px;
}

.not-current-month {
    background-color: #f9fafb;
    color: #9ca3af;
}

.today {
    background-color: rgba(59, 130, 246, 0.05);
}

.today .day-number {
    color: #3b82f6;
    font-weight: 700;
}

.weekend {
    color: #ef4444;
}

.holiday {
    color: #ef4444;
}

.day-events {
    display: flex;
    flex-direction: column;
    gap: 2px;
}

.day-event {
    padding: 2px 4px;
    border-radius: 2px;
    font-size: 0.8rem;
    color: white;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
}

.more-events {
    font-size: 0.75rem;
    color: #6b7280;
    text-align: right;
    margin-top: 2px;
}
</style>
