import { apiClient } from '@/shared/api';
import type {
    User,
    LoginResponse,
    LoginRequest,
    SignupResponse,
    SignupRequest,
} from '@/entities/user';

export async function loginApi(
    loginRequest: LoginRequest,
): Promise<LoginResponse> {
    const response = await apiClient.post('/api/v1/auth/login', loginRequest);
    return response.data;
}

export async function signupApi(
    signupRequest: SignupRequest,
): Promise<SignupResponse> {
    const response = await apiClient.post(
        '/api/v1/auth/register',
        signupRequest,
    );
    return response.data;
}

export async function fetchUserProfile(): Promise<User> {
    const response = await apiClient.get('/api/v1/users/profile');
    return response.data;
}

export async function logoutApi() {
    await apiClient.post('/api/v1/auth/logout');
}
