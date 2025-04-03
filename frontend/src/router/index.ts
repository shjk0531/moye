import { createRouter, createWebHistory } from 'vue-router';
import { AuthPage, MainPage, ChatPage, LoginPage, SignupPage } from '@/pages';
import AppLayout from '@/shared/layout/components/AppLayout.vue';
import { ChannelListSidebar } from '@/widgets/sidebar';

const routes = [
    {
        path: '/auth',
        children: [
            {
                path: 'login',
                component: LoginPage,
            },
            {
                path: 'signup',
                component: SignupPage,
            },
        ],
    },
    {
        path: '/',
        component: AppLayout,
        children: [
            {
                path: '/',
                components: {
                    page: ChatPage,
                    leftSide: ChannelListSidebar,
                },
            },
            {
                path: '/me',
                components: {
                    page: ChatPage,

                    leftSide: ChannelListSidebar,
                },
            },
            {
                path: '/login',
                components: { page: AuthPage, leftSide: ChannelListSidebar },
            },
            {
                path: '/study/:studyId/channel/:channelId?',
                components: { page: ChatPage, leftSide: ChannelListSidebar },
            },
            {
                path: '/study/:studyId',
                components: { page: ChatPage, leftSide: ChannelListSidebar },
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
    // 로그인 없이 접근이 허용된 경로를 배열로 정의
    const publicPaths = ['/login', '/signup'];

    // 현재 접근하려는 경로가 publicPaths에 포함되어 있는지 체크
    const isPublicRoute = publicPaths.includes(to.path);

    // localStorage나 다른 저장소에서 JWT 토큰을 확인 (여기서는 localStorage 사용)
    const jwtToken = localStorage.getItem('jwt');
    const isLoggedIn = !!jwtToken; // 토큰이 있으면 true

    // 만약 사용자가 로그인하지 않았고(publicPaths가 아닌) 접근하려고 한다면
    if (!isLoggedIn && !isPublicRoute) {
        // /login 페이지로 리디렉션
        // return next('/login');
    }

    // 조건을 만족하면 정상적으로 라우팅 진행
    next();
});

export default router;
