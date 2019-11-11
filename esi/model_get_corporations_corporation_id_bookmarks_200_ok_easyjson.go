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

func easyjson3851f0d1DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCorporationsCorporationIdBookmarks200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCorporationsCorporationIdBookmarks200OkList, 0, 1)
			} else {
				*out = GetCorporationsCorporationIdBookmarks200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCorporationsCorporationIdBookmarks200Ok
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
func easyjson3851f0d1EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCorporationsCorporationIdBookmarks200OkList) {
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
func (v GetCorporationsCorporationIdBookmarks200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3851f0d1EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdBookmarks200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3851f0d1EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdBookmarks200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3851f0d1DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdBookmarks200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3851f0d1DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson3851f0d1DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCorporationsCorporationIdBookmarks200Ok) {
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
		case "bookmark_id":
			out.BookmarkId = int32(in.Int32())
		case "coordinates":
			easyjson3851f0d1DecodeGithubComAntihaxGoesiEsi2(in, &out.Coordinates)
		case "created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Created).UnmarshalJSON(data))
			}
		case "creator_id":
			out.CreatorId = int32(in.Int32())
		case "folder_id":
			out.FolderId = int32(in.Int32())
		case "item":
			(out.Item).UnmarshalEasyJSON(in)
		case "label":
			out.Label = string(in.String())
		case "location_id":
			out.LocationId = int32(in.Int32())
		case "notes":
			out.Notes = string(in.String())
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
func easyjson3851f0d1EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCorporationsCorporationIdBookmarks200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.BookmarkId != 0 {
		const prefix string = ",\"bookmark_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.BookmarkId))
	}
	if true {
		const prefix string = ",\"coordinates\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson3851f0d1EncodeGithubComAntihaxGoesiEsi2(out, in.Coordinates)
	}
	if true {
		const prefix string = ",\"created\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Created).MarshalJSON())
	}
	if in.CreatorId != 0 {
		const prefix string = ",\"creator_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.CreatorId))
	}
	if in.FolderId != 0 {
		const prefix string = ",\"folder_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.FolderId))
	}
	if true {
		const prefix string = ",\"item\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Item).MarshalEasyJSON(out)
	}
	if in.Label != "" {
		const prefix string = ",\"label\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Label))
	}
	if in.LocationId != 0 {
		const prefix string = ",\"location_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.LocationId))
	}
	if in.Notes != "" {
		const prefix string = ",\"notes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Notes))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCorporationsCorporationIdBookmarks200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3851f0d1EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdBookmarks200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3851f0d1EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdBookmarks200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3851f0d1DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdBookmarks200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3851f0d1DecodeGithubComAntihaxGoesiEsi1(l, v)
}
func easyjson3851f0d1DecodeGithubComAntihaxGoesiEsi2(in *jlexer.Lexer, out *GetCorporationsCorporationIdBookmarksCoordinates) {
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
func easyjson3851f0d1EncodeGithubComAntihaxGoesiEsi2(out *jwriter.Writer, in GetCorporationsCorporationIdBookmarksCoordinates) {
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