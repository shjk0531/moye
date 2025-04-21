/** 세부 설정 항목 */
export interface SettingItem {
    /** i18n 키 혹은 라벨용 문자열 */
    name: string;
    /** 라우트 path or name */
    route: string;
    /** (선택) lazy-load할 컴포넌트 */
    component?: any;
    icon?: string;
}

/** 그룹별 설정 메뉴 */
export interface SettingGroup {
    /** 그룹 라벨(i18n 키) */
    title?: string;
    /** 그룹 내 항목들 */
    items: SettingItem[];
}

export const settingsList: SettingGroup[] = [
    {
        title: '사용자 설정', // ex: “내 정보” (i18n 키)
        items: [
            { name: 'account', route: 'account' },
            { name: 'profile', route: 'profile' },
        ],
    },
    {
        title: '앱 설정', // ex: “기기 설정”
        items: [
            { name: 'display', route: 'display' },
            { name: 'notifications', route: 'notifications' },
            { name: 'privacy', route: 'privacy' },
            { name: 'theme', route: 'theme' },
            { name: 'language', route: 'language' },
        ],
    },
];
