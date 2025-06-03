// src/store/lounge/index.ts
import { defineStore } from 'pinia';
import { state } from './state';
import { actions } from './actions';
import { getters } from './getters';
import { persist } from './persist';

export const useLoungeStore = defineStore('lounge', {
    state,
    actions,
    getters,
    persist,
});
