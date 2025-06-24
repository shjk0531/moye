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
import type { ChatMessage } from '../models/types';
import MessageItem from './MessageItem.vue';

const props = defineProps<{
    messages: ChatMessage[];
}>();
</script>

<style scoped></style>
