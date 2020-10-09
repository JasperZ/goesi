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

func easyjson5fa7ff4fDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdStatsSocialList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdStatsSocialList, 0, 0)
			} else {
				*out = GetCharactersCharacterIdStatsSocialList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdStatsSocial
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
func easyjson5fa7ff4fEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdStatsSocialList) {
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
func (v GetCharactersCharacterIdStatsSocialList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5fa7ff4fEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdStatsSocialList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5fa7ff4fEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdStatsSocialList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5fa7ff4fDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdStatsSocialList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5fa7ff4fDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson5fa7ff4fDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdStatsSocial) {
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
		case "add_contact_bad":
			out.AddContactBad = int64(in.Int64())
		case "add_contact_good":
			out.AddContactGood = int64(in.Int64())
		case "add_contact_high":
			out.AddContactHigh = int64(in.Int64())
		case "add_contact_horrible":
			out.AddContactHorrible = int64(in.Int64())
		case "add_contact_neutral":
			out.AddContactNeutral = int64(in.Int64())
		case "add_note":
			out.AddNote = int64(in.Int64())
		case "added_as_contact_bad":
			out.AddedAsContactBad = int64(in.Int64())
		case "added_as_contact_good":
			out.AddedAsContactGood = int64(in.Int64())
		case "added_as_contact_high":
			out.AddedAsContactHigh = int64(in.Int64())
		case "added_as_contact_horrible":
			out.AddedAsContactHorrible = int64(in.Int64())
		case "added_as_contact_neutral":
			out.AddedAsContactNeutral = int64(in.Int64())
		case "calendar_event_created":
			out.CalendarEventCreated = int64(in.Int64())
		case "chat_messages_alliance":
			out.ChatMessagesAlliance = int64(in.Int64())
		case "chat_messages_constellation":
			out.ChatMessagesConstellation = int64(in.Int64())
		case "chat_messages_corporation":
			out.ChatMessagesCorporation = int64(in.Int64())
		case "chat_messages_fleet":
			out.ChatMessagesFleet = int64(in.Int64())
		case "chat_messages_region":
			out.ChatMessagesRegion = int64(in.Int64())
		case "chat_messages_solarsystem":
			out.ChatMessagesSolarsystem = int64(in.Int64())
		case "chat_messages_warfaction":
			out.ChatMessagesWarfaction = int64(in.Int64())
		case "chat_total_message_length":
			out.ChatTotalMessageLength = int64(in.Int64())
		case "direct_trades":
			out.DirectTrades = int64(in.Int64())
		case "fleet_broadcasts":
			out.FleetBroadcasts = int64(in.Int64())
		case "fleet_joins":
			out.FleetJoins = int64(in.Int64())
		case "mails_received":
			out.MailsReceived = int64(in.Int64())
		case "mails_sent":
			out.MailsSent = int64(in.Int64())
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
func easyjson5fa7ff4fEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdStatsSocial) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AddContactBad != 0 {
		const prefix string = ",\"add_contact_bad\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.AddContactBad))
	}
	if in.AddContactGood != 0 {
		const prefix string = ",\"add_contact_good\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AddContactGood))
	}
	if in.AddContactHigh != 0 {
		const prefix string = ",\"add_contact_high\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AddContactHigh))
	}
	if in.AddContactHorrible != 0 {
		const prefix string = ",\"add_contact_horrible\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AddContactHorrible))
	}
	if in.AddContactNeutral != 0 {
		const prefix string = ",\"add_contact_neutral\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AddContactNeutral))
	}
	if in.AddNote != 0 {
		const prefix string = ",\"add_note\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AddNote))
	}
	if in.AddedAsContactBad != 0 {
		const prefix string = ",\"added_as_contact_bad\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AddedAsContactBad))
	}
	if in.AddedAsContactGood != 0 {
		const prefix string = ",\"added_as_contact_good\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AddedAsContactGood))
	}
	if in.AddedAsContactHigh != 0 {
		const prefix string = ",\"added_as_contact_high\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AddedAsContactHigh))
	}
	if in.AddedAsContactHorrible != 0 {
		const prefix string = ",\"added_as_contact_horrible\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AddedAsContactHorrible))
	}
	if in.AddedAsContactNeutral != 0 {
		const prefix string = ",\"added_as_contact_neutral\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AddedAsContactNeutral))
	}
	if in.CalendarEventCreated != 0 {
		const prefix string = ",\"calendar_event_created\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.CalendarEventCreated))
	}
	if in.ChatMessagesAlliance != 0 {
		const prefix string = ",\"chat_messages_alliance\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ChatMessagesAlliance))
	}
	if in.ChatMessagesConstellation != 0 {
		const prefix string = ",\"chat_messages_constellation\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ChatMessagesConstellation))
	}
	if in.ChatMessagesCorporation != 0 {
		const prefix string = ",\"chat_messages_corporation\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ChatMessagesCorporation))
	}
	if in.ChatMessagesFleet != 0 {
		const prefix string = ",\"chat_messages_fleet\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ChatMessagesFleet))
	}
	if in.ChatMessagesRegion != 0 {
		const prefix string = ",\"chat_messages_region\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ChatMessagesRegion))
	}
	if in.ChatMessagesSolarsystem != 0 {
		const prefix string = ",\"chat_messages_solarsystem\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ChatMessagesSolarsystem))
	}
	if in.ChatMessagesWarfaction != 0 {
		const prefix string = ",\"chat_messages_warfaction\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ChatMessagesWarfaction))
	}
	if in.ChatTotalMessageLength != 0 {
		const prefix string = ",\"chat_total_message_length\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ChatTotalMessageLength))
	}
	if in.DirectTrades != 0 {
		const prefix string = ",\"direct_trades\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.DirectTrades))
	}
	if in.FleetBroadcasts != 0 {
		const prefix string = ",\"fleet_broadcasts\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.FleetBroadcasts))
	}
	if in.FleetJoins != 0 {
		const prefix string = ",\"fleet_joins\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.FleetJoins))
	}
	if in.MailsReceived != 0 {
		const prefix string = ",\"mails_received\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.MailsReceived))
	}
	if in.MailsSent != 0 {
		const prefix string = ",\"mails_sent\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.MailsSent))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdStatsSocial) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5fa7ff4fEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdStatsSocial) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5fa7ff4fEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdStatsSocial) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5fa7ff4fDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdStatsSocial) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5fa7ff4fDecodeGithubComAntihaxGoesiEsi1(l, v)
}