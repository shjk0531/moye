<template>
    <div
        class="flex items-center p-2 bg-gray-700 text-gray-50 h-(--custom-bottom-pannel-height) rounded-lg p-2"
    >
        <InputText
            v-model="message"
            autoResize
            placeholder="Message #general"
            class="!bg-gray-700 !border-none !focus:ring-0 !focus:shadow-none !focus:border-none !flex-1"
            @keyup.enter="send"
        />
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
