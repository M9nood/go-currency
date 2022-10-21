package currency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTHB(t *testing.T) {
	c := THB(1000)
	assert.Equal(t, "1,000", c.String())
	assert.Equal(t, "฿ 1,000", c.Currency())
	assert.Equal(t, "หนึ่งพันบาท", c.Text())

	c = THB(425.5)
	assert.Equal(t, "425.5", c.String())
	assert.Equal(t, "฿ 425.5", c.Currency())
	assert.Equal(t, "สี่ร้อยยี่สิบห้าบาทห้าสิบสตางค์", c.Text())
}
