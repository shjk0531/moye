import { createRouter, createWebHistory } from 'vue-router';
import { AuthPage, MainPage, ChatPage } from '@/pages';
import AppLayout from '@/shared/layout/components/AppLayout.vue';

const routes = [
    {
        path: '/',
        component: AppLayout,
        children: [
            { path: '/', component: ChatPage },
            { path: '/me', component: MainPage },
            { path: '/login', component: AuthPage },
            { path: '/study/:studyId/:chennalId', component: ChatPage },
        ],
    },
];

export const router = createRouter({
    history: createWebHistory(),
    routes,
});
