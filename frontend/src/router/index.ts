// src/router/index.ts

import { createRouter, createWebHistory } from 'vue-router';
import authRoutes from './modules/auth';
import settingsRoutes from './modules/settings';
import appRoutes from './modules/app';
import loungeRoutes from './modules/lounge';
import { useUserStore } from '@/store/user';
import type { UserState } from '@/store/user/state';
import { PATHS } from './paths';

const checkAuth = (store: UserState) => {
    const token = store.accessToken;
    if (!token && store.user) {
        store.user = null;
    }
    return !!token;
};

const routes = [
    ...authRoutes,
    ...settingsRoutes,
    ...appRoutes,
    ...loungeRoutes,
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

// 전역 가드
router.beforeEach((to, _from, next) => {
    const userStore = useUserStore();
    const requiresAuth = to.matched.some((r) => r.meta.requiresAuth);

    if (requiresAuth && !checkAuth(userStore)) {
        return next(`${PATHS.AUTH_BASE}/${PATHS.AUTH_LOGIN}`);
    }
    if (!requiresAuth && checkAuth(userStore) && to.path.startsWith('/auth')) {
        return next(PATHS.FRIENDS);
    }
    next();
});

export default router;
