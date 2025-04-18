<template>
    <PageLayout>
        <template #notice>
            <ChatNotice />
        </template>
        <template #content>
            <div
                class="message-area flex flex-col flex-1 overflow-y-auto overflow-x-hidden"
            >
                <ScrollPanel style="height: 100%">
                    <MessageItem
                        v-for="message in messages"
                        :key="message.id"
                        :message="message"
                    />
                </ScrollPanel>
            </div>
            <div class="chat-area pb-4 px-2">
                <MessageInput @sendMessage="handleSendMessage" />
            </div>
        </template>
        <template #sidebar>
            <MemberList />
        </template>
    </PageLayout>
</template>

<script setup lang="ts">
import { MessageInput, MessageItem } from '@/features/chat';
import { ChatNotice } from '@/widgets/notice';
import { MemberList } from '@/widgets/sidebar';
import { PageLayout } from '@/shared/layout';
import { ref } from 'vue';

interface Message {
    id: number;
    avatar: string;
    username: string;
    timestamp: string;
    text: string;
}

const messages = ref<Message[]>([
    {
        id: 1,
        avatar: 'https://picsum.photos/200/300?random=1',
        username: 'User1',
        timestamp: '10:00 AM',
        text: 'Hello, world!',
    },
    {
        id: 2,
        avatar: 'https://picsum.photos/200/300?random=2',
        username: 'User2',
        timestamp: '10:05 AM',
        text: 'Hi there!',
    },
    {
        id: 3,
        avatar: 'https://picsum.photos/200/300?random=3',
        username: 'User3',
        timestamp: '10:10 AM',
        text: 'How are you?',
    },
    {
        id: 4,
        avatar: 'https://picsum.photos/200/300?random=4',
        username: 'User4',
        timestamp: '10:15 AM',
        text: 'I am fine, thanks!',
    },
    {
        id: 5,
        avatar: 'https://picsum.photos/200/300?random=5',
        username: 'User5',
        timestamp: '10:20 AM',
        text: 'What about you?',
    },
    {
        id: 6,
        avatar: 'https://picsum.photos/200/300?random=6',
        username: 'User6',
        timestamp: '10:25 AM',
        text: 'I am doing great!',
    },
]);

const handleSendMessage = (messageText: string) => {
    messages.value.push({
        id: messages.value.length + 1,
        avatar:
            'https://picsum.photos/200/300?random=0' +
            (messages.value.length + 1),
        username: 'You',
        timestamp: new Date().toLocaleTimeString(),
        text: messageText,
    });
};
</script>

<style scoped>
.message-area :deep(.p-scrollpanel-content) {
    height: 100%;
}

.message-area :deep(.p-scrollpanel-bar) {
    background-color: var(--custom-scrollbar-color);
    border-radius: 4px;
    transition: background-color 0.3s ease;
}

.message-area :deep(.p-scrollpanel-bar:hover) {
    background-color: var(--custom-scrollbar-color-hover);
}

.message-area :deep(.p-scrollpanel-bar:active) {
    background-color: var(--custom-scrollbar-color-active);
}

.message-area :deep(.p-scrollpanel-bar-y) {
    width: calc(var(--spacing) * 2) !important;
}
</style>
