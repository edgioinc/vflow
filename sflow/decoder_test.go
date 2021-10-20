//: ----------------------------------------------------------------------------
//: Copyright (C) 2017 Verizon.  All Rights Reserved.
//: All Rights Reserved
//:
//: file:    decoder_test.go
//: details: TODO
//: author:  Mehrdad Arshad Rad
//: date:    02/01/2017
//:
//: Licensed under the Apache License, Version 2.0 (the "License");
//: you may not use this file except in compliance with the License.
//: You may obtain a copy of the License at
//:
//:     http://www.apache.org/licenses/LICENSE-2.0
//:
//: Unless required by applicable law or agreed to in writing, software
//: distributed under the License is distributed on an "AS IS" BASIS,
//: WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//: See the License for the specific language governing permissions and
//: limitations under the License.
//: ----------------------------------------------------------------------------

package sflow

import (
	"bytes"
	"testing"
)

var TestsFlowRawPacket = []byte{0x00, 0x00, 0x00, 0x05, 0x00, 0x00,
	0x00, 0x01, 0x18, 0x03, 0x40, 0x21, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x8d, 0x63, 0x16, 0x1c,
	0x54, 0x89, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xe8, 0xa6, 0x5c,
	0xc8, 0xeb, 0x00, 0x00, 0x03, 0x56, 0x00, 0x00, 0x10, 0x00, 0xcc, 0x8e, 0xc0, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x02, 0x31, 0x00, 0x00, 0x02, 0xc3, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x00, 0x90, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x05, 0xee, 0x00, 0x00,
	0x00, 0x04, 0x00, 0x00, 0x00, 0x80, 0x40, 0x55, 0x39, 0x41, 0x04, 0xb8, 0xae, 0x4b, 0xc8, 0x41,
	0x3a, 0xe2, 0x08, 0x00, 0x45, 0x00, 0x05, 0xdc, 0xfa, 0x5d, 0x40, 0x00, 0x3e, 0x06, 0x27, 0x76,
	0x98, 0xc3, 0x21, 0x28, 0x45, 0x2a, 0x16, 0x33, 0x01, 0xbb, 0xd4, 0xd2, 0x81, 0x2c, 0x72, 0x9d,
	0x00, 0x05, 0x6d, 0x6f, 0x50, 0x10, 0x01, 0x6b, 0x03, 0xd0, 0x00, 0x00, 0xbb, 0x6e, 0xa1, 0x32,
	0xf3, 0x60, 0xcf, 0x2c, 0x45, 0x8e, 0x53, 0x02, 0x02, 0x3d, 0xd5, 0xe9, 0xda, 0x9d, 0x59, 0x40,
	0x4f, 0xf8, 0x1a, 0x48, 0x0e, 0x90, 0x16, 0xa0, 0x0a, 0x42, 0x37, 0x20, 0x28, 0x78, 0x36, 0x9f,
	0xdf, 0x7d, 0x7f, 0x8b, 0x80, 0xa2, 0xf3, 0x67, 0x83, 0x41, 0xfd, 0x76, 0xed, 0xac, 0xd7, 0x5b,
	0xbd, 0xcb, 0x5f, 0x5f, 0x65, 0xe4, 0xdc, 0xe4, 0x00, 0xa3, 0x56, 0x22, 0xe8, 0x47, 0x31, 0xc0,
	0x42, 0x8f, 0x87, 0x89, 0xb0, 0x82, 0x00, 0x00, 0x03, 0xe9, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x03, 0xea, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x01, 0xce, 0x48, 0xd2, 0x46, 0x00, 0x00,
	0x00, 0x18, 0x00, 0x00, 0x00, 0x13, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xe8, 0xa6, 0x5c,
	0xc8, 0xec, 0x00, 0x00, 0x03, 0x56, 0x00, 0x00, 0x10, 0x00, 0xcc, 0x8e, 0xd0, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x02, 0x31, 0x00, 0x00, 0x02, 0xc3, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x00, 0x90, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x05, 0xb2, 0x00, 0x00,
	0x00, 0x04, 0x00, 0x00, 0x00, 0x80, 0xd4, 0x6d, 0x50, 0x7f, 0x8a, 0xc9, 0xae, 0x4b, 0xc8, 0x41,
	0x3a, 0xe2, 0x08, 0x00, 0x45, 0x00, 0x05, 0xa0, 0x6a, 0x89, 0x40, 0x00, 0x3e, 0x06, 0x5e, 0xdc,
	0x98, 0xc3, 0x0d, 0x59, 0xac, 0x3a, 0x1b, 0x9c, 0x01, 0xbb, 0xb9, 0xf9, 0x03, 0xfa, 0xad, 0xec,
	0xf3, 0x37, 0xe3, 0x60, 0x50, 0x10, 0x01, 0x28, 0x28, 0xda, 0x00, 0x00, 0xb3, 0x6f, 0xc1, 0x7e,
	0x8a, 0x37, 0x74, 0x95, 0xbc, 0xb9, 0x7c, 0xaa, 0x85, 0x35, 0xcd, 0x05, 0x3f, 0x3a, 0x27, 0xcf,
	0xa8, 0x7d, 0xb0, 0x46, 0x51, 0xfc, 0x5c, 0xb8, 0x83, 0x76, 0xcb, 0x85, 0x2a, 0xb6, 0x42, 0x85,
	0x86, 0xa2, 0x61, 0x57, 0x92, 0xf0, 0x71, 0xf6, 0xa2, 0xa3, 0xfc, 0x58, 0x93, 0x99, 0x88, 0x9f,
	0x56, 0x21, 0x88, 0x22, 0x89, 0x66, 0xe8, 0x7a, 0xb2, 0x2e, 0x98, 0xaf, 0x70, 0xd6, 0xc0, 0x6e,
	0xe4, 0xbd, 0xc5, 0x78, 0x96, 0x05, 0x00, 0x00, 0x03, 0xe9, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x03, 0xea, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x01, 0xce, 0x48, 0xd3, 0x16, 0x00, 0x00,
	0x00, 0x18, 0x00, 0x00, 0x00, 0x15, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xe8, 0xa6, 0x5c,
	0xc8, 0xed, 0x00, 0x00, 0x03, 0x56, 0x00, 0x00, 0x10, 0x00, 0xcc, 0x8e, 0xe0, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x02, 0x31, 0x00, 0x00, 0x02, 0xc3, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x00, 0x90, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x05, 0x9e, 0x00, 0x00,
	0x00, 0x04, 0x00, 0x00, 0x00, 0x80, 0xd4, 0x6d, 0x50, 0x7f, 0x8a, 0xc9, 0xae, 0x4b, 0xc8, 0x41,
	0x3a, 0xe2, 0x08, 0x00, 0x45, 0x02, 0x05, 0x8c, 0xd1, 0xce, 0x40, 0x00, 0x3e, 0x06, 0xe0, 0x45,
	0x98, 0xc3, 0x21, 0x84, 0xac, 0x3a, 0x1e, 0xd4, 0x01, 0xbb, 0x9f, 0xd8, 0xaa, 0x45, 0xdc, 0x86,
	0x6f, 0x4c, 0xfd, 0x41, 0x50, 0x10, 0x01, 0x26, 0x91, 0x45, 0x00, 0x00, 0xd9, 0x89, 0x5f, 0x11,
	0x8f, 0x1c, 0xdc, 0xda, 0x35, 0x98, 0xc4, 0x03, 0xa4, 0x7b, 0x56, 0x11, 0xd3, 0x3d, 0x25, 0xe7,
	0xf9, 0x19, 0x57, 0xd0, 0x44, 0xa2, 0x59, 0x3d, 0xc9, 0x90, 0xca, 0x7a, 0xa5, 0xbf, 0x00, 0x1e,
	0x98, 0x1c, 0x8c, 0x00, 0x4f, 0x5c, 0xf7, 0x89, 0x86, 0xfe, 0x88, 0x2e, 0x32, 0x03, 0x59, 0xbc,
	0x51, 0x06, 0x56, 0xd9, 0x38, 0xe5, 0xbe, 0x6b, 0x79, 0x8a, 0xdf, 0xf8, 0x34, 0x6b, 0x86, 0xc7,
	0xb2, 0x91, 0x4c, 0x11, 0x47, 0x50, 0x00, 0x00, 0x03, 0xe9, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x03, 0xea, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x01, 0xce, 0x48, 0xd3, 0x16, 0x00, 0x00,
	0x00, 0x18, 0x00, 0x00, 0x00, 0x15, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xac, 0xa6, 0x5c,
	0xc8, 0xee, 0x00, 0x00, 0x03, 0x56, 0x00, 0x00, 0x10, 0x00, 0xcc, 0x8e, 0xf0, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x02, 0xc3, 0x00, 0x00, 0x02, 0x31, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x00, 0x54, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x46, 0x00, 0x00,
	0x00, 0x04, 0x00, 0x00, 0x00, 0x42, 0xae, 0x4b, 0xc8, 0x41, 0x3a, 0xe2, 0x40, 0x55, 0x39, 0x41,
	0x04, 0xb8, 0x08, 0x00, 0x45, 0x00, 0x00, 0x34, 0xd8, 0xb1, 0x40, 0x00, 0x38, 0x06, 0x06, 0xc5,
	0x68, 0xdc, 0xc5, 0x06, 0x5d, 0xb8, 0xd7, 0xb2, 0x95, 0x98, 0x01, 0xbb, 0xad, 0x33, 0xd4, 0x9c,
	0xf7, 0x0d, 0xcd, 0xc0, 0x80, 0x10, 0x2c, 0xcc, 0x54, 0x73, 0x00, 0x00, 0x01, 0x01, 0x05, 0x0a,
	0xf7, 0x0d, 0xd9, 0x28, 0xf7, 0x0d, 0xef, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x03, 0xe9, 0x00, 0x00,
	0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x03, 0xea, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x01, 0x98, 0xc3,
	0x4d, 0x83, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
	0x00, 0xe8, 0xa6, 0x5c, 0xc8, 0xef, 0x00, 0x00, 0x03, 0x56, 0x00, 0x00, 0x10, 0x00, 0xcc, 0x8f,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x32, 0x00, 0x00, 0x02, 0xc3, 0x00, 0x00,
	0x00, 0x03, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x90, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
	0x05, 0x8a, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x80, 0xd4, 0x6d, 0x50, 0x7f, 0x8a, 0xc9,
	0xae, 0x4b, 0xc8, 0x41, 0x3a, 0xe2, 0x08, 0x00, 0x45, 0x00, 0x05, 0x78, 0xdf, 0xfd, 0x40, 0x00,
	0x3e, 0x06, 0xfe, 0x28, 0xc0, 0xe5, 0xd2, 0xb5, 0xac, 0x3a, 0x19, 0x84, 0x01, 0xbb, 0x9d, 0x0e,
	0xd5, 0xf4, 0x53, 0xb6, 0x01, 0xe7, 0xe1, 0xc4, 0x50, 0x18, 0x01, 0x6b, 0x8b, 0xc9, 0x00, 0x00,
	0xa7, 0xd3, 0xc5, 0x76, 0x7f, 0x47, 0x38, 0xaf, 0x34, 0xc7, 0x01, 0xcb, 0xcc, 0xec, 0xa3, 0xc4,
	0x73, 0xac, 0xa9, 0xa2, 0x83, 0x26, 0x09, 0x43, 0x98, 0x8d, 0x88, 0x88, 0x84, 0x71, 0x8a, 0x21,
	0x72, 0xe0, 0xd6, 0x09, 0xf4, 0x31, 0x31, 0x4f, 0x18, 0xb3, 0x81, 0x71, 0xc3, 0x91, 0x52, 0xa0,
	0x73, 0xed, 0x97, 0xde, 0xa2, 0x2d, 0xff, 0x27, 0xd4, 0xb7, 0x8c, 0x9b, 0x3b, 0xb3, 0x92, 0x5b,
	0xdc, 0x6e, 0x51, 0x97, 0xaf, 0xa9, 0xde, 0xec, 0xcb, 0x8a, 0x00, 0x00, 0x03, 0xe9, 0x00, 0x00,
	0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x03, 0xea, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x01, 0xce, 0x48,
	0xd3, 0x16, 0x00, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00, 0x15}

