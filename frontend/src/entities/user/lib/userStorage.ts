import type { User } from '../models/model';

const USER_KEY = 'user_data';
const TOKEN_KEY = 'access_token';

/**
 * 사용자 정보를 로컬 스토리지에 저장합니다.
 */
export const saveUser = (user: User): void => {
    localStorage.setItem(USER_KEY, JSON.stringify(user));
};

/**
 * 로컬 스토리지에서 사용자 정보를 가져옵니다.
 */
export const getUser = (): User | null => {
    const userData = localStorage.getItem(USER_KEY);
    return userData ? JSON.parse(userData) : null;
};

/**
 * 로컬 스토리지에서 사용자 정보를 삭제합니다.
 */
export const removeUser = (): void => {
    localStorage.removeItem(USER_KEY);
};

/**
 * 인증 토큰을 저장합니다.
 */
export const saveToken = (token: string): void => {
    localStorage.setItem(TOKEN_KEY, token);
};

/**
 * 인증 토큰을 가져옵니다.
 */
export const getToken = (): string | null => {
    return localStorage.getItem(TOKEN_KEY);
};

/**
 * 인증 토큰을 삭제합니다.
 */
export const removeToken = (): void => {
    localStorage.removeItem(TOKEN_KEY);
};

/**
 * 사용자가 로그인되어 있는지 확인합니다.
 */
export const isAuthenticated = (): boolean => {
    return !!getToken();
};
