import { apiClient } from '@/shared/api';
import type { Study, StudyCreatePayload } from '../models/types';

/**
 * 스터디 목록을 가져옵니다.
 */
export async function fetchStudies(): Promise<Study[]> {
    try {
        // 실제 환경에서는 아래 API 호출을 사용
        // const response = await apiClient.get('/api/studies');
        // return response.data;

        // 예시 데이터를 반환합니다.
        return [
            {
                id: '1',
                name: 'Study A',
                profileUrl: 'https://picsum.photos/200/300?random=2',
                description: '스터디 A에 대한 설명입니다.',
            },
            {
                id: '2',
                name: 'Study B',
                profileUrl: 'https://picsum.photos/200/300?random=3',
                description: '스터디 B에 대한 설명입니다.',
            },
            {
                id: '3',
                name: 'Study C',
                profileUrl: 'https://picsum.photos/200/300?random=4',
                description: '스터디 C에 대한 설명입니다.',
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
export async function fetchStudyById(studyId: string): Promise<Study> {
    try {
        // 실제 환경에서는 아래 API 호출을 사용
        // const response = await apiClient.get(`/api/studies/${studyId}`);
        // return response.data;

        // 예시 데이터를 반환합니다.
        return {
            id: studyId,
            name: `Study ${studyId}`,
            profileUrl: `https://picsum.photos/200/300?random=${studyId}`,
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

/**
 * 스터디를 생성합니다.
 */
export async function createStudy(payload: StudyCreatePayload): Promise<Study> {
    try {
        const response = await apiClient.post('/studies', payload);
        return response.data;
    } catch (error) {
        console.error('스터디 생성 중 오류 발생:', error);
        throw error;
    }
}
