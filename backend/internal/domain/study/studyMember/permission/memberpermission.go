package permission

const (
    PermManageSchedule   int64 = 1 << 0 // 1 - 일정 관리
    PermCreateChannel    int64 = 1 << 1 // 2 - 채팅방 생성
    PermWriteNotice      int64 = 1 << 2 // 4 - 공지 작성
    PermKickMember       int64 = 1 << 3 // 8 - 멤버 강퇴
    PermManageInvite     int64 = 1 << 4 // 16 - 초대링크 생성
    PermReorderChannels  int64 = 1 << 5 // 32 - 채널 순서 변경
    PermManageRoles      int64 = 1 << 6 // 64 - 직위 생성/수정
)

func HasPermission(flags int64, permission int64) bool {
    return flags&permission != 0
}
