<template>
    <PageLayout>
        <template #notice>
            <ChatNotice />
        </template>
        <template #content>
            <div
                class="message-area flex flex-col flex-1 overflow-y-auto overflow-x-hidden"
            >
                <ScrollPanel style="height: 100%">
                    <MessageItem
                        v-for="message in messages"
                        :key="message.id"
                        :message="message"
                    />
                </ScrollPanel>
            </div>
            <div class="chat-area pb-4 px-2">
                <MessageInput @sendMessage="handleSendMessage" />
            </div>
        </template>
        <template #sidebar>
            <MemberList />
        </template>
    </PageLayout>
</template>

<script setup lang="ts">
import { MessageInput, MessageItem } from '@/features/chat';
import { ChatNotice } from '@/widgets/notice';
import { MemberList } from '@/widgets/sidebar';
import { PageLayout } from '@/shared/layout';
import { ref } from 'vue';

interface Message {
    id: number;
    avatar: string;
    username: string;
    timestamp: string;
    text: string;
}

const messages = ref<Message[]>([
    {
        id: 1,
        avatar: 'https://picsum.photos/200/300?random=1',
        username: '김민수',
        timestamp: '10:00 AM',
        text: '안녕하세요! 오늘 스터디 시작하기 전에 간단히 자기소개 부탁드립니다.',
    },
    {
        id: 2,
        avatar: 'https://picsum.photos/200/300?random=2',
        username: '이지은',
        timestamp: '10:05 AM',
        text: '안녕하세요! 저는 이지은이라고 합니다. 프론트엔드 개발자로 일하고 있고, Vue.js에 관심이 많습니다.',
    },
    {
        id: 3,
        avatar: 'https://picsum.photos/200/300?random=3',
        username: '박준호',
        timestamp: '10:10 AM',
        text: '안녕하세요! 저는 박준호입니다. 백엔드 개발자이고, 현재 Spring Boot를 주로 사용하고 있어요.',
    },
    {
        id: 4,
        avatar: 'https://picsum.photos/200/300?random=4',
        username: '최유진',
        timestamp: '10:15 AM',
        text: '안녕하세요! 저는 최유진입니다. 취준생이고, 웹 개발에 관심이 많아서 참여하게 되었습니다.',
    },
    {
        id: 5,
        avatar: 'https://picsum.photos/200/300?random=5',
        username: '김민수',
        timestamp: '10:20 AM',
        text: '오늘은 프로젝트 기획서 작성에 대해 논의하고 싶은데, 각자 아이디어가 있으신가요?',
    },
    {
        id: 6,
        avatar: 'https://picsum.photos/200/300?random=6',
        username: '이지은',
        timestamp: '10:25 AM',
        text: '저는 스터디원들의 학습 진도와 목표를 관리할 수 있는 대시보드를 만들어보면 좋을 것 같아요.',
    },
    {
        id: 7,
        avatar: 'https://picsum.photos/200/300?random=7',
        username: '박준호',
        timestamp: '10:30 AM',
        text: '좋은 아이디어네요! 그럼 백엔드에서는 사용자 인증과 데이터베이스 설계부터 시작하면 될 것 같습니다.',
    },
    {
        id: 8,
        avatar: 'https://picsum.photos/200/300?random=8',
        username: '최유진',
        timestamp: '10:35 AM',
        text: '저도 동의합니다. 그리고 모바일 환경에서도 잘 작동하도록 반응형으로 구현하면 좋을 것 같아요.',
    },
    {
        id: 9,
        avatar: 'https://picsum.photos/200/300?random=9',
        username: '김민수',
        timestamp: '10:40 AM',
        text: '좋은 의견들이네요. 그럼 다음 주까지 각자 담당할 부분 정리해서 공유해주시면 좋겠습니다.',
    },
    {
        id: 10,
        avatar: 'https://picsum.photos/200/300?random=10',
        username: '이지은',
        timestamp: '10:45 AM',
        text: '네, 프론트엔드 부분은 제가 담당하도록 하겠습니다. Vue.js와 TypeScript를 활용해서 구현해보겠습니다.',
    },
    {
        id: 11,
        avatar: 'https://picsum.photos/200/300?random=11',
        username: '박준호',
        timestamp: '10:50 AM',
        text: '백엔드는 Spring Boot로 RESTful API를 구현하고, PostgreSQL을 데이터베이스로 사용하는 건 어떨까요?',
    },
    {
        id: 12,
        avatar: 'https://picsum.photos/200/300?random=12',
        username: '최유진',
        timestamp: '10:55 AM',
        text: '저는 UI/UX 디자인과 문서화를 담당하도록 하겠습니다. Figma로 디자인 시안을 준비해볼게요.',
    },
]);

const handleSendMessage = (messageText: string) => {
    messages.value.push({
        id: messages.value.length + 1,
        avatar:
            'https://picsum.photos/200/300?random=0' +
            (messages.value.length + 1),
        username: 'You',
        timestamp: new Date().toLocaleTimeString(),
        text: messageText,
    });
};
</script>

<style scoped>
.message-area :deep(.p-scrollpanel-content) {
    height: 100%;
}

.message-area :deep(.p-scrollpanel-bar) {
    background-color: var(--custom-scrollbar-color);
    border-radius: 4px;
    transition: background-color 0.3s ease;
}

.message-area :deep(.p-scrollpanel-bar:hover) {
    background-color: var(--custom-scrollbar-color-hover);
}

.message-area :deep(.p-scrollpanel-bar:active) {
    background-color: var(--custom-scrollbar-color-active);
}

.message-area :deep(.p-scrollpanel-bar-y) {
    width: calc(var(--spacing) * 2) !important;
}
</style>
