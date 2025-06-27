import { apiClient } from '@/shared/api';
import type { MemberResponse } from '@/entities/member';

export async function fetchLoungeMembers(
    loungeId: string,
): Promise<MemberResponse> {
    const response = await apiClient.get(`/api/v1/lounges/${loungeId}/members`);
    console.log('members', response);
    return response.data;
}
