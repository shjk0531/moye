import type { RouteRecordRaw } from 'vue-router';
import { AppLayout } from '@/shared/layout';
import { ChatPage, CalendarPage, CreateLoungePage } from '@/pages';
import { ChannelListSidebar, CalendarListSidebar } from '@/widgets/sidebar';
import { PATHS } from '@/router/paths';

const loungeRoutes: RouteRecordRaw[] = [
    {
        path: PATHS.LOUNGE_BASE + `/:${PATHS.LOUNGE_PARAM}`,
        component: AppLayout,
        meta: { requiresAuth: true },
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
