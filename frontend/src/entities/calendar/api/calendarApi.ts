import type { CalendarEvent } from '../model/types';

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
