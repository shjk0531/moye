import { createRouter, createWebHistory } from 'vue-router';
import {
    ChatPage,
    LoginPage,
    SignupPage,
    CalendarPage,
    FindPasswordPage,
} from '@/pages';
import { ChannelListSidebar, CalendarListSidebar } from '@/widgets/sidebar';
import { AuthLayout, AppLayout, PageLayout } from '@/shared/layout';
import { useUserStore } from '@/store/user';
const routes = [
    {
        path: '/auth',
        component: AuthLayout,
        children: [
            {
                path: 'login',
                component: LoginPage,
                meta: { requiresAuth: false },
            },
            {
                path: 'signup',
                component: SignupPage,
                meta: { requiresAuth: false },
            },
            {
                path: 'find-password',
                component: FindPasswordPage,
                meta: { requiresAuth: false },
            },
        ],
    },
    {
        path: '/',
        component: AppLayout,
        meta: { requiresAuth: true },
        children: [
            {
                path: '',
                components: {
                    page: ChatPage,
                    sidebar: ChannelListSidebar,
                },
            },
            {
                path: 'me',
                components: {
                    page: ChatPage,
                    sidebar: ChannelListSidebar,
                },
            },
            {
                path: 'study/:studyId/channel/:channelId?',
                components: {
                    page: ChatPage,
                    sidebar: ChannelListSidebar,
                },
            },
            {
                path: 'study/:studyId/calendar/:calendarId?',
                components: {
                    page: CalendarPage,
                    sidebar: CalendarListSidebar,
                },
            },
        ],
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

// Global Navigation Guard 등록
router.beforeEach((to, _from, next) => {
    const userStore = useUserStore();
    const requiresAuth = to.matched.some((record) => record.meta.requiresAuth);

    // 인증이 필요한 페이지에 접근할 때 로그인 여부 확인
    if (requiresAuth && !userStore.checkAuth()) {
        // 로그인 페이지로 리디렉션
        return next('/auth/login');
    }

    // 이미 로그인된 사용자가 인증 페이지(로그인, 회원가입 등)에 접근하려고 할 때
    if (!requiresAuth && userStore.checkAuth() && to.path.startsWith('/auth')) {
        // 메인 페이지로 리디렉션
        return next('/');
    }

    // 조건을 만족하면 정상적으로 라우팅 진행
    next();
});

export default router;
