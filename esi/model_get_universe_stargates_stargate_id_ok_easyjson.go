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

func easyjson5e0a8a62DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetUniverseStargatesStargateIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseStargatesStargateIdOkList, 0, 1)
			} else {
				*out = GetUniverseStargatesStargateIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseStargatesStargateIdOk
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
func easyjson5e0a8a62EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetUniverseStargatesStargateIdOkList) {
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
func (v GetUniverseStargatesStargateIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5e0a8a62EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseStargatesStargateIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5e0a8a62EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseStargatesStargateIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5e0a8a62DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseStargatesStargateIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5e0a8a62DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson5e0a8a62DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetUniverseStargatesStargateIdOk) {
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
		case "destination":
			easyjson5e0a8a62DecodeGithubComAntihaxGoesiEsi2(in, &out.Destination)
		case "name":
			out.Name = string(in.String())
		case "position":
			easyjson5e0a8a62DecodeGithubComAntihaxGoesiEsi3(in, &out.Position)
		case "stargate_id":
			out.StargateId = int32(in.Int32())
		case "system_id":
			out.SystemId = int32(in.Int32())
		case "type_id":
			out.TypeId = int32(in.Int32())
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
func easyjson5e0a8a62EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetUniverseStargatesStargateIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		const prefix string = ",\"destination\":"
		first = false
		out.RawString(prefix[1:])
		easyjson5e0a8a62EncodeGithubComAntihaxGoesiEsi2(out, in.Destination)
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if true {
		const prefix string = ",\"position\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson5e0a8a62EncodeGithubComAntihaxGoesiEsi3(out, in.Position)
	}
	if in.StargateId != 0 {
		const prefix string = ",\"stargate_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.StargateId))
	}
	if in.SystemId != 0 {
		const prefix string = ",\"system_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.SystemId))
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseStargatesStargateIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5e0a8a62EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseStargatesStargateIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5e0a8a62EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseStargatesStargateIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5e0a8a62DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseStargatesStargateIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5e0a8a62DecodeGithubComAntihaxGoesiEsi1(l, v)
}
func easyjson5e0a8a62DecodeGithubComAntihaxGoesiEsi3(in *jlexer.Lexer, out *GetUniverseStargatesStargateIdPosition) {
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
		case "x":
			out.X = float64(in.Float64())
		case "y":
			out.Y = float64(in.Float64())
		case "z":
			out.Z = float64(in.Float64())
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
func easyjson5e0a8a62EncodeGithubComAntihaxGoesiEsi3(out *jwriter.Writer, in GetUniverseStargatesStargateIdPosition) {
	out.RawByte('{')
	first := true
	_ = first
	if in.X != 0 {
		const prefix string = ",\"x\":"
		first = false
		out.RawString(prefix[1:])
		out.Float64(float64(in.X))
	}
	if in.Y != 0 {
		const prefix string = ",\"y\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Y))
	}
	if in.Z != 0 {
		const prefix string = ",\"z\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Z))
	}
	out.RawByte('}')
}
func easyjson5e0a8a62DecodeGithubComAntihaxGoesiEsi2(in *jlexer.Lexer, out *GetUniverseStargatesStargateIdDestination) {
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
		case "stargate_id":
			out.StargateId = int32(in.Int32())
		case "system_id":
			out.SystemId = int32(in.Int32())
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
func easyjson5e0a8a62EncodeGithubComAntihaxGoesiEsi2(out *jwriter.Writer, in GetUniverseStargatesStargateIdDestination) {
	out.RawByte('{')
	first := true
	_ = first
	if in.StargateId != 0 {
		const prefix string = ",\"stargate_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.StargateId))
	}
	if in.SystemId != 0 {
		const prefix string = ",\"system_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.SystemId))
	}
	out.RawByte('}')
}