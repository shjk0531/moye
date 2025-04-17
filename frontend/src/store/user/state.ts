import type { User } from '../../entities/user/models/model';

export interface UserState {
    user: User | null;
}

export const state = (): UserState => ({
    user: null,
});
