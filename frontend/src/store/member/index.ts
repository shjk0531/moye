// src/store/member/index.ts
import { defineStore } from 'pinia';
import { state } from './state';
import { actions } from './actions';
import { getters } from './getters';
import { persist } from './persist';

export const useMemberStore = defineStore('member', {
    state,
    actions,
    getters,
    persist,
});
