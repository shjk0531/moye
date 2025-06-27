// src/store/member/state.ts
import type { Member } from '@/entities/member';

export interface MemberList {
    [memberId: string]: Member;
}

export interface MemberState {
    memberList: MemberList;
}

export const state = (): MemberState => ({
    memberList: {},
});
