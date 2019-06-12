package service

// CheckError 入力値がなかった場合に投げるエラー
type CheckError struct {
	Message string
	Cd      string
}

func (e *CheckError) Error() string {
	return e.Message
}

// DbError DBエラーの場合に投げるエラー
type DbError struct {
	Message string
}

func (e *DbError) Error() string {
	return e.Message
}
