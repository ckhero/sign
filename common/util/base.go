/**
 *@Description
 *@ClassName base
 *@Date 2021/3/4 下午8:38
 *@Author ckhero
 */

package util

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"
)

func BuildSignKey(userId uint64, date time.Time) string {
	return fmt.Sprintf("u:year:month:%d:%d:%d", userId, date.Year(), date.Month())
}

func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

