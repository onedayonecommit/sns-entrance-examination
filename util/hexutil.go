package util

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// 0x로 시작하는 hex string 반환 함수 => (매개변수로 원하는 길이 / 2) - 1 ex) 16자리 요청 == (16/2)-1 => 7
func GenerateHex(long int) (string, error)  {
	bytes := make([]byte, long)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return fmt.Sprintf("0x%s", hex.EncodeToString(bytes)), nil // 문자열 반환함수
}
