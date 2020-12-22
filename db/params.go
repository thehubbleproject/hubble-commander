package db

import "github.com/BOPR/core"

func (db *DB) InitSyncStatus(startBlock uint64) error {
	return db.Instance.Create(&core.SyncStatus{LastEthBlockRecorded: startBlock}).Error
}

func (db *DB) UpdateSyncStatusWithBlockNumber(blkNum uint64) error {
	syncStatus, err := db.GetSyncStatus()
	if err != nil {
		return err
	}
	var updatedSyncStatus core.SyncStatus
	updatedSyncStatus.LastEthBlockRecorded = blkNum
	if err := db.Instance.Table("sync_statuses").Where("id = ?", syncStatus.ID).Update(&updatedSyncStatus).Error; err != nil {
		return err
	}
	return nil
}

func (db *DB) GetSyncStatus() (status core.SyncStatus, err error) {
	if err := db.Instance.First(&status).Error; err != nil {
		return status, err
	}
	return status, nil
}

// UpdateStakeAmount updates the stake amount
func (db *DB) UpdateStakeAmount(newStakeAmount uint64) error {
	var updatedParams core.Params
	updatedParams.StakeAmount = newStakeAmount
	if err := db.Instance.Table("params").Assign(core.Params{StakeAmount: newStakeAmount}).FirstOrCreate(&updatedParams).Error; err != nil {
		return err
	}
	return nil
}

// UpdateFinalisationTimePerBatch updates the max depth
func (db *DB) UpdateFinalisationTimePerBatch(finalisationDuration uint64) error {
	var updatedParams core.Params
	updatedParams.MaxDepth = finalisationDuration
	if err := db.Instance.Table("params").Assign(core.Params{FinalisationTime: finalisationDuration}).FirstOrCreate(&updatedParams).Error; err != nil {
		return err
	}
	return nil
}

// UpdateMaxDepth updates the max depth
func (db *DB) UpdateMaxDepth(newDepth uint64) error {
	var updatedParams core.Params
	updatedParams.MaxDepth = newDepth
	if err := db.Instance.Table("params").Assign(core.Params{MaxDepth: newDepth}).FirstOrCreate(&updatedParams).Error; err != nil {
		return err
	}
	return nil
}

// UpdateDepositSubTreeHeight updates the max height of deposit sub tree
func (db *DB) UpdateDepositSubTreeHeight(newHeight uint64) error {
	var updatedParams core.Params
	updatedParams.MaxDepositSubTreeHeight = newHeight
	if err := db.Instance.Table("params").Assign(core.Params{MaxDepositSubTreeHeight: newHeight}).FirstOrCreate(&updatedParams).Error; err != nil {
		return err
	}
	return nil
}

func (db *DB) GetBatchFinalisationTime() (uint64, error) {
	var params core.Params
	if err := db.Instance.First(&params).Error; err != nil {
		return 0, err
	}
	return params.FinalisationTime, nil
}

// GetParams gets params from the DB
func (db *DB) GetParams() (params core.Params, err error) {
	if err := db.Instance.First(&params).Error; err != nil {
		return params, err
	}
	return params, nil
}
