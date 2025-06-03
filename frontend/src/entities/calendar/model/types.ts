export interface LoungeCalendarResponse {
    items: LoungeCalendarItem[];
}

export interface LoungeCalendarItem {
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
