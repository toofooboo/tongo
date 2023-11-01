package abi

// Code autogenerated. DO NOT EDIT.

import (
	"errors"

	"github.com/tonkeeper/tongo/boc"
	"github.com/tonkeeper/tongo/tlb"
)

func (j *JettonPayload) UnmarshalTLB(cell *boc.Cell, decoder *tlb.Decoder) error {
	if cell.BitsAvailableForRead() == 0 && cell.RefsAvailableForRead() == 0 {
		return nil
	}
	tempCell := cell.CopyRemaining()
	op64, err := tempCell.ReadUint(32)
	if errors.Is(err, boc.ErrNotEnoughBits) {
		j.SumType = UnknownJettonOp
		j.Value = cell.CopyRemaining()
		return nil
	}
	op := uint32(op64)
	j.OpCode = &op
	switch op {
	case TextCommentJettonOpCode: // 0x00000000
		var res TextCommentJettonPayload
		if err := tlb.Unmarshal(tempCell, &res); err == nil {
			j.SumType = TextCommentJettonOp
			j.Value = res
			return nil
		}
	case TegroJettonSwapJettonOpCode: // 0x01fb7a25
		var res TegroJettonSwapJettonPayload
		if err := tlb.Unmarshal(tempCell, &res); err == nil {
			j.SumType = TegroJettonSwapJettonOp
			j.Value = res
			return nil
		}
	case EncryptedTextCommentJettonOpCode: // 0x2167da4b
		var res EncryptedTextCommentJettonPayload
		if err := tlb.Unmarshal(tempCell, &res); err == nil {
			j.SumType = EncryptedTextCommentJettonOp
			j.Value = res
			return nil
		}
	case StonfiSwapJettonOpCode: // 0x25938561
		var res StonfiSwapJettonPayload
		if err := tlb.Unmarshal(tempCell, &res); err == nil {
			j.SumType = StonfiSwapJettonOp
			j.Value = res
			return nil
		}
	case TegroAddLiquidityJettonOpCode: // 0x287e167a
		var res TegroAddLiquidityJettonPayload
		if err := tlb.Unmarshal(tempCell, &res); err == nil {
			j.SumType = TegroAddLiquidityJettonOp
			j.Value = res
			return nil
		}
	case StonfiSwapOkRefJettonOpCode: // 0x45078540
		var res StonfiSwapOkRefJettonPayload
		if err := tlb.Unmarshal(tempCell, &res); err == nil {
			j.SumType = StonfiSwapOkRefJettonOp
			j.Value = res
			return nil
		}
	case StonfiSwapOkJettonOpCode: // 0xc64370e5
		var res StonfiSwapOkJettonPayload
		if err := tlb.Unmarshal(tempCell, &res); err == nil {
			j.SumType = StonfiSwapOkJettonOp
			j.Value = res
			return nil
		}
	case DedustSwapJettonOpCode: // 0xe3a0d482
		var res DedustSwapJettonPayload
		if err := tlb.Unmarshal(tempCell, &res); err == nil {
			j.SumType = DedustSwapJettonOp
			j.Value = res
			return nil
		}
	case StonfiProvideLiquidityJettonOpCode: // 0xfcf9e58f
		var res StonfiProvideLiquidityJettonPayload
		if err := tlb.Unmarshal(tempCell, &res); err == nil {
			j.SumType = StonfiProvideLiquidityJettonOp
			j.Value = res
			return nil
		}

	}
	j.SumType = UnknownJettonOp
	j.Value = cell.CopyRemaining()

	return nil
}

