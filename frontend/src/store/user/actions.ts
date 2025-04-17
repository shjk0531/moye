import type { User } from '@/entities/user';
import type { UserState } from './state';

export const actions = {
    setUser(this: UserState, user: User) {
        this.user = user;
    },

    clearUser(this: UserState) {
        this.user = null;
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
};
