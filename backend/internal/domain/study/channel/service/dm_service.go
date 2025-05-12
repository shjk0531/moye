package service


func SortUserIDs(uid1, uid2 string) []string {
    if uid1 < uid2 {
        return []string{uid1, uid2}
    }
    return []string{uid2, uid1}
}
