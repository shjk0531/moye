import { ref, computed } from 'vue';
import { fetchCalendarEvents } from '@/entities/calendar';
import type { CalendarEvent, CalendarDay } from '@/entities/calendar';

export function useCalendar() {
    const currentDate = ref(new Date());
    const events = ref<CalendarEvent[]>([]);
    const loading = ref(false);
    const error = ref<string | null>(null);

    // 현재 연도
    const currentYear = computed(() => currentDate.value.getFullYear());

    // 현재 월 (0-11)
    const currentMonth = computed(() => currentDate.value.getMonth());

    // 이전 월로 이동
    const prevMonth = () => {
        const date = new Date(currentDate.value);
        date.setMonth(date.getMonth() - 1);
        currentDate.value = date;
        loadEvents();
    };

    // 다음 월로 이동
    const nextMonth = () => {
        const date = new Date(currentDate.value);
        date.setMonth(date.getMonth() + 1);
        currentDate.value = date;
        loadEvents();
    };

    // 오늘 날짜로 이동
    const goToToday = () => {
        currentDate.value = new Date();
        loadEvents();
    };

    // 특정 월의 날짜 배열 생성 (달력용)
    const getDaysInMonth = computed(() => {
        const year = currentYear.value;
        const month = currentMonth.value;
        const daysInMonth = new Date(year, month + 1, 0).getDate();
        const firstDayOfMonth = new Date(year, month, 1).getDay();

        const days: CalendarDay[] = [];

        // 이전 달의 마지막 날짜들 (달력 첫 주 채우기)
        const prevMonthLastDay = new Date(year, month, 0).getDate();
        for (let i = firstDayOfMonth - 1; i >= 0; i--) {
            const date = new Date(year, month - 1, prevMonthLastDay - i);
            days.push({
                date,
                isCurrentMonth: false,
                isToday: isSameDay(date, new Date()),
                isWeekend: isWeekend(date),
                isHoliday: false, // 휴일 로직 추가 필요
                events: getEventsForDay(date),
            });
        }

        // 현재 달 날짜들
        for (let i = 1; i <= daysInMonth; i++) {
            const date = new Date(year, month, i);
            days.push({
                date,
                isCurrentMonth: true,
                isToday: isSameDay(date, new Date()),
                isWeekend: isWeekend(date),
                isHoliday: false, // 휴일 로직 추가 필요
                events: getEventsForDay(date),
            });
        }

        // 다음 달 날짜들 (달력 마지막 주 채우기)
        const remainingDays = 42 - days.length; // 6주 달력을 채우기 위해
        for (let i = 1; i <= remainingDays; i++) {
            const date = new Date(year, month + 1, i);
            days.push({
                date,
                isCurrentMonth: false,
                isToday: isSameDay(date, new Date()),
                isWeekend: isWeekend(date),
                isHoliday: false, // 휴일 로직 추가 필요
                events: getEventsForDay(date),
            });
        }

        return days;
    });

    // 특정 날짜에 해당하는 이벤트 필터링
    const getEventsForDay = (date: Date) => {
        return events.value.filter((event) => {
            const eventStart = new Date(event.startDate);
            const eventEnd = new Date(event.endDate);
            return (
                date >= new Date(eventStart.setHours(0, 0, 0, 0)) &&
                date <= new Date(eventEnd.setHours(23, 59, 59, 999))
            );
        });
    };

    // 날짜가 같은지 비교
    const isSameDay = (date1: Date, date2: Date) => {
        return (
            date1.getDate() === date2.getDate() &&
            date1.getMonth() === date2.getMonth() &&
            date1.getFullYear() === date2.getFullYear()
        );
    };

    // 주말인지 확인
    const isWeekend = (date: Date) => {
        const day = date.getDay();
        return day === 0 || day === 6;
    };

    // 이벤트 데이터 로드
    const loadEvents = async () => {
        try {
            loading.value = true;
            error.value = null;
            events.value = await fetchCalendarEvents(
                currentYear.value,
                currentMonth.value,
            );
        } catch (e) {
            console.error('캘린더 이벤트를 불러오는 중 오류 발생:', e);
            error.value = '이벤트를 불러올 수 없습니다.';
            events.value = [];
        } finally {
            loading.value = false;
        }
    };

    return {
        currentDate,
        events,
        loading,
        error,
        currentYear,
        currentMonth,
        prevMonth,
        nextMonth,
        goToToday,
        getDaysInMonth,
        loadEvents,
    };
}
