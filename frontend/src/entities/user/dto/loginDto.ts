import type { User } from '@/entities/user';

export interface LoginResponse {
    access_token: string;
    user: User;
}

export interface LoginRequest {
    email: string;
    password: string;
}
