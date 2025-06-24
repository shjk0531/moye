// src/store/chat/state.ts
import type { ChatMessage } from '@/entities/chat';

export interface ChatState {
    channelId: string;
    connected: boolean;
    messages: ChatMessage[];
}

export const state = (): ChatState => ({
    channelId: '',
    connected: false,
    messages: [] as ChatMessage[],
});
