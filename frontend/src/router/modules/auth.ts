import type { RouteRecordRaw } from 'vue-router';
import { AuthLayout } from '@/shared';
import { LoginPage, SignupPage, FindPasswordPage } from '@/pages';

const authRoutes: RouteRecordRaw[] = [
    {
        path: '/auth',
        component: AuthLayout,
        meta: { requiresAuth: false },
        children: [
            { path: 'login', component: LoginPage },
            { path: 'signup', component: SignupPage },
            { path: 'find-password', component: FindPasswordPage },
        ],
    },
];

export default authRoutes;
