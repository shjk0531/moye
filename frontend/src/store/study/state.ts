export interface StudyState {
    // studyId별 각 목록(listType별)의 마지막 활성 항목
    // 예: { '1': { channel: '채널1의 id', calendar: '캘린더2의 id' } }
    activeItems: Record<string, Record<string, string>>;
    studyName: string;
    studyIcon: string;
}

export const state = (): StudyState => ({
    activeItems: {},
    studyName: '',
    studyIcon: '',
});
