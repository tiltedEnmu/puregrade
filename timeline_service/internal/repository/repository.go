package repository

type Repository interface {
	Push(userId string, postIds ...string) (int64, error)
	GetRange(userId string, start, end int64) ([]string, error)
	GetIndexByPostId(userId, postId string, maxLen int64) (int64, error)
	Trim(userId string, cap int32) (string, error)
}
