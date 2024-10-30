package history

import (
	"github.com/protolambda/zrnt/eth2/beacon/common"
	"github.com/protolambda/ztyp/codec"
	"github.com/protolambda/ztyp/tree"
)

const beaconBlockBodyProofLen = 8

type BeaconBlockBodyProof [beaconBlockBodyProofLen]common.Root

func (b *BeaconBlockBodyProof) Deserialize(dr *codec.DecodingReader) error {
	roots := b[:]
	return tree.ReadRoots(dr, &roots, beaconBlockBodyProofLen)
}

func (b *BeaconBlockBodyProof) Serialize(w *codec.EncodingWriter) error {
	return tree.WriteRoots(w, b[:])
}

func (b BeaconBlockBodyProof) ByteLength() (out uint64) {
	return beaconBlockBodyProofLen * 32
}

func (b BeaconBlockBodyProof) FixedLength() uint64 {
	return beaconBlockBodyProofLen * 32
}

func (b *BeaconBlockBodyProof) HashTreeRoot(hFn tree.HashFn) common.Root {
	return hFn.ComplexVectorHTR(func(i uint64) tree.HTR {
		if i < beaconBlockBodyProofLen {
			return &b[i]
		}
		return nil
	}, beaconBlockBodyProofLen)
}

const beaconBlockHeaderProofLen = 3

type BeaconBlockHeaderProof [beaconBlockHeaderProofLen]common.Root

func (b *BeaconBlockHeaderProof) Deserialize(dr *codec.DecodingReader) error {
	roots := b[:]
	return tree.ReadRoots(dr, &roots, beaconBlockHeaderProofLen)
}

func (b *BeaconBlockHeaderProof) Serialize(w *codec.EncodingWriter) error {
	return tree.WriteRoots(w, b[:])
}

func (b BeaconBlockHeaderProof) ByteLength() (out uint64) {
	return beaconBlockHeaderProofLen * 32
}

func (b BeaconBlockHeaderProof) FixedLength() uint64 {
	return beaconBlockHeaderProofLen * 32
}

func (b *BeaconBlockHeaderProof) HashTreeRoot(hFn tree.HashFn) common.Root {
	return hFn.ComplexVectorHTR(func(i uint64) tree.HTR {
		if i < beaconBlockHeaderProofLen {
			return &b[i]
		}
		return nil
	}, beaconBlockHeaderProofLen)
}

const historicalRootsProofLen = 14

type HistoricalRootsProof [historicalRootsProofLen]common.Root

func (b *HistoricalRootsProof) Deserialize(dr *codec.DecodingReader) error {
	roots := b[:]
	return tree.ReadRoots(dr, &roots, historicalRootsProofLen)
}

func (b *HistoricalRootsProof) Serialize(w *codec.EncodingWriter) error {
	return tree.WriteRoots(w, b[:])
}

func (b HistoricalRootsProof) ByteLength() (out uint64) {
	return historicalRootsProofLen * 32
}

func (b HistoricalRootsProof) FixedLength() uint64 {
	return historicalRootsProofLen * 32
}

func (b *HistoricalRootsProof) HashTreeRoot(hFn tree.HashFn) common.Root {
	return hFn.ComplexVectorHTR(func(i uint64) tree.HTR {
		if i < historicalRootsProofLen {
			return &b[i]
		}
		return nil
	}, historicalRootsProofLen)
}

type HistoricalRootsBlockProof struct {
	BeaconBlockBodyProof   BeaconBlockBodyProof   `yaml:"beacon_block_body_proof" json:"beacon_block_body_proof"`
	BeaconBlockBodyRoot    common.Root            `yaml:"beacon_block_body_root" json:"beacon_block_body_root"`
	BeaconBlockHeaderProof BeaconBlockHeaderProof `yaml:"beacon_block_header_proof" json:"beacon_block_header_proof"`
	BeaconBlockHeaderRoot  common.Root            `yaml:"beacon_block_header_root" json:"beacon_block_header_root"`
	HistoricalRootsProof   HistoricalRootsProof   `yaml:"historical_roots_proof" json:"historical_roots_proof"`
	Slot                   common.Slot            `yaml:"slot" json:"slot"`
}

func (h *HistoricalRootsBlockProof) Deserialize(dr *codec.DecodingReader) error {
	return dr.FixedLenContainer(
		&h.BeaconBlockBodyProof,
		&h.BeaconBlockBodyRoot,
		&h.BeaconBlockHeaderProof,
		&h.BeaconBlockHeaderProof,
		&h.HistoricalRootsProof,
		&h.Slot,
	)
}

