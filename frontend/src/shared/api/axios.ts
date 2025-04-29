import axios from 'axios';
import { API_BASE_URL } from '../config';
import router from '@/router';
import { useUserStore } from '@/store';
import { handleAuthError } from './authError';

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
        const token = localStorage.getItem('access_token');
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
    (response) => {
        return response;
    },
    (error) => {
        if (error.response) {
            // 인증 관련 오류 (401: 권한 없음, 403: 금지됨)
            if (
                error.response.status === 401 ||
                error.response.status === 403
            ) {
                handleAuthError(error.response.data);
            }
        }
        return Promise.reject(error);
    },
);

export default apiClient;