func TestSFHeaderDecode(t *testing.T) {
	filter := []uint32{DataCounterSample}
	reader := bytes.NewReader(TestsFlowRawPacket)
	d := NewSFDecoder(reader, filter)
	datagram, err := d.sfHeaderDecode()

	if err != nil {
		t.Error("unexpected error", err)
	}

	if datagram.Version != 5 {
		t.Error("expected version: 5, got",
			datagram.Version)
	}

	if datagram.SysUpTime != 370955401 {
		t.Error("expected SysUpTime: 370955401, got",
			datagram.SysUpTime)
	}

	if datagram.SamplesNo != 5 {
		t.Error("expected SamplesNo: 5, got",
			datagram.SamplesNo)
	}

	if datagram.SequenceNo != 36195 {
		t.Error("expected SequenceNo: 36195, got",
			datagram.SequenceNo)
	}

	if datagram.IPVersion != 1 {
		t.Error("expected IPVersion: 1, got",
			datagram.IPVersion)
	}

	if datagram.IPAddress.String() != "24.3.64.33" {
		t.Error("expected agent ip address: 24.3.64.33, got",
			datagram.IPAddress.String())
	}
}

func TestGetSampleInfo(t *testing.T) {
	filter := []uint32{DataCounterSample}
	reader := bytes.NewReader(TestsFlowRawPacket)
	// skip sflow header
	skip := make([]byte, 4*7)
	reader.Read(skip)

	d := NewSFDecoder(reader, filter)

	sizes := []uint32{232, 232, 232, 172, 232}

	for i := 0; i < 5; i++ {
		sfTypeFormat, sfDataLength, err := d.getSampleInfo()
		if err != nil {
			t.Error("unexpected error", err)
		}
		if sfTypeFormat != 1 {
			t.Error("expected type format# 1, got", sfTypeFormat)
		}
		if sfDataLength != sizes[i] {
			t.Error("expected data length: ", sizes[i], ", got", sfDataLength)
		}

		d.reader.Seek(int64(sfDataLength), 1)
	}
}

