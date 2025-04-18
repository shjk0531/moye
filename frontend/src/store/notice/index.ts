// src/store/notice/index.ts
import { defineStore } from 'pinia';
import { state } from './state';
import { actions } from './actions';
import { getters } from './getters';
import { persist } from './persist';

export const useNoticeStore = defineStore('notice', {
    state,
    actions,
    getters,
    persist,
});
