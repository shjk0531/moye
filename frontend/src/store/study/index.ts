// src/store/study/index.ts
import { defineStore } from 'pinia';
import { state } from './state';
import { actions } from './actions';
import { getters } from './getters';

export const useStudyStore = defineStore('study', {
    state,
    actions,
    getters,
});
