package block

import (
	"errors"

	"github.com/dojimanetwork/go-polka-rpc/v5/registry"
	"github.com/dojimanetwork/go-polka-rpc/v5/types"
)

type SignedBlock struct {
	Block         Block               `json:"block"`
	Justification types.Justification `json:"justification"`
}

type Block struct {
	Header     types.Header `json:"header"`
	Extrinsics []string     `json:"extrinsics"`
}

// DecodeExtrinsics attempts to decode the extrinsics using the provided decoder.
func (s *SignedBlock) DecodeExtrinsics(decoder *registry.ExtrinsicDecoder) ([]*registry.DecodedExtrinsic, error) {
	if s == nil {
		return nil, errors.New("nil SignedBlock")
	}

	var decodedExtrinsics []*registry.DecodedExtrinsic

	for _, hexEncodedExtrinsic := range s.Block.Extrinsics {
		decodedExtrinsic, err := decoder.DecodeHex(hexEncodedExtrinsic)

		if err != nil {
			return nil, err
		}

		decodedExtrinsics = append(decodedExtrinsics, decodedExtrinsic)
	}

	return decodedExtrinsics, nil
}
