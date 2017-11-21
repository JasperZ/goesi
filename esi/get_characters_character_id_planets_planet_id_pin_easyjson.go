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

func easyjson4f4069f8DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdPinList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdPlanetsPlanetIdPinList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdPlanetsPlanetIdPinList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdPlanetsPlanetIdPin
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
func easyjson4f4069f8EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdPinList) {
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
func (v GetCharactersCharacterIdPlanetsPlanetIdPinList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4f4069f8EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdPinList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4f4069f8EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdPinList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4f4069f8DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdPinList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4f4069f8DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson4f4069f8DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdPin) {
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
		case "latitude":
			out.Latitude = float32(in.Float32())
		case "longitude":
			out.Longitude = float32(in.Float32())
		case "pin_id":
			out.PinId = int64(in.Int64())
		case "type_id":
			out.TypeId = int32(in.Int32())
		case "schematic_id":
			out.SchematicId = int32(in.Int32())
		case "extractor_details":
			easyjson4f4069f8DecodeGithubComAntihaxGoesiEsi2(in, &out.ExtractorDetails)
		case "factory_details":
			easyjson4f4069f8DecodeGithubComAntihaxGoesiEsi3(in, &out.FactoryDetails)
		case "contents":
			if in.IsNull() {
				in.Skip()
				out.Contents = nil
			} else {
				in.Delim('[')
				if out.Contents == nil {
					if !in.IsDelim(']') {
						out.Contents = make([]GetCharactersCharacterIdPlanetsPlanetIdContent, 0, 4)
					} else {
						out.Contents = []GetCharactersCharacterIdPlanetsPlanetIdContent{}
					}
				} else {
					out.Contents = (out.Contents)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetCharactersCharacterIdPlanetsPlanetIdContent
					easyjson4f4069f8DecodeGithubComAntihaxGoesiEsi4(in, &v4)
					out.Contents = append(out.Contents, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "install_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.InstallTime).UnmarshalJSON(data))
			}
		case "expiry_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ExpiryTime).UnmarshalJSON(data))
			}
		case "last_cycle_start":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.LastCycleStart).UnmarshalJSON(data))
			}
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
func easyjson4f4069f8EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdPin) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Latitude != 0 {
		const prefix string = ",\"latitude\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Latitude))
	}
	if in.Longitude != 0 {
		const prefix string = ",\"longitude\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Longitude))
	}
	if in.PinId != 0 {
		const prefix string = ",\"pin_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.PinId))
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
	if in.SchematicId != 0 {
		const prefix string = ",\"schematic_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.SchematicId))
	}
	if true {
		const prefix string = ",\"extractor_details\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson4f4069f8EncodeGithubComAntihaxGoesiEsi2(out, in.ExtractorDetails)
	}
	if true {
		const prefix string = ",\"factory_details\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson4f4069f8EncodeGithubComAntihaxGoesiEsi3(out, in.FactoryDetails)
	}
	if len(in.Contents) != 0 {
		const prefix string = ",\"contents\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.Contents {
				if v5 > 0 {
					out.RawByte(',')
				}
				easyjson4f4069f8EncodeGithubComAntihaxGoesiEsi4(out, v6)
			}
			out.RawByte(']')
		}
	}
	if true {
		const prefix string = ",\"install_time\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.InstallTime).MarshalJSON())
	}
	if true {
		const prefix string = ",\"expiry_time\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ExpiryTime).MarshalJSON())
	}
	if true {
		const prefix string = ",\"last_cycle_start\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.LastCycleStart).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdPin) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4f4069f8EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdPin) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4f4069f8EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdPin) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4f4069f8DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdPin) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4f4069f8DecodeGithubComAntihaxGoesiEsi1(l, v)
}
func easyjson4f4069f8DecodeGithubComAntihaxGoesiEsi4(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdContent) {
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
		case "type_id":
			out.TypeId = int32(in.Int32())
		case "amount":
			out.Amount = int64(in.Int64())
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
func easyjson4f4069f8EncodeGithubComAntihaxGoesiEsi4(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdContent) {
	out.RawByte('{')
	first := true
	_ = first
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
	if in.Amount != 0 {
		const prefix string = ",\"amount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Amount))
	}
	out.RawByte('}')
}
func easyjson4f4069f8DecodeGithubComAntihaxGoesiEsi3(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdFactoryDetails) {
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
		case "schematic_id":
			out.SchematicId = int32(in.Int32())
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
func easyjson4f4069f8EncodeGithubComAntihaxGoesiEsi3(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdFactoryDetails) {
	out.RawByte('{')
	first := true
	_ = first
	if in.SchematicId != 0 {
		const prefix string = ",\"schematic_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.SchematicId))
	}
	out.RawByte('}')
}
func easyjson4f4069f8DecodeGithubComAntihaxGoesiEsi2(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails) {
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
		case "heads":
			if in.IsNull() {
				in.Skip()
				out.Heads = nil
			} else {
				in.Delim('[')
				if out.Heads == nil {
					if !in.IsDelim(']') {
						out.Heads = make([]GetCharactersCharacterIdPlanetsPlanetIdHead, 0, 5)
					} else {
						out.Heads = []GetCharactersCharacterIdPlanetsPlanetIdHead{}
					}
				} else {
					out.Heads = (out.Heads)[:0]
				}
				for !in.IsDelim(']') {
					var v7 GetCharactersCharacterIdPlanetsPlanetIdHead
					easyjson4f4069f8DecodeGithubComAntihaxGoesiEsi5(in, &v7)
					out.Heads = append(out.Heads, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "product_type_id":
			out.ProductTypeId = int32(in.Int32())
		case "cycle_time":
			out.CycleTime = int32(in.Int32())
		case "head_radius":
			out.HeadRadius = float32(in.Float32())
		case "qty_per_cycle":
			out.QtyPerCycle = int32(in.Int32())
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
func easyjson4f4069f8EncodeGithubComAntihaxGoesiEsi2(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Heads) != 0 {
		const prefix string = ",\"heads\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v8, v9 := range in.Heads {
				if v8 > 0 {
					out.RawByte(',')
				}
				easyjson4f4069f8EncodeGithubComAntihaxGoesiEsi5(out, v9)
			}
			out.RawByte(']')
		}
	}
	if in.ProductTypeId != 0 {
		const prefix string = ",\"product_type_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.ProductTypeId))
	}
	if in.CycleTime != 0 {
		const prefix string = ",\"cycle_time\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.CycleTime))
	}
	if in.HeadRadius != 0 {
		const prefix string = ",\"head_radius\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.HeadRadius))
	}
	if in.QtyPerCycle != 0 {
		const prefix string = ",\"qty_per_cycle\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.QtyPerCycle))
	}
	out.RawByte('}')
}
func easyjson4f4069f8DecodeGithubComAntihaxGoesiEsi5(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdHead) {
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
		case "head_id":
			out.HeadId = int32(in.Int32())
		case "latitude":
			out.Latitude = float32(in.Float32())
		case "longitude":
			out.Longitude = float32(in.Float32())
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
func easyjson4f4069f8EncodeGithubComAntihaxGoesiEsi5(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdHead) {
	out.RawByte('{')
	first := true
	_ = first
	if in.HeadId != 0 {
		const prefix string = ",\"head_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.HeadId))
	}
	if in.Latitude != 0 {
		const prefix string = ",\"latitude\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Latitude))
	}
	if in.Longitude != 0 {
		const prefix string = ",\"longitude\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Longitude))
	}
	out.RawByte('}')
}
