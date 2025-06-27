import type { MemberState } from './state';

export const getters = {
    getMemberList: (state: MemberState) => state.memberList,

    getMemberById: (state: MemberState) => (memberId: string) => {
        return state.memberList[memberId];
    },
};
