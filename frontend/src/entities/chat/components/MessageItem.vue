<!-- MessageItem.vue -->
<template>
    <div class="flex mb-5 pl-4">
        <img
            :src="message.profileImage"
            class="w-10 h-10 rounded-full mr-2"
            alt="avatar"
        />
        <div class="message-content max-w-[90%]">
            <div class="message-header">
                <span class="username text-gray-50 text-md">
                    {{ message.senderName }}
                </span>
                <span class="timestamp text-gray-300 text-sm">
                    {{ message.createdAt }}
                </span>
            </div>
            <div class="flex message-body text-gray-50 text-md justify-start">
                <p>{{ message.content }}</p>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import type { ChatMessage } from '@/entities/chat';

defineProps<{
    message: ChatMessage;
}>();
</script>

<style scoped></style>
