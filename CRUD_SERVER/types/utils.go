package types

type ApiResponse struct {
	Description string      `json:"description"`
	Result      int16       `json:"result"`
	ErrCode     interface{} `json:"errCode"`
}

func NewApiResponse(description string, result int16, errCode interface{}) *ApiResponse {
	return &ApiResponse{
		Description: description,
		Result:      result,
		ErrCode:     errCode,
	}
}
