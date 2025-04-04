export async function fetchCalendarsGrouped() {
    return [
        {
            id: 1,
            label: '달력 그룹1',
            order: 2,
        },
        {
            id: 2,
            label: '달력 그룹2',
            order: 3,
        },
    ];
}

export async function fetchCalendarsUngrouped() {
    return [
        { id: 0, label: '일정1', order: 1, icon: 'calendar', groupId: 1 },
        { id: 1, label: '일정2', order: 2, icon: 'calendar', groupId: 1 },
        { id: 2, label: '일정3', order: 1, icon: 'calendar', groupId: null },
        { id: 3, label: '일정4', order: 4, icon: 'calendar', groupId: null },
        { id: 4, label: '일정5', order: 5, icon: 'calendar', groupId: null },
        { id: 5, label: '일정6', order: 6, icon: 'calendar', groupId: null },
        { id: 6, label: '일정7', order: 1, icon: 'calendar', groupId: 2 },
        { id: 7, label: '일정8', order: 2, icon: 'calendar', groupId: 2 },
        { id: 8, label: '일정9', order: 3, icon: 'calendar', groupId: 2 },
        { id: 9, label: '일정10', order: 4, icon: 'calendar', groupId: 2 },
        { id: 10, label: '일정11', order: 5, icon: 'calendar', groupId: 2 },
    ];
}
