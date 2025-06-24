<template>
    <PageLayout>
        <template #notice>
            <ChatNotice />
        </template>
        <template #content>
            <div
                class="message-area flex flex-col flex-1 overflow-y-auto overflow-x-hidden"
            >
                <ChatMessageList :messages="messages" />
            </div>
            <div class="chat-area pb-4 px-2">
                <MessageInput @sendMessage="sendMessage" />
            </div>
        </template>
        <template #sidebar>
            <MemberList />
        </template>
    </PageLayout>
</template>

<script setup lang="ts">
import { MessageInput } from '@/features/chat';
import { ChatNotice } from '@/widgets/notice';
import { MemberList } from '@/widgets/sidebar';
import { PageLayout } from '@/shared/layout';
import { useChatSocket, ChatMessageList } from '@/entities/chat';
import { computed, onMounted, watch } from 'vue';
import { useChatStore } from '@/store/chat';
import { useRoute } from 'vue-router';

const route = useRoute();
const chatStore = useChatStore();
const channelId = () => route.params.channelId as string;

onMounted(() => {
    if (channelId()) {
        chatStore.selectChannel(channelId());
    }
});

watch(
    () => route.params.channelId,
    (newId, oldId) => {
        if (newId && newId !== oldId) {
            chatStore.selectChannel(newId as string);
        }
    },
);

function sendMessage(content: string) {
    chatStore.sendMessage(content);
}

const messages = computed(() => chatStore.messages);
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
