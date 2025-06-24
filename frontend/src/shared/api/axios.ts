import axios from 'axios';
import { API_BASE_URL } from '../config';
import { useUserStore } from '@/store';
import { handleAuthError } from './authError';
import { refreshAccessToken } from './tokenService';

// API 클라이언트 인스턴스 생성
const apiClient = axios.create({
    baseURL: API_BASE_URL,
    headers: {
        'Content-Type': 'application/json',
    },
    withCredentials: true,
});

// 인터셉터 설정
apiClient.interceptors.request.use(
    (config) => {
        const userStore = useUserStore();
        const token = userStore.getToken;
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    },
);

// 응답 인터셉터 설정
apiClient.interceptors.response.use(
    (response) => response,
    async (error) => {
        const originalRequest: any = error.config;
        if (error.response) {
            const { status, data } = error.response;
            // 토큰 만료 시 재발급 시도
            if (
                status === 401 &&
                data.code === 'token_expired' &&
                !originalRequest._retry
            ) {
                originalRequest._retry = true;
                try {
                    const newToken = await refreshAccessToken();
                    originalRequest.headers[
                        'Authorization'
                    ] = `Bearer ${newToken}`;
                    return apiClient(originalRequest);
                } catch (refreshError) {
                    // 재발급 실패 시 인증 오류 처리
                    handleAuthError(data);
                    return Promise.reject(refreshError);
                }
            }
            // 그 외 인증 오류 처리
            if (status === 401 || status === 403) {
                handleAuthError(data);
            }
        }
        return Promise.reject(error);
    },
);

export default apiClient;
