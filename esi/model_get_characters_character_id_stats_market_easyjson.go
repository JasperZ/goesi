// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package esi

import (
	json "encoding/json"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonE319ea96DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdStatsMarketList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdStatsMarketList, 0, 0)
			} else {
				*out = GetCharactersCharacterIdStatsMarketList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdStatsMarket
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE319ea96EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdStatsMarketList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdStatsMarketList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE319ea96EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdStatsMarketList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE319ea96EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdStatsMarketList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE319ea96DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdStatsMarketList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE319ea96DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonE319ea96DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdStatsMarket) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "accept_contracts_courier":
			out.AcceptContractsCourier = int64(in.Int64())
		case "accept_contracts_item_exchange":
			out.AcceptContractsItemExchange = int64(in.Int64())
		case "buy_orders_placed":
			out.BuyOrdersPlaced = int64(in.Int64())
		case "cancel_market_order":
			out.CancelMarketOrder = int64(in.Int64())
		case "create_contracts_auction":
			out.CreateContractsAuction = int64(in.Int64())
		case "create_contracts_courier":
			out.CreateContractsCourier = int64(in.Int64())
		case "create_contracts_item_exchange":
			out.CreateContractsItemExchange = int64(in.Int64())
		case "deliver_courier_contract":
			out.DeliverCourierContract = int64(in.Int64())
		case "isk_gained":
			out.IskGained = int64(in.Int64())
		case "isk_spent":
			out.IskSpent = int64(in.Int64())
		case "modify_market_order":
			out.ModifyMarketOrder = int64(in.Int64())
		case "search_contracts":
			out.SearchContracts = int64(in.Int64())
		case "sell_orders_placed":
			out.SellOrdersPlaced = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE319ea96EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdStatsMarket) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AcceptContractsCourier != 0 {
		const prefix string = ",\"accept_contracts_courier\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.AcceptContractsCourier))
	}
	if in.AcceptContractsItemExchange != 0 {
		const prefix string = ",\"accept_contracts_item_exchange\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AcceptContractsItemExchange))
	}
	if in.BuyOrdersPlaced != 0 {
		const prefix string = ",\"buy_orders_placed\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.BuyOrdersPlaced))
	}
	if in.CancelMarketOrder != 0 {
		const prefix string = ",\"cancel_market_order\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.CancelMarketOrder))
	}
	if in.CreateContractsAuction != 0 {
		const prefix string = ",\"create_contracts_auction\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.CreateContractsAuction))
	}
	if in.CreateContractsCourier != 0 {
		const prefix string = ",\"create_contracts_courier\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.CreateContractsCourier))
	}
	if in.CreateContractsItemExchange != 0 {
		const prefix string = ",\"create_contracts_item_exchange\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.CreateContractsItemExchange))
	}
	if in.DeliverCourierContract != 0 {
		const prefix string = ",\"deliver_courier_contract\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.DeliverCourierContract))
	}
	if in.IskGained != 0 {
		const prefix string = ",\"isk_gained\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.IskGained))
	}
	if in.IskSpent != 0 {
		const prefix string = ",\"isk_spent\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.IskSpent))
	}
	if in.ModifyMarketOrder != 0 {
		const prefix string = ",\"modify_market_order\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ModifyMarketOrder))
	}
	if in.SearchContracts != 0 {
		const prefix string = ",\"search_contracts\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.SearchContracts))
	}
	if in.SellOrdersPlaced != 0 {
		const prefix string = ",\"sell_orders_placed\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.SellOrdersPlaced))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdStatsMarket) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE319ea96EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdStatsMarket) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE319ea96EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdStatsMarket) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE319ea96DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdStatsMarket) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE319ea96DecodeGithubComAntihaxGoesiEsi1(l, v)
}
