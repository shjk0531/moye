// src/store/chat/index.ts
import { defineStore } from 'pinia';
import { state } from './state';
import { actions } from './action';
import { getters } from './getter';
import { persist } from './persist';

export const useChatStore = defineStore('chat', {
    state,
    actions,
    getters,
    persist,
});
