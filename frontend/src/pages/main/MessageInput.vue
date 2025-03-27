<template>
    <div class="message-input-container">
        <InputText
            v-model="message"
            autoResize
            placeholder="Message #general"
            class="message-input"
        />
        <Button label="Send" @click="send" class="send-button" />
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

<style scoped>
.message-input-container {
    display: flex;
    align-items: center;
}

.message-input {
    flex: 1;
    margin-right: 10px;
    background-color: #40444b;
    border: none;
    color: #dcddde;
    padding: 10px;
    border-radius: 5px;
}

.send-button {
    background-color: #7289da;
    border: none;
    color: white;
    padding: 10px 20px;
    border-radius: 5px;
}
</style>
