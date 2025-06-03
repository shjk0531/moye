import type { RouteRecordRaw } from 'vue-router';
import { AppLayout } from '@/shared/layout';
import { ChatPage, CalendarPage, CreateLoungePage } from '@/pages';
import { ChannelListSidebar, CalendarListSidebar } from '@/widgets/sidebar';
import { PATHS } from '@/router/paths';

const loungeRoutes: RouteRecordRaw[] = [
    {
        path: PATHS.STUDY_BASE + `/:${PATHS.STUDY_PARAM}`,
        component: AppLayout,
        meta: { requiresAuth: true },
        children: [
            {
                path: `:${PATHS.STUDY_CHANNEL}/:${PATHS.STUDY_CHANNEL_PARAM}`,
                components: { page: ChatPage, sidebar: ChannelListSidebar },
            },
            {
                path: `:${PATHS.STUDY_CALENDAR}/:${PATHS.STUDY_CALENDAR_PARAM}`,
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
