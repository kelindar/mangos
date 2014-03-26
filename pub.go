// Copyright 2014 Garrett D'Amore
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sp

// pub is an implementation of the PUB protocol.
type pub struct {
	xpub
}

// Name implements the Protocol Name method.
func (*pub) Name() string {
	return PubName
}

// IsRaw implements the Protocol IsRaw method.
func (*pub) IsRaw() bool {
	return false
}

type pubFactory int

func (pubFactory) NewProtocol() Protocol {
	return new(pub)
}

// PubFactory implements the Protocol Factory for the PUB (Publish) protocol.
var PubFactory pubFactory