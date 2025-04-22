import type { RouteRecordRaw } from 'vue-router';
import { SettingLayout } from '@/shared/layout';
import { SettingPage } from '@/pages';

const settingsRoutes: RouteRecordRaw[] = [
    {
        path: '/settings',
        component: SettingLayout,
        meta: { requiresAuth: true },
        props: (route) => ({ section: route.params.section }),
        children: [
            {
                path: '',
                name: 'Settings',
                component: SettingPage,
            },
            { path: 'account', component: SettingPage },
            { path: 'profile', component: SettingPage },
            { path: 'display', component: SettingPage },
            { path: 'notifications', component: SettingPage },
            { path: 'privacy', component: SettingPage },
            { path: 'theme', component: SettingPage },
            { path: 'language', component: SettingPage },
        ],
    },
];

export default settingsRoutes;
