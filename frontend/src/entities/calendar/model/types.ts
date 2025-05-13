export interface StudyCalendarResponse {
    items: StudyCalendarItem[];
}

export interface StudyCalendarItem {
    itemType: 'calendar' | 'group';
    calendar?: CalendarDTO;
    group?: CalendarGroupDTO;
}

export interface CalendarDTO {
    id: string;
    name: string;
}

export interface CalendarGroupDTO {
    id: string;
    name: string;
    calendars: CalendarDTO[];
}
