export interface CalendarEvent {
    id: string;
    title: string;
    content: string;
    startDate: Date;
    endDate: Date;
    color: string;
}

export interface CalendarDay {
    date: Date;
    isCurrentMonth: boolean;
    isToday: boolean;
    isWeekend: boolean;
    isHoliday: boolean;
    events: CalendarEvent[];
}
