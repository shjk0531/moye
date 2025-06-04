export interface Lounge {
    id?: string;
    name: string;
    profileUrl: string;
    description: string;
    leaderId?: string;
    createdAt?: string;
    updatedAt?: string;
    content?: string;
    tags?: string[];
}

export interface LoungeCreatePayload {
    name: string;
    profile_url: string;
    description: string;
    content: string;
    tags: string[];
}

export interface LoungeIcon {
    id: string;
    name: string;
    profile_url: string;
}

export interface LoungeIconResponse {
    icons: LoungeIcon[];
}

export interface LoungeInfo {
    id: string;
    name: string;
    profile_url: string;
    description: string;
    tags: string[];
    updated_at: string;
}

export interface LoungeListResponse {
    lounges: LoungeInfo[];
}