func (h *HistoricalRootsBlockProof) Serialize(w *codec.EncodingWriter) error {
	return w.FixedLenContainer(
		&h.BeaconBlockBodyProof,
		&h.BeaconBlockBodyRoot,
		&h.BeaconBlockHeaderProof,
		&h.BeaconBlockHeaderProof,
		&h.HistoricalRootsProof,
		&h.Slot,
	)
}

func (h *HistoricalRootsBlockProof) ByteLength(spec *common.Spec) uint64 {
	return codec.ContainerLength(
		&h.BeaconBlockBodyProof,
		&h.BeaconBlockBodyRoot,
		&h.BeaconBlockHeaderProof,
		&h.BeaconBlockHeaderProof,
		&h.HistoricalRootsProof,
		&h.Slot,
	)
}

func (h *HistoricalRootsBlockProof) FixedLength(spec *common.Spec) uint64 {
	return codec.ContainerLength(
		&h.BeaconBlockBodyProof,
		&h.BeaconBlockBodyRoot,
		&h.BeaconBlockHeaderProof,
		&h.BeaconBlockHeaderProof,
		&h.HistoricalRootsProof,
		&h.Slot,
	)
}

func (h *HistoricalRootsBlockProof) HashTreeRoot(spec *common.Spec, hFn tree.HashFn) common.Root {
	return hFn.HashTreeRoot(
		&h.BeaconBlockBodyProof,
		&h.BeaconBlockBodyRoot,
		&h.BeaconBlockHeaderProof,
		&h.BeaconBlockHeaderProof,
		&h.HistoricalRootsProof,
		&h.Slot,
	)
}

type HistoricalRoots []common.Root

func (h *HistoricalRoots) Deserialize(spec *common.Spec, dr *codec.DecodingReader) error {
	return dr.List(func() codec.Deserializable {
		i := len(*h)
		*h = append(*h, common.Root{})
		return &(*h)[i]
	}, common.Root{}.ByteLength(), uint64(spec.HISTORICAL_ROOTS_LIMIT))
}

func (h HistoricalRoots) Serialize(spec *common.Spec, w *codec.EncodingWriter) error {
	return w.List(func(i uint64) codec.Serializable {
		return &h[i]
	}, common.Root{}.ByteLength(), uint64(spec.HISTORICAL_ROOTS_LIMIT))
}

func (h HistoricalRoots) ByteLength(spec *common.Spec) uint64 {
	return uint64(len(h)) * (common.Root{}.ByteLength())
}

func (h *HistoricalRoots) FixedLength(_ *common.Spec) uint64 {
	return 0
}

func (h HistoricalRoots) HashTreeRoot(spec *common.Spec, hFn tree.HashFn) common.Root {
	length := uint64(len(h))
	return hFn.ComplexListHTR(func(i uint64) tree.HTR {
		if i < length {
			return &h[i]
		}
		return nil
	}, length, uint64(spec.HISTORICAL_ROOTS_LIMIT))
}

// Proof for EL BlockHeader before TheMerge / Paris
const blockProofHistoricalHashesLength = 15
type BlockProofHistoricalHashesAccumulator [blockProofHistoricalHashesLength]common.Root

func (b *BlockProofHistoricalHashesAccumulator) Deserialize(dr *codec.DecodingReader) error {
	roots := b[:]
	return tree.ReadRoots(dr, &roots, blockProofHistoricalHashesLength)
}

func (b *BlockProofHistoricalHashesAccumulator) Serialize(w *codec.EncodingWriter) error {
	return tree.WriteRoots(w, b[:])
}

func (b BlockProofHistoricalHashesAccumulator) ByteLength() (out uint64) {
	return blockProofHistoricalHashesLength * 32
}

func (b BlockProofHistoricalHashesAccumulator) FixedLength() uint64 {
	return blockProofHistoricalHashesLength * 32
}

func (b *BlockProofHistoricalHashesAccumulator) HashTreeRoot(hFn tree.HashFn) common.Root {
	return hFn.ComplexVectorHTR(func(i uint64) tree.HTR {
		if i < blockProofHistoricalHashesLength {
			return &b[i]
		}
		return nil
	}, blockProofHistoricalHashesLength)
}

