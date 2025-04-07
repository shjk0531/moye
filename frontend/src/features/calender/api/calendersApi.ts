import { findFirstItem } from '@/shared/lib/panelUtils';
import type { PanelGroup, PanelItem } from '@/shared/lib/panelUtils';

// 캘린더 그룹 타입 정의
export interface CalendarGroup extends PanelGroup {
    id: number;
    label: string;
    order: number;
}

// 캘린더 타입 정의
export interface Calendar extends PanelItem {
    id: number;
    label: string;
    order: number;
    icon: string;
    groupId: number | null;
}

export async function fetchCalendarsGrouped(): Promise<CalendarGroup[]> {
    return [
        {
            id: 1,
            label: '달력 그룹1',
            order: 2,
        },
        {
            id: 2,
            label: '달력 그룹2',
            order: 3,
        },
    ];
}

export async function fetchCalendarsUngrouped(): Promise<Calendar[]> {
    return [
        { id: 0, label: '일정1', order: 1, icon: 'calendar', groupId: 1 },
        { id: 1, label: '일정2', order: 2, icon: 'calendar', groupId: 1 },
        { id: 2, label: '일정3', order: 1, icon: 'calendar', groupId: null },
        { id: 3, label: '일정4', order: 4, icon: 'calendar', groupId: null },
        { id: 4, label: '일정5', order: 5, icon: 'calendar', groupId: null },
        { id: 5, label: '일정6', order: 6, icon: 'calendar', groupId: null },
        { id: 6, label: '일정7', order: 1, icon: 'calendar', groupId: 2 },
        { id: 7, label: '일정8', order: 2, icon: 'calendar', groupId: 2 },
        { id: 8, label: '일정9', order: 3, icon: 'calendar', groupId: 2 },
        { id: 9, label: '일정10', order: 4, icon: 'calendar', groupId: 2 },
        { id: 10, label: '일정11', order: 5, icon: 'calendar', groupId: 2 },
    ];
}

/**
 * 그룹 및 캘린더 목록에서 첫 번째 캘린더를 찾는 함수
 * 내부적으로 공통 findFirstItem 함수를 사용합니다.
 *
 * @param groups 그룹 목록
 * @param calendars 캘린더 목록
 * @returns 첫 번째 캘린더 ID 또는 null
 */
export function findFirstCalendar(
    groups: CalendarGroup[],
    calendars: Calendar[],
): number | null {
    return findFirstItem(groups, calendars);
}
