import type { PanelGroup, PanelItem } from '@/shared/lib/panelUtils';

// 채널 그룹 타입 정의
export interface ChannelGroup extends PanelGroup {
    id: number;
    label: string;
    order: number;
}

// 채널 타입 정의
export interface Channel extends PanelItem {
    id: number;
    label: string;
    order: number;
    icon: string;
    groupId: number | null;
}
