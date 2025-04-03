<!-- MessageItem.vue -->
<template>
    <div class="flex mb-3 pl-4">
        <img
            :src="message.avatar"
            class="w-10 h-10 rounded-full"
            alt="avatar"
        />
        <div class="message-content max-w-[80%]">
            <div class="message-header">
                <span class="username text-gray-50 text-md">
                    {{ message.username }}
                </span>
                <span class="timestamp text-gray-300 text-sm">
                    {{ message.timestamp }}
                </span>
            </div>
            <div class="flex message-body text-gray-50 text-md justify-start">
                <p>{{ message.text }}</p>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'MessageItem',
    props: {
        message: {
            type: Object,
            required: true,
        },
    },
};
</script>

<style scoped></style>
