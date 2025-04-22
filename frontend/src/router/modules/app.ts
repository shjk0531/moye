import type { RouteRecordRaw } from 'vue-router';
import { AppLayout } from '@/shared/layout';
import { ChatPage, CalendarPage } from '@/pages';
import {
    FriendListSidebar,
    ChannelListSidebar,
    CalendarListSidebar,
} from '@/widgets/sidebar';
import { useUserStore } from '@/store/user';
import { PATHS } from '@/router/paths';

const appRoutes: RouteRecordRaw[] = [
    {
        path: '',
        component: AppLayout,
        meta: { requiresAuth: true },
        children: [
            {
                path: '',
                redirect: () => {
                    const store = useUserStore();
                    return store.checkAuth()
                        ? `/${PATHS.ME}`
                        : `${PATHS.AUTH_BASE}/${PATHS.AUTH_LOGIN}`;
                },
            },
            {
                path: PATHS.STUDIES,
                components: { page: ChatPage, sidebar: FriendListSidebar },
            },
            {
                path: PATHS.ME,
                components: { page: ChatPage, sidebar: FriendListSidebar },
            },
        ],
    },
];

export default appRoutes;
