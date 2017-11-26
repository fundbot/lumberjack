package download

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_File(t *testing.T) {
	File("http://example.qutheory.io/json", "", "test")

	data, _ := ioutil.ReadFile("./build/data/test")
	assert.Equal(t, string(data), "{\"array\":[0,1,2,3],\"dict\":{\"lang\":\"Swift\",\"name\":\"Vapor\"},\"number\":123,\"string\":\"test\"}")
}
