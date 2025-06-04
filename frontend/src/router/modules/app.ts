import type { RouteRecordRaw } from 'vue-router';
import { AppLayout } from '@/shared/layout';
import { ChatPage } from '@/pages';
import { FriendListSidebar } from '@/widgets/sidebar';
import { useUserStore } from '@/store/user';
import { PATHS } from '@/router/paths';
import { LoungeListPage } from '@/pages';

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
                        ? `/${PATHS.FRIENDS}`
                        : `${PATHS.AUTH_BASE}/${PATHS.AUTH_LOGIN}`;
                },
            },
            {
                path: PATHS.LOUNGES,
                components: {
                    page: LoungeListPage,
                    sidebar: FriendListSidebar,
                },
            },
            {
                path: PATHS.FRIENDS,
                components: { page: ChatPage, sidebar: FriendListSidebar },
            },
        ],
    },
];

export default appRoutes;
