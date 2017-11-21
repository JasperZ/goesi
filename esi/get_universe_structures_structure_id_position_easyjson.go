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

func easyjson1145c69dDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetUniverseStructuresStructureIdPositionList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseStructuresStructureIdPositionList, 0, 5)
			} else {
				*out = GetUniverseStructuresStructureIdPositionList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseStructuresStructureIdPosition
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
func easyjson1145c69dEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetUniverseStructuresStructureIdPositionList) {
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
func (v GetUniverseStructuresStructureIdPositionList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1145c69dEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseStructuresStructureIdPositionList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1145c69dEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseStructuresStructureIdPositionList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1145c69dDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseStructuresStructureIdPositionList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1145c69dDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson1145c69dDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetUniverseStructuresStructureIdPosition) {
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
		case "x":
			out.X = float32(in.Float32())
		case "y":
			out.Y = float32(in.Float32())
		case "z":
			out.Z = float32(in.Float32())
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
func easyjson1145c69dEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetUniverseStructuresStructureIdPosition) {
	out.RawByte('{')
	first := true
	_ = first
	if in.X != 0 {
		const prefix string = ",\"x\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.X))
	}
	if in.Y != 0 {
		const prefix string = ",\"y\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Y))
	}
	if in.Z != 0 {
		const prefix string = ",\"z\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Z))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseStructuresStructureIdPosition) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1145c69dEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseStructuresStructureIdPosition) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1145c69dEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseStructuresStructureIdPosition) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1145c69dDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseStructuresStructureIdPosition) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1145c69dDecodeGithubComAntihaxGoesiEsi1(l, v)
}
