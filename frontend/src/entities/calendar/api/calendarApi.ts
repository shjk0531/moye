import { apiClient } from '@/shared/api';
import type { StudyCalendarResponse } from '../model/types';

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

// 캘린더 이벤트 타입 정의
export interface CalendarEvent {
    id: string;
    title: string;
    content: string;
    startDate: Date;
    endDate: Date;
    color?: string;
}

export interface CalendarDay {
    date: Date;
    isCurrentMonth: boolean;
    isToday: boolean;
    isWeekend: boolean;
    isHoliday: boolean;
    events: CalendarEvent[];
}

/**
 * 캘린더 목록 조회
 * @param studyId 스터디 ID
 * @returns 캘린더 목록
 */
export async function fetchCalendars(
    studyId: string,
): Promise<StudyCalendarResponse> {
    const response = await apiClient.get<StudyCalendarResponse>(
        `/api/v1/studies/${studyId}/calendars`,
    );
    console.log(response.data);
    return response.data;
}
