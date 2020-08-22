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

func easyjson67942d1eDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *PostCharactersCharacterIdAssetsLocations200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostCharactersCharacterIdAssetsLocations200OkList, 0, 2)
			} else {
				*out = PostCharactersCharacterIdAssetsLocations200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 PostCharactersCharacterIdAssetsLocations200Ok
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
func easyjson67942d1eEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in PostCharactersCharacterIdAssetsLocations200OkList) {
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
func (v PostCharactersCharacterIdAssetsLocations200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson67942d1eEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostCharactersCharacterIdAssetsLocations200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson67942d1eEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostCharactersCharacterIdAssetsLocations200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson67942d1eDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostCharactersCharacterIdAssetsLocations200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson67942d1eDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson67942d1eDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *PostCharactersCharacterIdAssetsLocations200Ok) {
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
		case "item_id":
			out.ItemId = int64(in.Int64())
		case "position":
			(out.Position).UnmarshalEasyJSON(in)
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
func easyjson67942d1eEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in PostCharactersCharacterIdAssetsLocations200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ItemId != 0 {
		const prefix string = ",\"item_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.ItemId))
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostCharactersCharacterIdAssetsLocations200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson67942d1eEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostCharactersCharacterIdAssetsLocations200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson67942d1eEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostCharactersCharacterIdAssetsLocations200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson67942d1eDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostCharactersCharacterIdAssetsLocations200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson67942d1eDecodeGithubComAntihaxGoesiEsi1(l, v)
}
