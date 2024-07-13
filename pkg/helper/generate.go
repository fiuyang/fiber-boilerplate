package helper

import (
	"fmt"
	"strconv"
)

func GenerateShipmentNo(lastShipmentNo string) string {
	if lastShipmentNo == "" {
		return "000001"
	}

	lastNo, err := strconv.Atoi(lastShipmentNo)
	if err != nil {
		return "000001"
	}

	newNo := lastNo + 1
	newShipmentNo := fmt.Sprintf("%06d", newNo)

	return newShipmentNo
}
