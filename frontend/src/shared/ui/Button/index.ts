import PrimeButton from 'primevue/button';
export { PrimeButton as Button };

export { default as ChannelButton } from './notice/ChannelButton.vue';
export { default as CalendarButton } from './notice/CalendarButton.vue';
export { default as PostButton } from './notice/PostButton.vue';
export { default as NoteButton } from './notice/NoteButton.vue';
export { default as MemberButton } from './notice/MemberButton.vue';

export const noticeImporters = {
    ChannelButton: () => import('./notice/ChannelButton.vue'),
    CalendarButton: () => import('./notice/CalendarButton.vue'),
    PostButton: () => import('./notice/PostButton.vue'),
    NoteButton: () => import('./notice/NoteButton.vue'),
    MemberButton: () => import('./notice/MemberButton.vue'),
};
