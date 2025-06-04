package dto

// UserLoungesResponse는 사용자가 속한 라운지 목록을 반환하는 응답 DTO입니다.
type UserLoungesResponse struct {
	Lounges []LoungeListDTO `json:"lounges"`
}

// LoungeListResponse는 라운지 목록을 조회하는 응답 DTO
type LoungeListResponse struct {
	Lounges []LoungeListDTO `json:"lounges"`
}


// LoungeIconResponse는 라운지 아이콘 목록을 조회하는 응답 DTO
type LoungeIconResponse struct {
	Icons []LoungeIconDTO `json:"icons"`
}