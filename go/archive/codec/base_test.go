package codec

import (
	"fmt"
	"testing"
)

func TestCodec_deserialize(t *testing.T) {
	c := Constructor()
	data := "[]"
	result := c.deserialize(data)
	fmt.Println(c.serialize(result))
}
