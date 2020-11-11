package serialization

import (
	"github.com/kaspanet/kaspad/domain/consensus/model"
	"github.com/kaspanet/kaspad/domain/consensus/model/externalapi"
)

// BlockGHOSTDAGDataToDBBlockGHOSTDAGData converts BlockGHOSTDAGData to DbBlockGhostdagData
func BlockGHOSTDAGDataToDBBlockGHOSTDAGData(blockGHOSTDAGData *model.BlockGHOSTDAGData) *DbBlockGhostdagData {
	var selectedParent *DbHash
	if blockGHOSTDAGData.SelectedParent != nil {
		selectedParent = DomainHashToDbHash(blockGHOSTDAGData.SelectedParent)
	}

	return &DbBlockGhostdagData{
		BlueScore:          blockGHOSTDAGData.BlueScore,
		SelectedParent:     selectedParent,
		MergeSetBlues:      DomainHashesToDbHashes(blockGHOSTDAGData.MergeSetBlues),
		MergeSetReds:       DomainHashesToDbHashes(blockGHOSTDAGData.MergeSetReds),
		BluesAnticoneSizes: bluesAnticoneSizesToDBBluesAnticoneSizes(blockGHOSTDAGData.BluesAnticoneSizes),
	}
}

// DBBlockGHOSTDAGDataToBlockGHOSTDAGData converts DbBlockGhostdagData to BlockGHOSTDAGData
func DBBlockGHOSTDAGDataToBlockGHOSTDAGData(dbBlockGHOSTDAGData *DbBlockGhostdagData) (*model.BlockGHOSTDAGData, error) {
	var selectedParent *externalapi.DomainHash
	if dbBlockGHOSTDAGData.SelectedParent != nil {
		var err error
		selectedParent, err = DbHashToDomainHash(dbBlockGHOSTDAGData.SelectedParent)
		if err != nil {
			return nil, err
		}
	}

	mergetSetBlues, err := DbHashesToDomainHashes(dbBlockGHOSTDAGData.MergeSetBlues)
	if err != nil {
		return nil, err
	}

	mergetSetReds, err := DbHashesToDomainHashes(dbBlockGHOSTDAGData.MergeSetReds)
	if err != nil {
		return nil, err
	}

	bluesAnticoneSizes, err := dbBluesAnticoneSizesToBluesAnticoneSizes(dbBlockGHOSTDAGData.BluesAnticoneSizes)
	if err != nil {
		return nil, err
	}

	return &model.BlockGHOSTDAGData{
		BlueScore:          dbBlockGHOSTDAGData.BlueScore,
		SelectedParent:     selectedParent,
		MergeSetBlues:      mergetSetBlues,
		MergeSetReds:       mergetSetReds,
		BluesAnticoneSizes: bluesAnticoneSizes,
	}, nil
}