// Proof that EL block_hash is in BeaconBlock -> BeaconBlockBody -> ExecutionPayload
const executionBlockProofLength = 11
type ExecutionBlockProof [executionBlockProofLength]common.Root

func (b *ExecutionBlockProof) Deserialize(dr *codec.DecodingReader) error {
	roots := b[:]
	return tree.ReadRoots(dr, &roots, executionBlockProofLength)
}

func (b *ExecutionBlockProof) Serialize(w *codec.EncodingWriter) error {
	return tree.WriteRoots(w, b[:])
}

func (b ExecutionBlockProof) ByteLength() (out uint64) {
	return executionBlockProofLength * 32
}

func (b ExecutionBlockProof) FixedLength() uint64 {
	return executionBlockProofLength * 32
}

func (b *ExecutionBlockProof) HashTreeRoot(hFn tree.HashFn) common.Root {
	return hFn.ComplexVectorHTR(func(i uint64) tree.HTR {
		if i < executionBlockProofLength {
			return &b[i]
		}
		return nil
	}, executionBlockProofLength)
}

// Proof that BeaconBlock root is part of historical_roots and thus canonical
// From TheMerge until Capella -> Bellatrix fork.
const beaconBlockProofHistoricalRootsLength = 14
type BeaconBlockProofHistoricalRoots [beaconBlockProofHistoricalRootsLength]common.Root

func (b *BeaconBlockProofHistoricalRoots) Deserialize(dr *codec.DecodingReader) error {
	roots := b[:]
	return tree.ReadRoots(dr, &roots, beaconBlockProofHistoricalRootsLength)
}

func (b *BeaconBlockProofHistoricalRoots) Serialize(w *codec.EncodingWriter) error {
	return tree.WriteRoots(w, b[:])
}

func (b BeaconBlockProofHistoricalRoots) ByteLength() (out uint64) {
	return beaconBlockProofHistoricalRootsLength * 32
}

func (b BeaconBlockProofHistoricalRoots) FixedLength() uint64 {
	return beaconBlockProofHistoricalRootsLength * 32
}

func (b *BeaconBlockProofHistoricalRoots) HashTreeRoot(hFn tree.HashFn) common.Root {
	return hFn.ComplexVectorHTR(func(i uint64) tree.HTR {
		if i < beaconBlockProofHistoricalRootsLength {
			return &b[i]
		}
		return nil
	}, beaconBlockProofHistoricalRootsLength)
}

// Proof for EL BlockHeader from TheMerge until Capella
type BlockProofHistoricalRoots struct {
	BeaconBlockProof BeaconBlockProofHistoricalRoots // Proof that the BeaconBlock is part of the historical_roots and thus part of the canonical chain
	BeaconBlockRoot common.Root // hash_tree_root of BeaconBlock used to verify the proofs
	ExecutionBlockProof ExecutionBlockProof // Proof that EL BlockHash is part of the BeaconBlock
	Slot common.Slot // Slot of BeaconBlock, used to calculate the historical_roots index
}

func (h *BlockProofHistoricalRoots) Deserialize(dr *codec.DecodingReader) error {
	return dr.FixedLenContainer(
		&h.BeaconBlockProof,
		&h.BeaconBlockRoot,
		&h.ExecutionBlockProof,
		&h.Slot,
	)
}

func (h *BlockProofHistoricalRoots) Serialize(w *codec.EncodingWriter) error {
	return w.FixedLenContainer(
		&h.BeaconBlockProof,
		&h.BeaconBlockRoot,
		&h.ExecutionBlockProof,
		&h.Slot,
	)
}

func (h *BlockProofHistoricalRoots) ByteLength() uint64 {
	return codec.ContainerLength(
		&h.BeaconBlockProof,
		&h.BeaconBlockRoot,
		&h.ExecutionBlockProof,
		&h.Slot,
	)
}

func (h *BlockProofHistoricalRoots) FixedLength() uint64 {
	return codec.ContainerLength(
		&h.BeaconBlockProof,
		&h.BeaconBlockRoot,
		&h.ExecutionBlockProof,
		&h.Slot,
	)
}

func (h *BlockProofHistoricalRoots) HashTreeRoot(hFn tree.HashFn) common.Root {
	return hFn.HashTreeRoot(
		&h.BeaconBlockProof,
		&h.BeaconBlockRoot,
		&h.ExecutionBlockProof,
		&h.Slot,
	)
}

