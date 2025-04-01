<template>
    <div class="flex border-none rounded-lg p-2 gap-2 mr-2 h-12">
        <InputText
            v-model="message"
            autoResize
            placeholder="Message #general"
            class="!bg-gray-800 !border-none !text-white !focus:ring-0 !focus:shadow-none !focus:border-none !p-3 !flex-1"
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
