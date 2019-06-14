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
	err     error
}

func (e *DbError) Error() string {
	return e.Message
}

// ErrorMessage DBエラーの場合に投げるエラー
type ErrorMessage struct {
	StatusCd int // httpステータスコードを設定
	Message  string
	ErrorCd  string // エラーコード
	Detail   error  // 発生したエラーを設定
}

func (e *ErrorMessage) Error() string {
	return e.Message
}
