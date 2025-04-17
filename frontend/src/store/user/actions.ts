import type { User } from '../../entities/user/models/model';
import type { UserState } from './state';

export const actions = {
    setUser(this: UserState, user: User) {
        this.user = user;
    },

    logout(this: UserState) {
        this.user = null;
        localStorage.removeItem('token');
    },
};
