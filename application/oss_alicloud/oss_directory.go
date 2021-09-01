package OSSAlicloud

const dirBasic = "/sticker"

func DirUserRoot(userID string) string {
	return dirBasic + userID
}

func DirUserInfo(userID string) string {
	return DirUserRoot(userID) + "/info"
}

func PathUserInfoAvatar(userID string) string {
	return DirUserInfo(userID) + "/avatar"
}
