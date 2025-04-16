import { apiClient } from '../../../shared/api';
import type { User } from '../../../entities/user';

export interface TokenResponse {
    access_token: string;
    token_type: string;
    expires_in: number;
}

export async function loginApi(
    email: string,
    password: string,
): Promise<TokenResponse> {
    const response = await apiClient.post('/auth/login', { email, password });
    return response.data;
}

export async function signupApi(
    email: string,
    password: string,
    nickname: string,
    profile: string,
): Promise<{ message: string }> {
    const response = await apiClient.post('/auth/register', {
        email,
        password,
        nickname,
        profile,
    });
    return response.data;
}
