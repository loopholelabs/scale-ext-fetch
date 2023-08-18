// Code generated by scale-signature-go v0.1.0, DO NOT EDIT.
// schema: HttpFetch:alpha
// output: types

package extfetch

import (
	"errors"
	"github.com/loopholelabs/polyglot"
)

var (
	NilDecode   = errors.New("cannot decode into a nil root struct")
	InvalidEnum = errors.New("invalid enum value")
)

type HttpConfig struct {
	Timeout int32
}

func NewHttpConfig() *HttpConfig {
	return &HttpConfig{

		Timeout: 60,
	}
}

func (x *HttpConfig) Encode(b *polyglot.Buffer) {
	e := polyglot.Encoder(b)
	if x == nil {
		e.Nil()
	} else {

		e.Int32(x.Timeout)

	}
}

func DecodeHttpConfig(x *HttpConfig, b []byte) (*HttpConfig, error) {
	d := polyglot.GetDecoder(b)
	defer d.Return()
	return _decodeHttpConfig(x, d)
}

func _decodeHttpConfig(x *HttpConfig, d *polyglot.Decoder) (*HttpConfig, error) {
	if d.Nil() {
		return nil, nil
	}

	err, _ := d.Error()
	if err != nil {
		return nil, err
	}

	if x == nil {
		x = NewHttpConfig()
	}

	x.Timeout, err = d.Int32()
	if err != nil {
		return nil, err
	}

	return x, nil
}

type HttpResponse struct {
	Headers map[string]StringList

	StatusCode int32

	Body []byte
}

func NewHttpResponse() *HttpResponse {
	return &HttpResponse{

		Headers: make(map[string]StringList),

		StatusCode: 0,

		Body: make([]byte, 0, 0),
	}
}

func (x *HttpResponse) Encode(b *polyglot.Buffer) {
	e := polyglot.Encoder(b)
	if x == nil {
		e.Nil()
	} else {

		e.Map(uint32(len(x.Headers)), polyglot.StringKind, polyglot.AnyKind)
		for k, v := range x.Headers {
			e.String(k)
			v.Encode(b)
		}

		e.Int32(x.StatusCode)

		e.Bytes(x.Body)

	}
}

func DecodeHttpResponse(x *HttpResponse, b []byte) (*HttpResponse, error) {
	d := polyglot.GetDecoder(b)
	defer d.Return()
	return _decodeHttpResponse(x, d)
}

func _decodeHttpResponse(x *HttpResponse, d *polyglot.Decoder) (*HttpResponse, error) {
	if d.Nil() {
		return nil, nil
	}

	err, _ := d.Error()
	if err != nil {
		return nil, err
	}

	if x == nil {
		x = NewHttpResponse()
	}

	mapSizeHeaders, err := d.Map(polyglot.StringKind, polyglot.AnyKind)
	if err != nil {
		return nil, err
	}

	if uint32(len(x.Headers)) != mapSizeHeaders {
		x.Headers = make(map[string]StringList, mapSizeHeaders)
	}

	for i := uint32(0); i < mapSizeHeaders; i++ {
		k, err := d.String()
		if err != nil {
			return nil, err
		}
		v, err := _decodeStringList(nil, d)
		if err != nil {
			return nil, err
		}
		x.Headers[k] = *v
	}

	x.StatusCode, err = d.Int32()
	if err != nil {
		return nil, err
	}

	x.Body, err = d.Bytes(nil)
	if err != nil {
		return nil, err
	}

	return x, nil
}

type StringList struct {
	Values []string
}

func NewStringList() *StringList {
	return &StringList{

		Values: make([]string, 0, 0),
	}
}

func (x *StringList) Encode(b *polyglot.Buffer) {
	e := polyglot.Encoder(b)
	if x == nil {
		e.Nil()
	} else {

		e.Slice(uint32(len(x.Values)), polyglot.StringKind)
		for _, a := range x.Values {
			e.String(a)
		}

	}
}

func DecodeStringList(x *StringList, b []byte) (*StringList, error) {
	d := polyglot.GetDecoder(b)
	defer d.Return()
	return _decodeStringList(x, d)
}

func _decodeStringList(x *StringList, d *polyglot.Decoder) (*StringList, error) {
	if d.Nil() {
		return nil, nil
	}

	err, _ := d.Error()
	if err != nil {
		return nil, err
	}

	if x == nil {
		x = NewStringList()
	}

	sliceSizeValues, err := d.Slice(polyglot.StringKind)
	if err != nil {
		return nil, err
	}

	if uint32(len(x.Values)) != sliceSizeValues {
		x.Values = make([]string, sliceSizeValues)
	}

	for i := uint32(0); i < sliceSizeValues; i++ {
		x.Values[i], err = d.String()
		if err != nil {
			return nil, err
		}
	}

	return x, nil
}

type ConnectionDetails struct {
	Url string
}

func NewConnectionDetails() *ConnectionDetails {
	return &ConnectionDetails{

		Url: "https://google.com",
	}
}

func (x *ConnectionDetails) Encode(b *polyglot.Buffer) {
	e := polyglot.Encoder(b)
	if x == nil {
		e.Nil()
	} else {

		e.String(x.Url)

	}
}

func DecodeConnectionDetails(x *ConnectionDetails, b []byte) (*ConnectionDetails, error) {
	d := polyglot.GetDecoder(b)
	defer d.Return()
	return _decodeConnectionDetails(x, d)
}

func _decodeConnectionDetails(x *ConnectionDetails, d *polyglot.Decoder) (*ConnectionDetails, error) {
	if d.Nil() {
		return nil, nil
	}

	err, _ := d.Error()
	if err != nil {
		return nil, err
	}

	if x == nil {
		x = NewConnectionDetails()
	}

	x.Url, err = d.String()
	if err != nil {
		return nil, err
	}

	return x, nil
}