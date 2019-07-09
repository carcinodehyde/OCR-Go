package constant

const (
	APP_VERSION       = "1.0.0"
	PREFIX_ENV        = "OCR_"
	LOG_MODULE        = "OCR"
	INTERNAL_JWT_TIME = 8640000

	ERR = 1
	OK  = 0

	KTP    = "KTP"
	TERANG = "TERANG"
	GELAP  = "GELAP"
	OFF    = "OFF"
	AUTO   = "AUTO"
)

type GenericResponse struct {
	Status   int         `json:"status"`
	Success  bool        `json:"success"`
	Messages []string    `json:"messages"`
	Data     interface{} `json:"data"`
}

func NewGenericResponse(stsCd, isError int, messages []string, data interface{}) *GenericResponse {

	return &GenericResponse{
		Status:   stsCd,
		Success:  isError == 0,
		Messages: messages,
		Data:     data,
	}
}
