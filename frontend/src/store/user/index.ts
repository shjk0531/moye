// src/store/user/index.ts
import { defineStore } from 'pinia';
import { state } from './state';
import { actions } from './actions';
import { getters } from './getters';
import { persist } from './persist';

export const useUserStore = defineStore('user', {
    state,
    actions,
    getters,
    persist,
});