// Proof that EL block_hash is in BeaconBlock -> BeaconBlockBody -> ExecutionPayload
const executionBlockProofCapellaLimit = 12
type ExecutionBlockProofCapella []common.Root

func (r *ExecutionBlockProofCapella) Deserialize(dr *codec.DecodingReader) error {
	return dr.List(func() codec.Deserializable {
		i := len(*r)
		*r = append(*r, common.Root{})
		return &((*r)[i])
	}, 0, executionBlockProofCapellaLimit)
}

func (r ExecutionBlockProofCapella) Serialize(w *codec.EncodingWriter) error {
	return w.List(func(i uint64) codec.Serializable {
		return &r[i]
	}, 0, uint64(len(r)))
}

func (r ExecutionBlockProofCapella) ByteLength() (out uint64) {
	for _, v := range r {
		out += v.ByteLength() + codec.OFFSET_SIZE
	}
	return
}

func (r *ExecutionBlockProofCapella) FixedLength() uint64 {
	return 0
}

func (r ExecutionBlockProofCapella) HashTreeRoot(hFn tree.HashFn) common.Root {
	length := uint64(len(r))
	return hFn.ComplexListHTR(func(i uint64) tree.HTR {
		if i < length {
			return &r[i]
		}
		return nil
	}, length, executionBlockProofCapellaLimit)
}


// Proof that BeaconBlock root is part of historical_summaries and thus canonical
// For Capella and onwards
const beaconBlockProofHistoricalSummariesLength = 13
type BeaconBlockProofHistoricalSummaries [beaconBlockProofHistoricalSummariesLength]common.Root

func (b *BeaconBlockProofHistoricalSummaries) Deserialize(dr *codec.DecodingReader) error {
	roots := b[:]
	return tree.ReadRoots(dr, &roots, beaconBlockProofHistoricalSummariesLength)
}

func (b *BeaconBlockProofHistoricalSummaries) Serialize(w *codec.EncodingWriter) error {
	return tree.WriteRoots(w, b[:])
}

func (b BeaconBlockProofHistoricalSummaries) ByteLength() (out uint64) {
	return beaconBlockProofHistoricalSummariesLength * 32
}

func (b BeaconBlockProofHistoricalSummaries) FixedLength() uint64 {
	return beaconBlockProofHistoricalSummariesLength * 32
}

func (b *BeaconBlockProofHistoricalSummaries) HashTreeRoot(hFn tree.HashFn) common.Root {
	return hFn.ComplexVectorHTR(func(i uint64) tree.HTR {
		if i < beaconBlockProofHistoricalSummariesLength {
			return &b[i]
		}
		return nil
	}, beaconBlockProofHistoricalSummariesLength)
}

// Proof for EL BlockHeader for Capella and onwards
type BlockProofHistoricalSummaries struct {
	BeaconBlockProof BeaconBlockProofHistoricalSummaries // Proof that the BeaconBlock is part of the historical_roots and thus part of the canonical chain
	BeaconBlockRoot common.Root // hash_tree_root of BeaconBlock used to verify the proofs
	ExecutionBlockProof ExecutionBlockProofCapella // Proof that EL BlockHash is part of the BeaconBlock
	Slot common.Slot // Slot of BeaconBlock, used to calculate the historical_roots index
}

func (h *BlockProofHistoricalSummaries) Deserialize(dr *codec.DecodingReader) error {
	return dr.FixedLenContainer(
		&h.BeaconBlockProof,
		&h.BeaconBlockRoot,
		&h.ExecutionBlockProof,
		&h.Slot,
	)
}

func (h *BlockProofHistoricalSummaries) Serialize(w *codec.EncodingWriter) error {
	return w.FixedLenContainer(
		&h.BeaconBlockProof,
		&h.BeaconBlockRoot,
		&h.ExecutionBlockProof,
		&h.Slot,
	)
}

func (h *BlockProofHistoricalSummaries) ByteLength() uint64 {
	return codec.ContainerLength(
		&h.BeaconBlockProof,
		&h.BeaconBlockRoot,
		&h.ExecutionBlockProof,
		&h.Slot,
	)
}

func (h *BlockProofHistoricalSummaries) FixedLength() uint64 {
	return 0
}

func (h *BlockProofHistoricalSummaries) HashTreeRoot(hFn tree.HashFn) common.Root {
	return hFn.HashTreeRoot(
		&h.BeaconBlockProof,
		&h.BeaconBlockRoot,
		&h.ExecutionBlockProof,
		&h.Slot,
	)
}