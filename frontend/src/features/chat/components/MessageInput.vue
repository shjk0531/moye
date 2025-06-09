<template>
    <div
        class="flex items-center p-2 bg-gray-700 text-gray-50 h-(--custom-bottom-pannel-height) rounded-lg p-2"
    >
        <input
            type="text"
            v-model="message"
            class="!bg-gray-700 !border-none !focus:ring-0 !focus:shadow-none !focus:border-none !flex-1"
            @keyup.enter="send"
        />
        <button
            class="bg-gray-600 text-gray-50 px-4 py-2 rounded-lg"
            @click="send"
        >
            Send
        </button>
    </div>
</template>

<script>
export default {
    name: 'MessageInput',
    data() {
        return {
            message: '',
        };
    },
    methods: {
        send() {
            if (this.message.trim() !== '') {
                this.$emit('sendMessage', this.message);
                this.message = '';
            }
        },
    },
};
</script>

<style scoped></style>
