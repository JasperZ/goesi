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

func easyjson2acd0704DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetUniversePlanetsPlanetIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniversePlanetsPlanetIdOkList, 0, 1)
			} else {
				*out = GetUniversePlanetsPlanetIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniversePlanetsPlanetIdOk
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
func easyjson2acd0704EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetUniversePlanetsPlanetIdOkList) {
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
func (v GetUniversePlanetsPlanetIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2acd0704EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniversePlanetsPlanetIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2acd0704EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniversePlanetsPlanetIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2acd0704DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniversePlanetsPlanetIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2acd0704DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson2acd0704DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetUniversePlanetsPlanetIdOk) {
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
		case "planet_id":
			out.PlanetId = int32(in.Int32())
		case "name":
			out.Name = string(in.String())
		case "type_id":
			out.TypeId = int32(in.Int32())
		case "position":
			(out.Position).UnmarshalEasyJSON(in)
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
func easyjson2acd0704EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetUniversePlanetsPlanetIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.PlanetId != 0 {
		const prefix string = ",\"planet_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.PlanetId))
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
	if true {
		const prefix string = ",\"position\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Position).MarshalEasyJSON(out)
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

// MarshalJSON supports json.Marshaler interface
func (v GetUniversePlanetsPlanetIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2acd0704EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniversePlanetsPlanetIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2acd0704EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniversePlanetsPlanetIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2acd0704DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniversePlanetsPlanetIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2acd0704DecodeGithubComAntihaxGoesiEsi1(l, v)
}
