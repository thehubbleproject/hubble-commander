package db

import (
	"math/big"

	"github.com/BOPR/bazooka"
	"github.com/BOPR/core"
)

var (
	// EmptyWitness for batches with no commiments
	EmptyWitness = [][32]byte{core.ZeroLeaf}
)

// GetLatestBatch fetches the latest batch recorded in the DB
func (db *DB) GetLatestBatch() (batch core.Batch, err error) {
	if err := db.Instance.Order("batch_id desc").First(&batch).Error; err != nil {
		return batch, err
	}
	return batch, nil
}

// GetBatchCount fetches the total number of batchs stored
func (db *DB) GetBatchCount() (int, error) {
	var count int
	if err := db.Instance.Table("batches").Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

// AddNewBatch adds a new batch to the DB
// It also internally attaches batchID and offset to commitments and stores them in a seperate table
func (db *DB) AddNewBatch(batch core.Batch, commitments []core.Commitment) (batchID uint64, err error) {
	// get batch count
	batchCount, err := db.GetBatchCount()
	if err != nil {
		return batchID, err
	}
	// this is because batch ID starts from 0
	batch.BatchID = uint64(batchCount)
	// add commitments to DB
	err = db.addCommitmentsWithBatchID(commitments, batch.BatchID)
	if err != nil {
		return
	}
	return batch.BatchID, db.Instance.Create(batch).Error
}

// GetBatchByIndex fetches the batch using batch ID
// NOTE: ID for first batch is 0
func (db *DB) GetBatchByIndex(index uint64) (batch core.Batch, err error) {
	if err := db.Instance.Scopes(QueryByBatchID(index)).Find(&batch).Error; err != nil {
		return batch, err
	}
	return batch, nil
}

// CommitBatch changes the status of batch to committed
// this function only runs for the operator
func (db *DB) CommitBatch(ID uint64) error {
	batch, err := db.GetBatchByIndex(ID)
	if err != nil {
		return err
	}

	if batch.Status == core.BATCH_COMMITTED {
		return core.ErrBatchAlreadyCommitted
	}

	batch.Status = core.BATCH_COMMITTED
	return db.Instance.Model(&batch).Scopes(QueryByBatchID(batch.BatchID)).Update(batch).Error
}

// attaches post batch submission information and adds commitments to DB
func (db *DB) addCommitmentsWithBatchID(commitments []core.Commitment, batchID uint64) error {
	if len(commitments) == 0 {
		return nil
	}
	for i, commitment := range commitments {
		commitment.AttachPostSubmissionInfo(batchID, uint64(i))
		if err := db.Instance.Create(&commitment).Error; err != nil {
			return err
		}
	}
	return nil
}

// GetCommitmentsForBatch getter for all commitments sent in a batch in order of offset
func (db *DB) GetCommitmentsForBatch(id uint64) (commitments []core.Commitment, err error) {
	if err = db.Instance.Model(&commitments).Scopes(QueryByBatchID(id)).Order("offset asc").Find(&commitments).Error; err != nil {
		return
	}
	return
}

// GetLastCommitmentMP creates the merkle proof for last commitment pushed on-chain
func (db *DB) GetLastCommitmentMP(id uint64) (commitmentMP bazooka.TypesCommitmentInclusionProof, err error) {
	// fetch all commitments stored for the batch_Id
	commitments, err := db.GetCommitmentsForBatch(id)
	if err != nil {
		return
	}

	// create commitment leaves
	var commitmentLeaves []core.ByteArray
	for _, commitment := range commitments {
		leaf, err := commitment.CommitmentData.Leaf()
		if err != nil {
			return commitmentMP, err
		}
		commitmentLeaves = append(commitmentLeaves, leaf)
	}

	// if there are no commitments, fetch the state root
	if len(commitments) == 0 {
		root, _ := db.GetStateRoot()
		rootHash, _ := core.HexToByteArray(root.Hash)
		commitmentMP = bazooka.TypesCommitmentInclusionProof{
			Commitment: *core.NewCommitmentData(rootHash, core.ZeroLeaf),
			Path:       big.NewInt(0),
			Witness:    EmptyWitness,
		}
		return commitmentMP, nil
	}

	lastCommitment := len(commitments)
	lastCommitmentIdx := lastCommitment - 1

	// create a merkle tree from all commitment leaves
	tree, err := core.NewTree(commitmentLeaves)
	if err != nil {
		return
	}

	// create a merkle proof for last commitment
	_, witnesses, err := tree.GetWitnessForLeaf(uint64(lastCommitmentIdx))
	if err != nil {
		return
	}

	// convert withness to sol witness
	var solWitness [][32]byte
	for _, witness := range witnesses {
		solWitness = append(solWitness, witness)
	}

	// create commitment MP
	commitmentMP = bazooka.TypesCommitmentInclusionProof{
		Commitment: commitments[lastCommitmentIdx].CommitmentData,
		Path:       big.NewInt(int64(lastCommitmentIdx)),
		Witness:    solWitness,
	}

	return commitmentMP, nil
}
