/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/authorizations/v1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalExportControlReviewRequest writes a value of the 'export_control_review_request' type to the given writer.
func MarshalExportControlReviewRequest(object *ExportControlReviewRequest, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeExportControlReviewRequest(object, stream)
	stream.Flush()
	return stream.Error
}

// writeExportControlReviewRequest writes a value of the 'export_control_review_request' type to the given stream.
func writeExportControlReviewRequest(object *ExportControlReviewRequest, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	if object.accountUsername != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("account_username")
		stream.WriteString(*object.accountUsername)
		count++
	}
	stream.WriteObjectEnd()
}

// UnmarshalExportControlReviewRequest reads a value of the 'export_control_review_request' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalExportControlReviewRequest(source interface{}) (object *ExportControlReviewRequest, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readExportControlReviewRequest(iterator)
	err = iterator.Error
	return
}

// readExportControlReviewRequest reads a value of the 'export_control_review_request' type from the given iterator.
func readExportControlReviewRequest(iterator *jsoniter.Iterator) *ExportControlReviewRequest {
	object := &ExportControlReviewRequest{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "account_username":
			value := iterator.ReadString()
			object.accountUsername = &value
		default:
			iterator.ReadAny()
		}
	}
	return object
}