func TestSFDecode(t *testing.T) {
	filter := []uint32{DataCounterSample}
	reader := bytes.NewReader(TestsFlowRawPacket)
	d := NewSFDecoder(reader, filter)
	_, err := d.SFDecode()
	if err != nil {
		t.Error("unexpected error", err)
	}
}

func TestDecodeSampleHeader(t *testing.T) {
	filter := []uint32{DataCounterSample}
	reader := bytes.NewReader(TestsFlowRawPacket)

	d := NewSFDecoder(reader, filter)

	datagram, err := d.SFDecode()
	if err != nil {
		t.Error("unexpected error", err)
	}

	if len(datagram.Samples) != 5 {
		t.Error("expected samples## 5, got", len(datagram.Samples))
	}

	sample := datagram.Samples[0].(*FlowSample)

	if sample.SequenceNo != 0xa65cc8eb {
		t.Error("expected SequenceNo 0xa65cc8eb, got", sample.SequenceNo)
	}

	if sample.SourceIDType != 0 {
		t.Error("expected SourceIDType 0, got", sample.SourceIDType)
	}

	if sample.SamplingRate != 0x1000 {
		t.Error("expected SamplingRate 0x1000, got", sample.SamplingRate)
	}

	if sample.SamplePool != 0xcc8ec000 {
		t.Error("expected SamplePool 0xcc8ec000, got", sample.SamplePool)
	}

	if sample.Drops != 0 {
		t.Error("expected Drops 0, got", sample.Drops)
	}

	if sample.InputIdx != 0x231 {
		t.Error("expected InputIdx 0x231, got", sample.InputIdx)
	}

	if sample.OutputIdx != 0x2c3 {
		t.Error("expected Output 0x2c3, got", sample.OutputIdx)
	}

	if sample.RecordsNo != 0x3 {
		t.Error("expected RecordsNo 0x3, got", sample.RecordsNo)
	}

}

func BenchmarkSFDecode(b *testing.B) {
	filter := []uint32{DataCounterSample}
	for i := 0; i < b.N; i++ {
		reader := bytes.NewReader(TestsFlowRawPacket)
		d := NewSFDecoder(reader, filter)
		d.SFDecode()
	}
}
