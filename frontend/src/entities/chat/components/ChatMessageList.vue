<template>
    <div class="chat-message-list">
        <MessageItem
            v-for="message in messages"
            :key="message.id"
            :message="message"
        />
    </div>
</template>

<script setup lang="ts">
import { useChatStore } from '@/store/chat';
import MessageItem from './MessageItem.vue';
import { computed } from 'vue';

const chatStore = useChatStore();
const messages = computed(() => chatStore.messages);
</script>

<style scoped></style>
