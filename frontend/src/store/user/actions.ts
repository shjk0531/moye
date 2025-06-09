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
        this.accessToken = token;
    },

    logout(this: UserState) {
        this.user = null;
        this.accessToken = null;
    },
};
