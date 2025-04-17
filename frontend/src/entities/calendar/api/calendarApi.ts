import type { CalendarEvent } from '../model/types';
import { findFirstItem } from '@/shared/lib/panelUtils';

// 캘린더 그룹 타입 정의
export interface CalendarGroup {
    id: number;
    label: string;
    order: number;
}

// 캘린더 타입 정의
export interface Calendar {
    id: number;
    label: string;
    order: number;
    icon: string;
    groupId: number | null;
}

// 실제 API 호출로 대체해야 합니다
export const fetchCalendarEvents = async (
    year: number,
    month: number,
): Promise<CalendarEvent[]> => {
    // 임시 데이터
    return [
        {
            id: '1',
            title: '프로젝트 시작',
            content: '새로운 프로젝트 시작일',
            startDate: new Date(2024, 3, 15),
            endDate: new Date(2024, 3, 20),
            color: '#FF5733',
        },
        {
            id: '2',
            title: '회의',
            content: '주간 회의',
            startDate: new Date(2024, 3, 10),
            endDate: new Date(2024, 3, 10),
            color: '#33FF57',
        },
    ];
};

/**
 * 캘린더 그룹 목록을 가져옵니다.
 */
export async function fetchCalendarsGrouped(): Promise<CalendarGroup[]> {
    try {
        // 예시 데이터를 반환합니다.
        return [
            {
                id: 1,
                label: '일정 그룹1',
                order: 1,
            },
            {
                id: 2,
                label: '일정 그룹2',
                order: 2,
            },
        ];
    } catch (error) {
        console.error('캘린더 그룹을 가져오는 중 오류 발생:', error);
        throw error;
    }
}

/**
 * 캘린더 목록을 가져옵니다.
 */
export async function fetchCalendarsUngrouped(): Promise<Calendar[]> {
    try {
        // 예시 데이터를 반환합니다.
        return [
            { id: 0, label: '일정1', order: 1, icon: 'calendar', groupId: 1 },
            {
                id: 1,
                label: '일정2',
                order: 2,
                icon: 'calendar-days',
                groupId: 1,
            },
            {
                id: 2,
                label: '일정3',
                order: 1,
                icon: 'calendar-check',
                groupId: null,
            },
            {
                id: 3,
                label: '일정4',
                order: 2,
                icon: 'calendar-plus',
                groupId: null,
            },
            {
                id: 4,
                label: '일정5',
                order: 3,
                icon: 'calendar-week',
                groupId: 2,
            },
            {
                id: 5,
                label: '일정6',
                order: 4,
                icon: 'calendar-minus',
                groupId: 2,
            },
        ];
    } catch (error) {
        console.error('캘린더 목록을 가져오는 중 오류 발생:', error);
        throw error;
    }
}

/**
 * 그룹 및 캘린더 목록에서 첫 번째 캘린더를 찾는 함수
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
