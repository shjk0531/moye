// src/store/user/state.ts
import type { User } from '@/entities/user';

export interface UserState {
    user: User | null;
    accessToken: string | null;
}

export const state = (): UserState => ({
    user: null,
    accessToken: null,
});
