// src/entities/chat/api/useChatSocket.ts
import { onMounted, onBeforeUnmount, watch, toRef } from 'vue';
import { socketClient } from '@/shared';
import { useChatStore } from '@/store/chat';
import { useRoute } from 'vue-router';

export function useChatSocket() {
    const route = useRoute();
    const channelId = toRef(route.params, 'channelId');
    const chatStore = useChatStore();

    function setup(id: string) {
        const path = `/ws/channels/${id}`;
        socketClient.connect(path);
        socketClient.on('text', handleIncoming);
    }

    function cleanup() {
        socketClient.off('text', handleIncoming);
        socketClient.disconnect();
    }

    // 1) 들어오는 메시지를 처리할 핸들러
    function handleIncoming(raw: any) {
        chatStore.addMessage(raw);
    }

    // 2) 소켓 연결·이벤트 바인딩
    onMounted(() => {
        setup(channelId.value as string);
    });

    // 3) 컴포넌트 언마운트 시 정리
    onBeforeUnmount(cleanup);

    // 4) 채널 변경 시 재연결
    watch(channelId, (newId, oldId) => {
        if (newId === oldId) return;
        cleanup();
        chatStore.clearMessages();
        setup(newId as string);
    });

    // 5) 메시지 보내기 함수
    function sendMessage(content: string) {
        socketClient.send('text', { content });
    }

    return {
        messages: chatStore.messages,
        connected: chatStore.connected,
        sendMessage,
    };
}
