import { apiClient } from '@/shared/api';
import type { Study } from '../models/types';

/**
 * 스터디 목록을 가져옵니다.
 */
export async function fetchStudies(): Promise<Study[]> {
    try {
        // 실제 환경에서는 아래 API 호출을 사용
        // const response = await apiClient.get('/studies');
        // return response.data;

        // 예시 데이터를 반환합니다.
        return [
            {
                id: 1,
                name: 'Study A',
                icon: 'https://picsum.photos/200/300?random=2',
            },
            {
                id: 2,
                name: 'Study B',
                icon: 'https://picsum.photos/200/300?random=3',
            },
            {
                id: 3,
                name: 'Study C',
                icon: 'https://picsum.photos/200/300?random=4',
            },
        ];
    } catch (error) {
        console.error('스터디 목록을 가져오는 중 오류 발생:', error);
        throw error;
    }
}

/**
 * 스터디 상세 정보를 가져옵니다.
 */
export async function fetchStudyById(studyId: number): Promise<Study> {
    try {
        // 실제 환경에서는 아래 API 호출을 사용
        // const response = await apiClient.get(`/studies/${studyId}`);
        // return response.data;

        // 예시 데이터를 반환합니다.
        return {
            id: studyId,
            name: `Study ${studyId}`,
            icon: `https://picsum.photos/200/300?random=${studyId + 1}`,
            description: `스터디 ${studyId}에 대한 설명입니다.`,
        };
    } catch (error) {
        console.error(
            `스터디 ID ${studyId}의 정보를 가져오는 중 오류 발생:`,
            error,
        );
        throw error;
    }
}
