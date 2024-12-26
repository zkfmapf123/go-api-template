package configs

import (
	"log"
	"testing"
)

func Test_KafkaPing(t *testing.T) {
	_, err := NewKakfa()

	if err != nil {
		log.Fatalln(err)
	}
}