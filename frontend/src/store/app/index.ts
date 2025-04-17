// src/store/app/index.ts
import { defineStore } from 'pinia';
import { state } from './state';
import { actions } from './actions';
import { getters } from './getters';

export const useAppStore = defineStore('app', {
    state,
    actions,
    getters,
});
