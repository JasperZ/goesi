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

func easyjson9551bd61DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetMarketsStructuresStructureId200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetMarketsStructuresStructureId200OkList, 0, 1)
			} else {
				*out = GetMarketsStructuresStructureId200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetMarketsStructuresStructureId200Ok
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
func easyjson9551bd61EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetMarketsStructuresStructureId200OkList) {
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
func (v GetMarketsStructuresStructureId200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9551bd61EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetMarketsStructuresStructureId200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9551bd61EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetMarketsStructuresStructureId200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9551bd61DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetMarketsStructuresStructureId200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9551bd61DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson9551bd61DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetMarketsStructuresStructureId200Ok) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "order_id":
			out.OrderId = int64(in.Int64())
		case "type_id":
			out.TypeId = int32(in.Int32())
		case "location_id":
			out.LocationId = int64(in.Int64())
		case "volume_total":
			out.VolumeTotal = int32(in.Int32())
		case "volume_remain":
			out.VolumeRemain = int32(in.Int32())
		case "min_volume":
			out.MinVolume = int32(in.Int32())
		case "price":
			out.Price = float32(in.Float32())
		case "is_buy_order":
			out.IsBuyOrder = bool(in.Bool())
		case "duration":
			out.Duration = int32(in.Int32())
		case "issued":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Issued).UnmarshalJSON(data))
			}
		case "range":
			out.Range_ = string(in.String())
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
func easyjson9551bd61EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetMarketsStructuresStructureId200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.OrderId != 0 {
		const prefix string = ",\"order_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.OrderId))
	}
	if in.TypeId != 0 {
		const prefix string = ",\"type_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.TypeId))
	}
	if in.LocationId != 0 {
		const prefix string = ",\"location_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.LocationId))
	}
	if in.VolumeTotal != 0 {
		const prefix string = ",\"volume_total\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.VolumeTotal))
	}
	if in.VolumeRemain != 0 {
		const prefix string = ",\"volume_remain\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.VolumeRemain))
	}
	if in.MinVolume != 0 {
		const prefix string = ",\"min_volume\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.MinVolume))
	}
	if in.Price != 0 {
		const prefix string = ",\"price\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Price))
	}
	if in.IsBuyOrder {
		const prefix string = ",\"is_buy_order\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsBuyOrder))
	}
	if in.Duration != 0 {
		const prefix string = ",\"duration\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Duration))
	}
	if true {
		const prefix string = ",\"issued\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Issued).MarshalJSON())
	}
	if in.Range_ != "" {
		const prefix string = ",\"range\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Range_))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetMarketsStructuresStructureId200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9551bd61EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetMarketsStructuresStructureId200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9551bd61EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetMarketsStructuresStructureId200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9551bd61DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetMarketsStructuresStructureId200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9551bd61DecodeGithubComAntihaxGoesiEsi1(l, v)
}
