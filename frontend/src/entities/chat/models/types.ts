export interface ChatMessage {
    id: string;
    content: string;
    senderId: string;
    senderName: string;
    profileImage: string;
    type: 'text' | 'image' | 'file';
    createdAt: string;
}
