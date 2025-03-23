import { apiClient } from '../../../shared/api';
import type { User } from '../../../entities/user';

export interface LoginResponse {
    token: string;
    user: User;
}

export async function loginApi(
    email: string,
    password: string,
): Promise<LoginResponse> {
    const response = await apiClient.post('/auth/login', { email, password });
    return response.data;
}
