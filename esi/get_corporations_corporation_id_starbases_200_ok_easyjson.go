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

func easyjson7a6ec264DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCorporationsCorporationIdStarbases200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCorporationsCorporationIdStarbases200OkList, 0, 1)
			} else {
				*out = GetCorporationsCorporationIdStarbases200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCorporationsCorporationIdStarbases200Ok
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
func easyjson7a6ec264EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCorporationsCorporationIdStarbases200OkList) {
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
func (v GetCorporationsCorporationIdStarbases200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7a6ec264EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdStarbases200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7a6ec264EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdStarbases200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7a6ec264DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdStarbases200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7a6ec264DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson7a6ec264DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCorporationsCorporationIdStarbases200Ok) {
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
		case "starbase_id":
			out.StarbaseId = int64(in.Int64())
		case "type_id":
			out.TypeId = int32(in.Int32())
		case "system_id":
			out.SystemId = int32(in.Int32())
		case "moon_id":
			out.MoonId = int32(in.Int32())
		case "state":
			out.State = string(in.String())
		case "unanchor_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UnanchorAt).UnmarshalJSON(data))
			}
		case "reinforced_until":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ReinforcedUntil).UnmarshalJSON(data))
			}
		case "onlined_since":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.OnlinedSince).UnmarshalJSON(data))
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
func easyjson7a6ec264EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCorporationsCorporationIdStarbases200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.StarbaseId != 0 {
		const prefix string = ",\"starbase_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.StarbaseId))
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
	if in.MoonId != 0 {
		const prefix string = ",\"moon_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.MoonId))
	}
	if in.State != "" {
		const prefix string = ",\"state\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.State))
	}
	if true {
		const prefix string = ",\"unanchor_at\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.UnanchorAt).MarshalJSON())
	}
	if true {
		const prefix string = ",\"reinforced_until\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ReinforcedUntil).MarshalJSON())
	}
	if true {
		const prefix string = ",\"onlined_since\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.OnlinedSince).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCorporationsCorporationIdStarbases200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7a6ec264EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdStarbases200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7a6ec264EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdStarbases200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7a6ec264DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdStarbases200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7a6ec264DecodeGithubComAntihaxGoesiEsi1(l, v)
}
