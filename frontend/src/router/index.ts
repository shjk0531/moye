import { createRouter, createWebHistory } from 'vue-router';
import { AuthPage, MainPage } from '@/pages';
import AppLayout from '@/shared/layout/components/AppLayout.vue';

const routes = [
    {
        path: '/',
        component: AppLayout,
        children: [
            { path: '', component: MainPage },
            { path: '/login', component: AuthPage },
        ],
    },
];

export const router = createRouter({
    history: createWebHistory(),
    routes,
});
