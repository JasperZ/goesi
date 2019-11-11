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

func easyjson79bc02b5DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetKillmailsKillmailIdKillmailHashOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetKillmailsKillmailIdKillmailHashOkList, 0, 1)
			} else {
				*out = GetKillmailsKillmailIdKillmailHashOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetKillmailsKillmailIdKillmailHashOk
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
func easyjson79bc02b5EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetKillmailsKillmailIdKillmailHashOkList) {
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
func (v GetKillmailsKillmailIdKillmailHashOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson79bc02b5EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetKillmailsKillmailIdKillmailHashOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson79bc02b5EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetKillmailsKillmailIdKillmailHashOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson79bc02b5DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetKillmailsKillmailIdKillmailHashOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson79bc02b5DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson79bc02b5DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetKillmailsKillmailIdKillmailHashOk) {
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
		case "attackers":
			if in.IsNull() {
				in.Skip()
				out.Attackers = nil
			} else {
				in.Delim('[')
				if out.Attackers == nil {
					if !in.IsDelim(']') {
						out.Attackers = make([]GetKillmailsKillmailIdKillmailHashAttacker, 0, 1)
					} else {
						out.Attackers = []GetKillmailsKillmailIdKillmailHashAttacker{}
					}
				} else {
					out.Attackers = (out.Attackers)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetKillmailsKillmailIdKillmailHashAttacker
					(v4).UnmarshalEasyJSON(in)
					out.Attackers = append(out.Attackers, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "killmail_id":
			out.KillmailId = int32(in.Int32())
		case "killmail_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.KillmailTime).UnmarshalJSON(data))
			}
		case "moon_id":
			out.MoonId = int32(in.Int32())
		case "solar_system_id":
			out.SolarSystemId = int32(in.Int32())
		case "victim":
			(out.Victim).UnmarshalEasyJSON(in)
		case "war_id":
			out.WarId = int32(in.Int32())
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
func easyjson79bc02b5EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetKillmailsKillmailIdKillmailHashOk) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Attackers) != 0 {
		const prefix string = ",\"attackers\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v5, v6 := range in.Attackers {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if in.KillmailId != 0 {
		const prefix string = ",\"killmail_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.KillmailId))
	}
	if true {
		const prefix string = ",\"killmail_time\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.KillmailTime).MarshalJSON())
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
	if in.SolarSystemId != 0 {
		const prefix string = ",\"solar_system_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.SolarSystemId))
	}
	if true {
		const prefix string = ",\"victim\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Victim).MarshalEasyJSON(out)
	}
	if in.WarId != 0 {
		const prefix string = ",\"war_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.WarId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetKillmailsKillmailIdKillmailHashOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson79bc02b5EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetKillmailsKillmailIdKillmailHashOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson79bc02b5EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetKillmailsKillmailIdKillmailHashOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson79bc02b5DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetKillmailsKillmailIdKillmailHashOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson79bc02b5DecodeGithubComAntihaxGoesiEsi1(l, v)
}