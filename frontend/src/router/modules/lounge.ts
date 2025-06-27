import type { RouteRecordRaw } from 'vue-router';
import { AppLayout } from '@/shared/layout';
import { ChatPage, CalendarPage, CreateLoungePage } from '@/pages';
import { ChannelListSidebar, CalendarListSidebar } from '@/widgets/sidebar';
import { PATHS } from '@/router/paths';
import { useMemberStore } from '@/store/member';

const loungeRoutes: RouteRecordRaw[] = [
    {
        path: PATHS.LOUNGE_BASE + `/:${PATHS.LOUNGE_PARAM}`,
        component: AppLayout,
        meta: { requiresAuth: true },
        beforeEnter: async (to, _from, next) => {
            const loungeId = String(to.params[PATHS.LOUNGE_PARAM]);
            const memberStore = useMemberStore();
            try {
                await memberStore.loadLoungeMembers(loungeId);
                next();
            } catch (err) {
                console.error('멤버 로드 실패', err);
                next(); // 실패해도 라우팅은 계속
            }
        },
        children: [
            {
                path: `:${PATHS.LOUNGE_CHANNEL}/:${PATHS.LOUNGE_CHANNEL_PARAM}`,
                name: 'chat',
                components: { page: ChatPage, sidebar: ChannelListSidebar },
            },
            {
                path: `:${PATHS.LOUNGE_CALENDAR}/:${PATHS.LOUNGE_CALENDAR_PARAM}`,
                name: 'calendar',
                components: {
                    page: CalendarPage,
                    sidebar: CalendarListSidebar,
                },
            },
        ],
    },
    {
        path: `/${PATHS.CREATE}`,
        component: CreateLoungePage,
        meta: { requiresAuth: true },
    },
];

export default loungeRoutes;
