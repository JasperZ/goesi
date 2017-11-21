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

func easyjson6213c735DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdClonesJumpCloneList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdClonesJumpCloneList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdClonesJumpCloneList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdClonesJumpClone
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
func easyjson6213c735EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdClonesJumpCloneList) {
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
func (v GetCharactersCharacterIdClonesJumpCloneList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6213c735EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdClonesJumpCloneList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6213c735EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdClonesJumpCloneList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6213c735DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdClonesJumpCloneList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6213c735DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson6213c735DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdClonesJumpClone) {
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
		case "location_id":
			out.LocationId = int64(in.Int64())
		case "location_type":
			out.LocationType = string(in.String())
		case "implants":
			if in.IsNull() {
				in.Skip()
				out.Implants = nil
			} else {
				in.Delim('[')
				if out.Implants == nil {
					if !in.IsDelim(']') {
						out.Implants = make([]int32, 0, 16)
					} else {
						out.Implants = []int32{}
					}
				} else {
					out.Implants = (out.Implants)[:0]
				}
				for !in.IsDelim(']') {
					var v4 int32
					v4 = int32(in.Int32())
					out.Implants = append(out.Implants, v4)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson6213c735EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdClonesJumpClone) {
	out.RawByte('{')
	first := true
	_ = first
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
	if in.LocationType != "" {
		const prefix string = ",\"location_type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LocationType))
	}
	if len(in.Implants) != 0 {
		const prefix string = ",\"implants\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.Implants {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdClonesJumpClone) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6213c735EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdClonesJumpClone) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6213c735EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdClonesJumpClone) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6213c735DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdClonesJumpClone) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6213c735DecodeGithubComAntihaxGoesiEsi1(l, v)
}
