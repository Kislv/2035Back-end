// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package domain

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

func easyjsonF642ad3eDecodeEventoolInternalPkgDomain(in *jlexer.Lexer, out *EventListResponse) {
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
		case "eventlist":
			if in.IsNull() {
				in.Skip()
				out.EventList = nil
			} else {
				in.Delim('[')
				if out.EventList == nil {
					if !in.IsDelim(']') {
						out.EventList = make([]EventCreatingResponse, 0, 0)
					} else {
						out.EventList = []EventCreatingResponse{}
					}
				} else {
					out.EventList = (out.EventList)[:0]
				}
				for !in.IsDelim(']') {
					var v1 EventCreatingResponse
					(v1).UnmarshalEasyJSON(in)
					out.EventList = append(out.EventList, v1)
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
func easyjsonF642ad3eEncodeEventoolInternalPkgDomain(out *jwriter.Writer, in EventListResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"eventlist\":"
		out.RawString(prefix[1:])
		if in.EventList == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.EventList {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventListResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF642ad3eEncodeEventoolInternalPkgDomain(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventListResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF642ad3eEncodeEventoolInternalPkgDomain(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventListResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF642ad3eDecodeEventoolInternalPkgDomain(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventListResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF642ad3eDecodeEventoolInternalPkgDomain(l, v)
}
func easyjsonF642ad3eDecodeEventoolInternalPkgDomain1(in *jlexer.Lexer, out *EventCreatingResponse) {
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
		case "id":
			out.Id = uint64(in.Uint64())
		case "posterpath":
			out.PosterPath = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "rating":
			out.Rating = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "userid":
			out.UserId = string(in.String())
		case "longitude":
			out.Longitude = string(in.String())
		case "latitude":
			out.Latitude = string(in.String())
		case "currentmembersquantity":
			out.CurrentMembersQuantity = uint64(in.Uint64())
		case "maxmembersquantity":
			out.MaxMembersQuantity = uint64(in.Uint64())
		case "minmembersquantity":
			out.MinMembersQuantity = uint64(in.Uint64())
		case "creatingdate":
			out.CreatingDate = string(in.String())
		case "startdate":
			out.StartDate = string(in.String())
		case "enddate":
			out.EndDate = string(in.String())
		case "minage":
			out.MinAge = string(in.String())
		case "maxage":
			out.MaxAge = string(in.String())
		case "price":
			out.Price = string(in.String())
		case "categories":
			if in.IsNull() {
				in.Skip()
				out.Categories = nil
			} else {
				in.Delim('[')
				if out.Categories == nil {
					if !in.IsDelim(']') {
						out.Categories = make([]string, 0, 4)
					} else {
						out.Categories = []string{}
					}
				} else {
					out.Categories = (out.Categories)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Categories = append(out.Categories, v4)
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
func easyjsonF642ad3eEncodeEventoolInternalPkgDomain1(out *jwriter.Writer, in EventCreatingResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"posterpath\":"
		out.RawString(prefix)
		out.String(string(in.PosterPath))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		out.String(string(in.Rating))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"userid\":"
		out.RawString(prefix)
		out.String(string(in.UserId))
	}
	{
		const prefix string = ",\"longitude\":"
		out.RawString(prefix)
		out.String(string(in.Longitude))
	}
	{
		const prefix string = ",\"latitude\":"
		out.RawString(prefix)
		out.String(string(in.Latitude))
	}
	{
		const prefix string = ",\"currentmembersquantity\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.CurrentMembersQuantity))
	}
	{
		const prefix string = ",\"maxmembersquantity\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.MaxMembersQuantity))
	}
	{
		const prefix string = ",\"minmembersquantity\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.MinMembersQuantity))
	}
	{
		const prefix string = ",\"creatingdate\":"
		out.RawString(prefix)
		out.String(string(in.CreatingDate))
	}
	{
		const prefix string = ",\"startdate\":"
		out.RawString(prefix)
		out.String(string(in.StartDate))
	}
	{
		const prefix string = ",\"enddate\":"
		out.RawString(prefix)
		out.String(string(in.EndDate))
	}
	{
		const prefix string = ",\"minage\":"
		out.RawString(prefix)
		out.String(string(in.MinAge))
	}
	{
		const prefix string = ",\"maxage\":"
		out.RawString(prefix)
		out.String(string(in.MaxAge))
	}
	{
		const prefix string = ",\"price\":"
		out.RawString(prefix)
		out.String(string(in.Price))
	}
	{
		const prefix string = ",\"categories\":"
		out.RawString(prefix)
		if in.Categories == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Categories {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventCreatingResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF642ad3eEncodeEventoolInternalPkgDomain1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventCreatingResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF642ad3eEncodeEventoolInternalPkgDomain1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventCreatingResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF642ad3eDecodeEventoolInternalPkgDomain1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventCreatingResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF642ad3eDecodeEventoolInternalPkgDomain1(l, v)
}
func easyjsonF642ad3eDecodeEventoolInternalPkgDomain2(in *jlexer.Lexer, out *EventCreatingRequest) {
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
		case "posterpath":
			out.PosterPath = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "rating":
			out.Rating = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "userid":
			out.UserId = string(in.String())
		case "longitude":
			out.Longitude = string(in.String())
		case "latitude":
			out.Latitude = string(in.String())
		case "currentmembersquantity":
			out.CurrentMembersQuantity = uint32(in.Uint32())
		case "maxmembersquantity":
			out.MaxMembersQuantity = uint32(in.Uint32())
		case "minmembersquantity":
			out.MinMembersQuantity = uint32(in.Uint32())
		case "creatingdate":
			out.CreatingDate = string(in.String())
		case "startdate":
			out.StartDate = string(in.String())
		case "enddate":
			out.EndDate = string(in.String())
		case "minage":
			out.MinAge = string(in.String())
		case "maxage":
			out.MaxAge = string(in.String())
		case "price":
			out.Price = string(in.String())
		case "categories":
			if in.IsNull() {
				in.Skip()
				out.Categories = nil
			} else {
				in.Delim('[')
				if out.Categories == nil {
					if !in.IsDelim(']') {
						out.Categories = make([]string, 0, 4)
					} else {
						out.Categories = []string{}
					}
				} else {
					out.Categories = (out.Categories)[:0]
				}
				for !in.IsDelim(']') {
					var v7 string
					v7 = string(in.String())
					out.Categories = append(out.Categories, v7)
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
func easyjsonF642ad3eEncodeEventoolInternalPkgDomain2(out *jwriter.Writer, in EventCreatingRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"posterpath\":"
		out.RawString(prefix[1:])
		out.String(string(in.PosterPath))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		out.String(string(in.Rating))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"userid\":"
		out.RawString(prefix)
		out.String(string(in.UserId))
	}
	{
		const prefix string = ",\"longitude\":"
		out.RawString(prefix)
		out.String(string(in.Longitude))
	}
	{
		const prefix string = ",\"latitude\":"
		out.RawString(prefix)
		out.String(string(in.Latitude))
	}
	{
		const prefix string = ",\"currentmembersquantity\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.CurrentMembersQuantity))
	}
	{
		const prefix string = ",\"maxmembersquantity\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.MaxMembersQuantity))
	}
	{
		const prefix string = ",\"minmembersquantity\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.MinMembersQuantity))
	}
	{
		const prefix string = ",\"creatingdate\":"
		out.RawString(prefix)
		out.String(string(in.CreatingDate))
	}
	{
		const prefix string = ",\"startdate\":"
		out.RawString(prefix)
		out.String(string(in.StartDate))
	}
	{
		const prefix string = ",\"enddate\":"
		out.RawString(prefix)
		out.String(string(in.EndDate))
	}
	{
		const prefix string = ",\"minage\":"
		out.RawString(prefix)
		out.String(string(in.MinAge))
	}
	{
		const prefix string = ",\"maxage\":"
		out.RawString(prefix)
		out.String(string(in.MaxAge))
	}
	{
		const prefix string = ",\"price\":"
		out.RawString(prefix)
		out.String(string(in.Price))
	}
	{
		const prefix string = ",\"categories\":"
		out.RawString(prefix)
		if in.Categories == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Categories {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventCreatingRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF642ad3eEncodeEventoolInternalPkgDomain2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventCreatingRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF642ad3eEncodeEventoolInternalPkgDomain2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventCreatingRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF642ad3eDecodeEventoolInternalPkgDomain2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventCreatingRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF642ad3eDecodeEventoolInternalPkgDomain2(l, v)
}
