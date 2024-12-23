package utils

import "informant-crypto/internal/services"

func BaseResponseBuilder(cur, rate, source, url string) services.BaseResponse {
	return services.BaseResponse{
		Currency: cur,
		Rate:     rate,
		Source:   source,
		URL:      url}
}
