package service

import "context"

// GetCreatorIDFromContext はContextからセッション作成者のIDを取得します。
func GetUserIDFromContext(ctx context.Context) (string, bool) {
	v := ctx.Value("userID")
	userID, ok := v.(string)
	return userID, ok
}

// SetUserIDToContext はユーザIDをContextにセットします。
func SetUserIDToContext(ctx context.Context, userID string) context.Context {
	if userID != "" {
		return context.WithValue(ctx, "userID", userID)
	}
	return ctx
}