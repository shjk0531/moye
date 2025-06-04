import { apiClient } from '@/shared/api';
import type { LoungeChannelResponse } from '../models/types';

/**
 * 채널 목록 조회
 * @param loungeId 라운지 ID
 * @returns 채널 목록
 */
export async function fetchChannels(
    loungeId: string,
): Promise<LoungeChannelResponse> {
    try {
        const response = await apiClient.get<LoungeChannelResponse>(
            `/api/v1/lounges/${loungeId}/channels`,
        );
        console.log('채널 목록 조회 성공:', response.data);
        return response.data;
    } catch (error) {
        console.error('채널 목록 조회 중 오류 발생:', error);
        throw error;
    }
}
