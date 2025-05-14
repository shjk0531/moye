import { apiClient } from '@/shared/api';
import type {
    Study,
    StudyCreatePayload,
    StudyIconResponse,
} from '../models/types';

export async function fetchStudyIcons(): Promise<StudyIconResponse> {
    try {
        const response = await apiClient.get('/api/v1/studies/user');
        return response.data;
    } catch (error) {
        console.error('스터디 아이콘 목록을 가져오는 중 오류 발생:', error);
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
