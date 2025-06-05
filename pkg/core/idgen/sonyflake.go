package idgen

import (
	"github.com/sony/sonyflake"
	"log"
)

var sf *sonyflake.Sonyflake

func Init(machineID uint16) {
	settings := sonyflake.Settings{
		MachineID: func() (uint16, error) {
			return machineID, nil
		},
	}
	sf = sonyflake.NewSonyflake(settings)
	if sf == nil {
		log.Fatal("failed to initialize sonyflake")
	}
}

func NextID() (uint64, error) {
	return sf.NextID()
}
