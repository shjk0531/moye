export interface Study {
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

export interface StudyCreatePayload {
    name: string;
    profile_url: string;
    description: string;
    content: string;
    tags: string[];
}

export interface StudyIcon {
    id: string;
    name: string;
    profile_url: string;
}

export interface StudyIconResponse {
    icons: StudyIcon[];
}
