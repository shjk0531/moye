<!-- MessageItem.vue -->
<template>
    <div class="message-item">
        <img :src="message.avatar" class="avatar" alt="avatar" />
        <div class="message-content">
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

<style scoped>
.message-item {
    display: flex;
    margin-bottom: 15px;
}

.avatar {
    width: 40px;
    height: 40px;
    border-radius: 20px;
    margin-right: 10px;
}

.message-content {
    max-width: 80%;
}
</style>
