import { defineStore } from 'pinia';
import type { User } from '../models/model';

export const useUserStore = defineStore('user', {
    state: () => ({
        user: null as User | null,
    }),
    actions: {
        setUser(user: User) {
            this.user = user;
        },
        logout() {
            this.user = null;
            localStorage.removeItem('token');
        },
    },
});
