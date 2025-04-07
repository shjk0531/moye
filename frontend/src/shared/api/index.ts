import apiClient from './axios';

export { apiClient };

// 기본 API 요청 함수
export const api = {
    get: <T>(url: string, params?: any) => apiClient.get<T>(url, { params }),
    post: <T>(url: string, data?: any) => apiClient.post<T>(url, data),
    put: <T>(url: string, data?: any) => apiClient.put<T>(url, data),
    delete: <T>(url: string) => apiClient.delete<T>(url),
};
