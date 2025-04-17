// src/store/user/actions.ts
import type { User } from '@/entities/user';
import type { UserState } from './state';

export const actions = {
    setUser(this: UserState, user: User) {
        this.user = user;
    },

    clearUser(this: UserState) {
        this.user = null;
    },

    login(this: UserState, user: User, token: string) {
        this.user = user;
        localStorage.setItem('access_token', token);
    },

    logout(this: UserState) {
        this.user = null;
        localStorage.removeItem('access_token');
    },

    checkAuth(this: UserState) {
        const token = localStorage.getItem('access_token');
        if (!token && this.user) {
            this.user = null;
        }
        return !!token;
    },

    getToken(this: UserState) {
        if (!this.user) {
            return null;
        }
        return localStorage.getItem('access_token');
    },
};
