package constant

const (
	AppVersion      = "1.0.0"
	PrefixEnv       = "OCR_"
	LogModule       = "OCR"
	InternalJwtTime = 8640000

	ERR = 1
	OK  = 0

	KTP    = "KTP"
	TERANG = "TERANG"
	GELAP  = "GELAP"
	OFF    = "OFF"
	AUTO   = "AUTO"

	ThresholdKTP     = 55000
	ThresholdGelap   = 55800
	ThresholdBalance = 65000
	ThresholdTerang  = 65000
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
