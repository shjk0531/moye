export interface StudyChannelResponse {
    items: StudyChannelItem[];
}

export interface StudyChannelItem {
    itemType: 'channel' | 'group';
    channel?: ChannelDTO;
    group?: ChannelGroupDTO;
}

export interface ChannelDTO {
    id: string;
    name: string;
}

export interface ChannelGroupDTO {
    id: string;
    name: string;
    channels: ChannelDTO[];
}
