package usb

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Open_Honeywell_1900(t *testing.T) {

	hids, err := enumerateHid(0x0c2e, 0x0907)
	assert.Nil(t, err)
	assert.NotNil(t, hids)

	d , err := hids[0].Open()
	assert.Nil(t, err)
	buf := make([]byte, 2048)
	for ; ;  {
		n , err := d.Read(buf)
		assert.NotNil(t, err)
		fmt.Printf("read %d bytes\n", n)
		fmt.Println(string(buf))
	}

}

func Test_Open_Datalogic_gbt4400(t *testing.T) {

	hids, err := enumerateHid(0x05f9, 0x220a)
	assert.Nil(t, err)
	assert.NotNil(t, hids)

	d , err := hids[0].Open()
	assert.NotNil(t, d)
	assert.Nil(t, err)
	if err != nil {
		t.FailNow()
	}
	buf := make([]byte, 2048)
	for ; ;  {
		n , err := d.Read(buf)
		assert.NotNil(t, err)
		fmt.Printf("read %d bytes\n", n)
		fmt.Println(string(buf))
	}

}