const (
	TextCommentJettonOp            JettonOpName = "TextComment"
	TegroJettonSwapJettonOp        JettonOpName = "TegroJettonSwap"
	EncryptedTextCommentJettonOp   JettonOpName = "EncryptedTextComment"
	StonfiSwapJettonOp             JettonOpName = "StonfiSwap"
	TegroAddLiquidityJettonOp      JettonOpName = "TegroAddLiquidity"
	StonfiSwapOkRefJettonOp        JettonOpName = "StonfiSwapOkRef"
	StonfiSwapOkJettonOp           JettonOpName = "StonfiSwapOk"
	DedustSwapJettonOp             JettonOpName = "DedustSwap"
	StonfiProvideLiquidityJettonOp JettonOpName = "StonfiProvideLiquidity"

	TextCommentJettonOpCode            JettonOpCode = 0x00000000
	TegroJettonSwapJettonOpCode        JettonOpCode = 0x01fb7a25
	EncryptedTextCommentJettonOpCode   JettonOpCode = 0x2167da4b
	StonfiSwapJettonOpCode             JettonOpCode = 0x25938561
	TegroAddLiquidityJettonOpCode      JettonOpCode = 0x287e167a
	StonfiSwapOkRefJettonOpCode        JettonOpCode = 0x45078540
	StonfiSwapOkJettonOpCode           JettonOpCode = 0xc64370e5
	DedustSwapJettonOpCode             JettonOpCode = 0xe3a0d482
	StonfiProvideLiquidityJettonOpCode JettonOpCode = 0xfcf9e58f
)

var KnownJettonTypes = map[string]any{
	TextCommentJettonOp:            TextCommentJettonPayload{},
	TegroJettonSwapJettonOp:        TegroJettonSwapJettonPayload{},
	EncryptedTextCommentJettonOp:   EncryptedTextCommentJettonPayload{},
	StonfiSwapJettonOp:             StonfiSwapJettonPayload{},
	TegroAddLiquidityJettonOp:      TegroAddLiquidityJettonPayload{},
	StonfiSwapOkRefJettonOp:        StonfiSwapOkRefJettonPayload{},
	StonfiSwapOkJettonOp:           StonfiSwapOkJettonPayload{},
	DedustSwapJettonOp:             DedustSwapJettonPayload{},
	StonfiProvideLiquidityJettonOp: StonfiProvideLiquidityJettonPayload{},
}
var jettonOpCodes = map[JettonOpName]JettonOpCode{
	TextCommentJettonOp:            TextCommentJettonOpCode,
	TegroJettonSwapJettonOp:        TegroJettonSwapJettonOpCode,
	EncryptedTextCommentJettonOp:   EncryptedTextCommentJettonOpCode,
	StonfiSwapJettonOp:             StonfiSwapJettonOpCode,
	TegroAddLiquidityJettonOp:      TegroAddLiquidityJettonOpCode,
	StonfiSwapOkRefJettonOp:        StonfiSwapOkRefJettonOpCode,
	StonfiSwapOkJettonOp:           StonfiSwapOkJettonOpCode,
	DedustSwapJettonOp:             DedustSwapJettonOpCode,
	StonfiProvideLiquidityJettonOp: StonfiProvideLiquidityJettonOpCode,
}

type TextCommentJettonPayload struct {
	Text tlb.Text
}

type TegroJettonSwapJettonPayload struct {
	Extract          bool
	MaxIn            tlb.VarUInteger16
	MinOut           tlb.VarUInteger16
	Destination      tlb.MsgAddress
	ErrorDestination tlb.MsgAddress
	Payload          *tlb.Any `tlb:"maybe^"`
}

type EncryptedTextCommentJettonPayload struct {
	CipherText tlb.Bytes
}

type StonfiSwapJettonPayload struct {
	TokenWallet     tlb.MsgAddress
	MinOut          tlb.VarUInteger16
	ToAddress       tlb.MsgAddress
	ReferralAddress *tlb.MsgAddress `tlb:"maybe"`
}

type TegroAddLiquidityJettonPayload struct {
	AmountA tlb.VarUInteger16
	AbountB tlb.VarUInteger16
}

type StonfiSwapOkRefJettonPayload struct{}

type StonfiSwapOkJettonPayload struct{}

type DedustSwapJettonPayload struct {
	Step       DedustSwapStep
	SwapParams DedustSwapParams `tlb:"^"`
}

type StonfiProvideLiquidityJettonPayload struct {
	TokenWallet tlb.MsgAddress
	MinLpOut    tlb.VarUInteger16
}
