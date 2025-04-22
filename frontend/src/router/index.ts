// src/router/index.ts

import { createRouter, createWebHistory } from 'vue-router';
import authRoutes from './modules/auth';
import settingsRoutes from './modules/settings';
import appRoutes from './modules/app';
import studyRoutes from './modules/study';
import { useUserStore } from '@/store/user';

const routes = [...authRoutes, ...settingsRoutes, ...appRoutes, ...studyRoutes];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

// 전역 가드
router.beforeEach((to, _from, next) => {
    const userStore = useUserStore();
    const requiresAuth = to.matched.some((r) => r.meta.requiresAuth);

    if (requiresAuth && !userStore.checkAuth()) {
        return next('/auth/login');
    }
    if (!requiresAuth && userStore.checkAuth() && to.path.startsWith('/auth')) {
        return next('/');
    }
    next();
});

export default router;
