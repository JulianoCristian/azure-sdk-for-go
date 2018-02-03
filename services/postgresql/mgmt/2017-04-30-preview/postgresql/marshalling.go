package postgresql

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/xml"
	"reflect"
	"time"
	"unsafe"
)

const (
	rfc3339Format = "2006-01-02T15:04:05.0000000Z07:00"
)

// used to convert times from UTC to GMT before sending across the wire
var gmt = time.FixedZone("GMT", 0)

// internal type used for marshalling time in RFC1123 format
type timeRFC1123 struct {
	time.Time
}

// MarshalText implements the encoding.TextMarshaler interface for timeRFC1123.
func (t timeRFC1123) MarshalText() ([]byte, error) {
	return []byte(t.Format(time.RFC1123)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for timeRFC1123.
func (t *timeRFC1123) UnmarshalText(data []byte) (err error) {
	t.Time, err = time.Parse(time.RFC1123, string(data))
	return
}

// internal type used for marshalling time in RFC3339 format
type timeRFC3339 struct {
	time.Time
}

// MarshalText implements the encoding.TextMarshaler interface for timeRFC3339.
func (t timeRFC3339) MarshalText() ([]byte, error) {
	return []byte(t.Format(rfc3339Format)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for timeRFC3339.
func (t *timeRFC3339) UnmarshalText(data []byte) (err error) {
	t.Time, err = time.Parse(rfc3339Format, string(data))
	return
}

// internal type used for marshalling
type serverPropertiesForRestore struct {
	StorageMB          *int64                 `json:"storageMB,omitempty"`
	Version            ServerVersionType      `json:"version,omitempty"`
	SslEnforcement     SslEnforcementEnumType `json:"sslEnforcement,omitempty"`
	CreateMode         CreateModeType         `json:"createMode,omitempty"`
	SourceServerID     string                 `json:"sourceServerId,omitempty"`
	RestorePointInTime timeRFC3339            `json:"restorePointInTime,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for ServerPropertiesForRestore.
func (spfr ServerPropertiesForRestore) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if reflect.TypeOf((*ServerPropertiesForRestore)(nil)).Elem().Size() != reflect.TypeOf((*serverPropertiesForRestore)(nil)).Elem().Size() {
		panic("size mismatch between ServerPropertiesForRestore and serverPropertiesForRestore")
	}
	spfr2 := (*serverPropertiesForRestore)(unsafe.Pointer(&spfr))
	return e.EncodeElement(*spfr2, start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for ServerPropertiesForRestore.
func (spfr *ServerPropertiesForRestore) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if reflect.TypeOf((*ServerPropertiesForRestore)(nil)).Elem().Size() != reflect.TypeOf((*serverPropertiesForRestore)(nil)).Elem().Size() {
		panic("size mismatch between ServerPropertiesForRestore and serverPropertiesForRestore")
	}
	spfr2 := (*serverPropertiesForRestore)(unsafe.Pointer(spfr))
	return d.DecodeElement(spfr2, &start)
}

// internal type used for marshalling
type logFileProperties struct {
	Name             *string      `json:"name,omitempty"`
	SizeInKB         *int64       `json:"sizeInKB,omitempty"`
	CreatedTime      *timeRFC3339 `json:"createdTime,omitempty"`
	LastModifiedTime *timeRFC3339 `json:"lastModifiedTime,omitempty"`
	Type             *string      `json:"type,omitempty"`
	URL              *string      `json:"url,omitempty"`
}

// MarshalXML implements the xml.Marshaler interface for LogFileProperties.
func (lfp LogFileProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if reflect.TypeOf((*LogFileProperties)(nil)).Elem().Size() != reflect.TypeOf((*logFileProperties)(nil)).Elem().Size() {
		panic("size mismatch between LogFileProperties and logFileProperties")
	}
	lfp2 := (*logFileProperties)(unsafe.Pointer(&lfp))
	return e.EncodeElement(*lfp2, start)
}

// UnmarshalXML implements the xml.Unmarshaler interface for LogFileProperties.
func (lfp *LogFileProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if reflect.TypeOf((*LogFileProperties)(nil)).Elem().Size() != reflect.TypeOf((*logFileProperties)(nil)).Elem().Size() {
		panic("size mismatch between LogFileProperties and logFileProperties")
	}
	lfp2 := (*logFileProperties)(unsafe.Pointer(lfp))
	return d.DecodeElement(lfp2, &start)
}
