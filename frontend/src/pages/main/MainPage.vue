<template>
    <div class="bg-gray-850 text-gray-50">
        <h3># general</h3>
    </div>
    <div class="chat-messages bg-gray-850">
        <!-- ScrollPanel 대신에 채팅 메시지 영역에 스크롤을 적용하거나, 
                 ScrollPanel 내부에 100% 높이 지정 -->
        <ScrollPanel style="height: 100%">
            <div
                v-for="message in messages"
                :key="message.id"
                class="message-item"
            >
                <img :src="message.avatar" class="avatar" alt="avatar" />
                <div class="message-content">
                    <div class="message-header">
                        <span class="username text-gray-50 text-md">
                            {{ message.username }}</span
                        >
                        <span class="timestamp text-gray-300 text-sm">
                            {{ message.timestamp }}</span
                        >
                    </div>
                    <div
                        class="flex message-body text-gray-50 text-md justify-start"
                    >
                        <p>{{ message.text }}</p>
                    </div>
                </div>
            </div>
        </ScrollPanel>
    </div>

    <MessageInput @sendMessage="handleSendMessage" />
</template>

<script>
import MessageInput from './MessageInput.vue';

export default {
    name: 'ChatWindow',
    components: {
        MessageInput,
    },
    data() {
        return {
            messages: [
                {
                    id: 1,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User1',
                    timestamp: '10:00 AM',
                    text: 'Hello, world!',
                },
                {
                    id: 2,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User2',
                    timestamp: '10:05 AM',
                    text: 'Hi there!',
                },
                {
                    id: 1,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User1',
                    timestamp: '10:00 AM',
                    text: 'Hello, world!',
                },
                {
                    id: 2,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User2',
                    timestamp: '10:05 AM',
                    text: 'Hi there!',
                },
                {
                    id: 1,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User1',
                    timestamp: '10:00 AM',
                    text: 'Hello, world!',
                },
                {
                    id: 2,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User2',
                    timestamp: '10:05 AM',
                    text: 'Hi there!',
                },
                {
                    id: 1,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User1',
                    timestamp: '10:00 AM',
                    text: 'Hello, world!',
                },
                {
                    id: 2,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User2',
                    timestamp: '10:05 AM',
                    text: 'Hi there!',
                },
                {
                    id: 1,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User1',
                    timestamp: '10:00 AM',
                    text: 'Hello, world!',
                },
                {
                    id: 2,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User2',
                    timestamp: '10:05 AM',
                    text: 'Hi there!',
                },
                {
                    id: 1,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User1',
                    timestamp: '10:00 AM',
                    text: 'Hello, world!',
                },
                {
                    id: 2,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User2',
                    timestamp: '10:05 AM',
                    text: 'Hi there!',
                },
                {
                    id: 1,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User1',
                    timestamp: '10:00 AM',
                    text: 'Hello, world!',
                },
                {
                    id: 2,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User2',
                    timestamp: '10:05 AM',
                    text: 'Hi there!',
                },
                {
                    id: 1,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User1',
                    timestamp: '10:00 AM',
                    text: 'Hello, world!',
                },
                {
                    id: 2,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User2',
                    timestamp: '10:05 AM',
                    text: 'Hi there!',
                },
                {
                    id: 1,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User1',
                    timestamp: '10:00 AM',
                    text: 'Hello, world!',
                },
                {
                    id: 2,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User2',
                    timestamp: '10:05 AM',
                    text: 'Hi there!',
                },
                {
                    id: 1,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User1',
                    timestamp: '10:00 AM',
                    text: 'Hello, world!',
                },
                {
                    id: 2,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User2',
                    timestamp: '10:05 AM',
                    text: 'Hi there!',
                },
                {
                    id: 1,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User1',
                    timestamp: '10:00 AM',
                    text: 'Hello, world!',
                },
                {
                    id: 2,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User2',
                    timestamp: '10:05 AM',
                    text: 'Hi there!',
                },
                {
                    id: 1,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User1',
                    timestamp: '10:00 AM',
                    text: 'Hello, world!',
                },
                {
                    id: 2,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User2',
                    timestamp: '10:05 AM',
                    text: 'Hi there!',
                },
                {
                    id: 1,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User1',
                    timestamp: '10:00 AM',
                    text: 'Hello, world!',
                },
                {
                    id: 2,
                    avatar: 'https://via.placeholder.com/40',
                    username: 'User2',
                    timestamp: '10:05 AM',
                    text: 'Hi there!',
                },
                // 메시지가 많아질 경우에도 메시지 영역에만 스크롤이 발생합니다.
            ],
        };
    },
    methods: {
        handleSendMessage(messageText) {
            const newMessage = {
                id: this.messages.length + 1,
                avatar:
                    'https://picsum.photos/200/300?random=' +
                    (this.messages.length + 1),
                username: 'You',
                timestamp: new Date().toLocaleTimeString(),
                text: messageText,
            };
            this.messages.push(newMessage);
        },
    },
};
</script>

<style scoped>
/* 전체 채팅 영역이 뷰포트 전체를 차지하고, 오버플로우는 숨김 */
.chat-window {
    height: 100vh;
    overflow: hidden;
    display: grid;
    grid-template-rows: subgrid;
    grid-template-columns: subgrid;
    background-color: #36393f;
    position: relative;
}

/* 채팅 메시지 영역은 flex로 남은 공간을 채우고, 스크롤이 적용됨 */
.chat-messages {
    flex: 1;
    overflow-y: auto;
    overflow-x: hidden;
}

/* ScrollPanel 내부에 100% 높이를 부여하여 부모의 높이를 상속받도록 함 */
.chat-messages ::v-deep .p-scrollpanel-content {
    height: 100%;
}

/* ScrollPanel 기본 스타일 오버라이드 */
.chat-messages ::v-deep .p-scrollpanel-bar {
    background-color: #ccc; /* 연한 회색 배경 */
    border-radius: 4px;
    transition: background-color 0.3s ease;
}

/* 마우스 오버 시 조금 더 진한 회색 */
.chat-messages ::v-deep .p-scrollpanel-bar:hover {
    background-color: #aaa;
}

/* 수직 스크롤바(p-scrollpanel-bar-y)의 너비를 고정 */
.chat-messages ::v-deep .p-scrollpanel-bar-y {
    width: 8px !important;
}

/* 수평 스크롤바(p-scrollpanel-bar-x)의 높이를 고정 */
.chat-messages ::v-deep .p-scrollpanel-bar-x {
    height: 8px !important;
}

/* 메시지 항목 스타일 */
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
