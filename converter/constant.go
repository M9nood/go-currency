package converter

const (
	MILLION_POSITION = 6
	TEN_POSITION     = 1
	UNIT_POSITION    = 0
)

type CurrencyFormat string

const (
	CURRENCY_FORMAT_THB CurrencyFormat = "thb"
)

var (
	THAI_NUMBER_TEXTS = []string{"ศูนย์", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า", "สิบ"}
	THAI_UNIT_TEXTS   = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน", "ล้าน"}
)
