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

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalNetwork writes a value of the 'network' type to the given writer.
func MarshalNetwork(object *Network, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeNetwork(object, stream)
	stream.Flush()
	return stream.Error
}

// writeNetwork writes a value of the 'network' type to the given stream.
func writeNetwork(object *Network, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	if object.hostPrefix != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("host_prefix")
		stream.WriteInt(*object.hostPrefix)
		count++
	}
	if object.machineCIDR != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("machine_cidr")
		stream.WriteString(*object.machineCIDR)
		count++
	}
	if object.podCIDR != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("pod_cidr")
		stream.WriteString(*object.podCIDR)
		count++
	}
	if object.serviceCIDR != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("service_cidr")
		stream.WriteString(*object.serviceCIDR)
		count++
	}
	stream.WriteObjectEnd()
}

// UnmarshalNetwork reads a value of the 'network' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalNetwork(source interface{}) (object *Network, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readNetwork(iterator)
	err = iterator.Error
	return
}

// readNetwork reads a value of the 'network' type from the given iterator.
func readNetwork(iterator *jsoniter.Iterator) *Network {
	object := &Network{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "host_prefix":
			value := iterator.ReadInt()
			object.hostPrefix = &value
		case "machine_cidr":
			value := iterator.ReadString()
			object.machineCIDR = &value
		case "pod_cidr":
			value := iterator.ReadString()
			object.podCIDR = &value
		case "service_cidr":
			value := iterator.ReadString()
			object.serviceCIDR = &value
		default:
			iterator.ReadAny()
		}
	}
	return object
}
