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

func easyjson9e1087fdDecodeEventoolInternalPkgDomain(in *jlexer.Lexer, out *UserPublicInfo) {
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
		case "ID":
			out.Id = uint64(in.Uint64())
		case "name":
			out.Username = string(in.String())
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
func easyjson9e1087fdEncodeEventoolInternalPkgDomain(out *jwriter.Writer, in UserPublicInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserPublicInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeEventoolInternalPkgDomain(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserPublicInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeEventoolInternalPkgDomain(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserPublicInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeEventoolInternalPkgDomain(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserPublicInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeEventoolInternalPkgDomain(l, v)
}
func easyjson9e1087fdDecodeEventoolInternalPkgDomain1(in *jlexer.Lexer, out *UserNotificationToken) {
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
		case "token":
			out.Token = string(in.String())
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
func easyjson9e1087fdEncodeEventoolInternalPkgDomain1(out *jwriter.Writer, in UserNotificationToken) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"token\":"
		out.RawString(prefix[1:])
		out.String(string(in.Token))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserNotificationToken) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeEventoolInternalPkgDomain1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserNotificationToken) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeEventoolInternalPkgDomain1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserNotificationToken) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeEventoolInternalPkgDomain1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserNotificationToken) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeEventoolInternalPkgDomain1(l, v)
}
func easyjson9e1087fdDecodeEventoolInternalPkgDomain2(in *jlexer.Lexer, out *UserBasic) {
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
		case "email":
			out.Email = string(in.String())
		case "password":
			out.Password = string(in.String())
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
func easyjson9e1087fdEncodeEventoolInternalPkgDomain2(out *jwriter.Writer, in UserBasic) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix[1:])
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserBasic) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeEventoolInternalPkgDomain2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserBasic) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeEventoolInternalPkgDomain2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserBasic) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeEventoolInternalPkgDomain2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserBasic) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeEventoolInternalPkgDomain2(l, v)
}
func easyjson9e1087fdDecodeEventoolInternalPkgDomain3(in *jlexer.Lexer, out *User) {
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
		case "ID":
			out.Id = uint64(in.Uint64())
		case "username":
			out.Username = string(in.String())
		case "password":
			out.Password = string(in.String())
		case "repeatpassword":
			out.RepeatPassword = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "imgsrc":
			out.Imgsrc = string(in.String())
		case "phonenumber":
			out.PhoneNumber = string(in.String())
		case "age":
			out.Age = uint64(in.Uint64())
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
					var v1 string
					v1 = string(in.String())
					out.Categories = append(out.Categories, v1)
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
func easyjson9e1087fdEncodeEventoolInternalPkgDomain3(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	if in.Password != "" {
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	if in.RepeatPassword != "" {
		const prefix string = ",\"repeatpassword\":"
		out.RawString(prefix)
		out.String(string(in.RepeatPassword))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"imgsrc\":"
		out.RawString(prefix)
		out.String(string(in.Imgsrc))
	}
	{
		const prefix string = ",\"phonenumber\":"
		out.RawString(prefix)
		out.String(string(in.PhoneNumber))
	}
	{
		const prefix string = ",\"age\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Age))
	}
	{
		const prefix string = ",\"categories\":"
		out.RawString(prefix)
		if in.Categories == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Categories {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeEventoolInternalPkgDomain3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeEventoolInternalPkgDomain3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeEventoolInternalPkgDomain3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeEventoolInternalPkgDomain3(l, v)
}
func easyjson9e1087fdDecodeEventoolInternalPkgDomain4(in *jlexer.Lexer, out *UpdUser) {
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
		case "username":
			out.Username = string(in.String())
		case "phonenumber":
			out.PhoneNumber = string(in.String())
		case "imgsrc":
			out.Imgsrc = string(in.String())
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
func easyjson9e1087fdEncodeEventoolInternalPkgDomain4(out *jwriter.Writer, in UpdUser) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix[1:])
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"phonenumber\":"
		out.RawString(prefix)
		out.String(string(in.PhoneNumber))
	}
	{
		const prefix string = ",\"imgsrc\":"
		out.RawString(prefix)
		out.String(string(in.Imgsrc))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UpdUser) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeEventoolInternalPkgDomain4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpdUser) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeEventoolInternalPkgDomain4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UpdUser) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeEventoolInternalPkgDomain4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpdUser) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeEventoolInternalPkgDomain4(l, v)
}
