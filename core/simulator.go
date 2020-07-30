package core

import "fmt"

const (
	STAGE_ACCOUNT_CREATE = 1
	STAGE_BURN_CONSENT   = 2
	STAGE_TRANSFER       = 3
	STAGE_BURN_EXEC      = 4
	STAGE_AIRDROP        = 5
)

type (
	Cycle struct {
		ID         uint64 `gorm:"AUTO_INCREMENT"`
		Stage      uint64
		StartIndex uint64
		EndIndex   uint64
	}
)

func (db *DB) LogCycle(stage, start, end uint64) error {
	fmt.Println("Logging a cycle", stage, start, end)
	cycle := Cycle{Stage: stage, StartIndex: start, EndIndex: end}
	return db.Instance.Create(&cycle).Error
}

func (db *DB) GetLastCycle() (cycle Cycle, err error) {
	if err := db.Instance.Order("id desc").First(&cycle).Error; err != nil {
		return cycle, err
	}
	return cycle, nil
}

func (db *DB) CycleCount() (count uint64, err error) {
	if err := db.Instance.Table("cycles").Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}
