// src/features/study/api/studiesApi.ts
// 이 파일은 스터디 관련 API 호출을 처리하는 모듈입니다.
// 실제 API 호출을 구현하는 대신 예시 데이터를 반환합니다.
// 이 예시에서는 스터디 목록을 가져오는 함수를 제공합니다.
// 이 파일은 TypeScript로 작성되었습니다.

export async function fetchStudies() {
    // 실제 API 호출 대신 예시 데이터를 반환합니다.
    return [
        {
            id: 1,
            name: 'Study A',
            icon: 'https://picsum.photos/200/300?random=2',
        },
        {
            id: 2,
            name: 'Study B',
            icon: 'https://picsum.photos/200/300?random=3',
        },
        {
            id: 3,
            name: 'Study C',
            icon: 'https://picsum.photos/200/300?random=4',
        },
    ];
}
