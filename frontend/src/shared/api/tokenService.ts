import axios from 'axios';
import { API_BASE_URL } from '../config';
import { useUserStore } from '@/store/user';
import { handleAuthError } from './authError';

// 리프레시 동시 호출 방지를 위한 플래그 및 구독자 목록
let isRefreshing = false;
let refreshSubscribers: ((token: string) => void)[] = [];

// 토큰 갱신 후 대기 중인 요청을 재실행하도록 호출
function onRefreshed(token: string) {
    refreshSubscribers.forEach((cb) => cb(token));
    refreshSubscribers = [];
}

// 새로운 토큰을 받아오면 구독자에 등록
function subscribeTokenRefresh(cb: (token: string) => void) {
    refreshSubscribers.push(cb);
}

/**
 * 리프레시 토큰 엔드포인트를 호출하여 accessToken을 갱신하고 저장합니다.
 * 이미 갱신 중인 경우 앞선 갱신이 완료될 때까지 대기합니다.
 */
export async function refreshAccessToken(): Promise<string> {
    const userStore = useUserStore();
    const refreshClient = axios.create({
        baseURL: API_BASE_URL,
        withCredentials: true,
    });
    if (isRefreshing) {
        return new Promise((resolve) => subscribeTokenRefresh(resolve));
    }
    isRefreshing = true;

    try {
        // 쿠키 기반 리프레시 토큰 사용
        const response = await refreshClient.post('/auth/refresh');
        const newToken: string = response.data.access_token;
        // 사용자 스토어에 새로운 토큰 저장
        if (userStore.user) {
            userStore.login(userStore.user, newToken);
        }
        // 대기 중인 요청에 새로운 토큰 전달
        onRefreshed(newToken);
        return newToken;
    } catch (err: any) {
        // 리프레시 실패 시 인증 오류 처리
        handleAuthError(err.response?.data);
        throw err;
    } finally {
        isRefreshing = false;
    }
}
