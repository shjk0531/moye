// src/store/member/action.ts
import type { MemberState } from './state';
import { fetchLoungeMembers } from '@/entities/member';
import type { Member } from '@/entities/member';

export const actions = {
    async loadLoungeMembers(this: MemberState, loungeId: string) {
        const response = await fetchLoungeMembers(loungeId);
        response.members.forEach((member) => {
            this.memberList[member.user_id] = member;
        });
        console.log('memberList:', this.memberList);
    },

    updateMember(this: MemberState, payload: Member) {
        this.memberList[payload.user_id] = payload;
    },

    setMembers(this: MemberState, payload: Member[]) {
        payload.forEach((member) => {
            this.memberList[member.user_id] = member;
        });
    },

    removeMember(this: MemberState, payload: string) {
        delete this.memberList[payload];
    },
};
