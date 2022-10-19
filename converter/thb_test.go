package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToThaiBahtText(t *testing.T) {
	assert.Equal(t, "หนึ่งร้อยยี่สิบสี่ล้านล้านสี่แสนบาท", ConvertToThaiBahtText(124000000400000))
	assert.Equal(t, "ยี่สิบเอ็ดบาท", ConvertToThaiBahtText(21))
	assert.Equal(t, "ห้าร้อยหนึ่งบาท", ConvertToThaiBahtText(501))
	assert.Equal(t, "ยี่สิบสตางค์", ConvertToThaiBahtText(0.2))
	assert.Equal(t, "สิบสตางค์", ConvertToThaiBahtText(0.1))
	assert.Equal(t, "สี่สิบบาทหนึ่งสตางค์", ConvertToThaiBahtText(40.01))
	assert.Equal(t, "เก้าหมื่นห้าพันสามร้อยแปดสิบสองบาทห้าสิบสตางค์", ConvertToThaiBahtText(95382.5))
	assert.Equal(t, "สี่สิบบาทจุดสองศูนย์หนึ่งสตางค์", ConvertToThaiBahtText(40.201))
	assert.Equal(t, "ศูนย์บาท", ConvertToThaiBahtText(0))
	assert.Equal(t, "สามร้อยสิบบาท", ConvertToThaiBahtText(310))
	assert.Equal(t, "หนึ่งแสนสองหมื่นสี่พันห้าล้านหกแสนเจ็ดหมื่นสี่พันสามร้อยสี่สิบล้านสี่แสนบาท", ConvertToThaiBahtText(124005674340400000))
}
