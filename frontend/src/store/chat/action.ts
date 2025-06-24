// src/store/chat/action.ts
import type { ChatMessage } from '@/entities/chat';
import type { ChatState } from './state';
import { socketClient } from '@/shared/api/socket';

type ChatStore = ChatState & {
    setConnected: (flag: boolean) => void;
    addMessage: (msg: ChatMessage) => void;
    clearMessages: () => void;
    selectChannel: (channelId: string) => Promise<void>;
    sendMessage: (content: string) => void;
};

export const actions = {
    setConnected(this: ChatStore, flag: boolean) {
        this.connected = flag;
    },
    addMessage(this: ChatStore, message: ChatMessage) {
        this.messages.push(message);
    },
    clearMessages(this: ChatStore) {
        this.messages = [];
    },
    async selectChannel(this: ChatStore, channelId: string) {
        if (this.connected) {
            socketClient.off('text', this.addMessage);
            socketClient.disconnect();
        }
        this.channelId = channelId;
        this.messages = [];

        await socketClient.connect(`/ws/channels/${channelId}`);
        socketClient.on('text', this.addMessage);
        this.connected = true;
    },
    sendMessage(this: ChatStore, content: string) {
        if (!this.connected) {
            console.error('Chat socket not connected');
            return;
        }
        socketClient.send('text', { channelId: this.channelId, content });
    },
};
