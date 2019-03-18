/*
 * Copyright 2019 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package image

import (
	"github.com/opencontainers/go-digest"
	"strings"
)

// Id is an image id, which happens to be represented as a digest string. An image id
// is based on the binary contents of an image, but not with its name (see Digest).
type Id struct {
	dig digest.Digest
}

var EmptyId Id

func init() {
	EmptyId = Id{""}
}

func NewId(id string) Id {
	return Id{digest.Digest(id)}
}

func (id Id) String() string {
	return string(id.dig)
}

// Filename returns a filesystem-friendly name for this image id.
// Filenames in Windows cannot contain ":".
func (id Id) Filename() string {
	return strings.Replace(id.String(), ":", "-", -1)
}
