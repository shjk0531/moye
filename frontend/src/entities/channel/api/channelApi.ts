import { apiClient } from '@/shared/api';
import type { StudyChannelResponse } from '../models/types';

/**
 * 채널 목록 조회
 * @param studyId 스터디 ID
 * @returns 채널 목록
 */
export async function fetchChannels(
    studyId: string,
): Promise<StudyChannelResponse> {
    const response = await apiClient.get<StudyChannelResponse>(
        `/api/v1/studies/${studyId}/channels`,
    );
    console.log(response.data);
    return response.data;
}
