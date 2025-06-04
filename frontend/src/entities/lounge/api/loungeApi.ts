import { apiClient } from '@/shared/api';
import type {
    Lounge,
    LoungeCreatePayload,
    LoungeIconResponse,
    LoungeListResponse,
} from '../models/types';

export async function fetchLoungeIcons(): Promise<LoungeIconResponse> {
    try {
        const response = await apiClient.get('/api/v1/lounges/user');
        return response.data;
    } catch (error) {
        console.error('라운지 아이콘 목록을 가져오는 중 오류 발생:', error);
        throw error;
    }
}

/**
 * 라운지 생성
 */
export async function createLounge(
    payload: LoungeCreatePayload,
): Promise<Lounge> {
    try {
        const response = await apiClient.post('/api/v1/lounges', payload);
        return response.data;
    } catch (error) {
        console.error('라운지 생성 중 오류 발생:', error);
        throw error;
    }
}

/**
 * 라운지 목록 조회
 */
export async function fetchLoungeList(): Promise<LoungeListResponse> {
    try {
        const response = await apiClient.get('/api/v1/lounges');
        console.log(response.data);
        return response.data;
    } catch (error) {
        console.error('라운지 목록 조회 중 오류 발생:', error);
        throw error;
    }
}
