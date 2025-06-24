// src/store/chat/getter.ts

import type { ChatState } from './state';

export const getters = {
    isConnected: (state: ChatState) => state.connected,
    getMessages: (state: ChatState) => state.messages,
};
