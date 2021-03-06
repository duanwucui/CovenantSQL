/*
 * Copyright 2018 The CovenantSQL Authors.
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

package proto

import (
	"strings"
	"testing"

	"time"

	"github.com/CovenantSQL/CovenantSQL/crypto/hash"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v2"
)

func TestNode_InitNodeCryptoInfo(t *testing.T) {
	Convey("InitNodeCryptoInfo", t, func() {
		node := NewNode()
		node.InitNodeCryptoInfo(1000 * time.Millisecond)
		hashTmp := hash.THashH(append(node.PublicKey.Serialize(),
			node.Nonce.Bytes()...))
		t.Logf("ComputeBlockNonce, Difficulty: %d, nonce %v, hash %s",
			node.ID.Difficulty(), node.Nonce, hashTmp.String())

		So(node.ID.Difficulty(), ShouldBeGreaterThan, 1)
		So(hashTmp.String(), ShouldEqual, node.ID)
	})
	Convey("NodeID Difficulty", t, func() {
		var node NodeID
		So(node.Difficulty(), ShouldEqual, -1)
		node = NodeID("")
		So(node.Difficulty(), ShouldEqual, -1)
		node = NodeID("1")
		So(node.Difficulty(), ShouldEqual, -1)
		node = NodeID("00000000011a34cb8142780f692a4097d883aa2ac8a534a070a134f11bcca573")
		So(node.Difficulty(), ShouldEqual, 39)
		node = NodeID("#0000000011a34cb8142780f692a4097d883aa2ac8a534a070a134f11bcca573")
		So(node.Difficulty(), ShouldEqual, -1)
		So((*NodeID)(nil).Difficulty(), ShouldEqual, -1)
	})
}

func TestNodeKey_Less(t *testing.T) {
	Convey("NodeID Difficulty", t, func() {
		k1 := NodeKey{}
		k2 := NodeKey{
			Hash: hash.Hash{0xa},
		}
		So(k1.Less(&k1), ShouldBeFalse)
		So(k2.Less(&k1), ShouldBeFalse)
		So(k1.Less(&k2), ShouldBeTrue)
	})
}

func TestServerRoles_Contains(t *testing.T) {
	Convey("ServerRoles Contains", t, func() {
		ss := make(ServerRoles, 0)

		So(ss.Contains(Follower), ShouldBeFalse)
		So(ss.Contains(Unknown), ShouldBeFalse)
		ss = append(ss, Leader)
		ss = append(ss, Follower)
		So(ss.Contains(Leader), ShouldBeTrue)
	})
}

func unmarshalAndMarshal(str string) string {
	var role ServerRole
	yaml.Unmarshal([]byte(str), &role)
	ret, _ := yaml.Marshal(role)

	return strings.TrimSpace(string(ret))
}

func TestServerRole_MarshalYAML(t *testing.T) {
	Convey("marshal unmarshal yaml", t, func() {
		var role ServerRole
		s, _ := role.MarshalYAML()
		So(s, ShouldResemble, "Unknown")
		So(unmarshalAndMarshal("unknown"), ShouldEqual, "Unknown")
		So(unmarshalAndMarshal("leader"), ShouldEqual, "Leader")
		So(unmarshalAndMarshal("follower"), ShouldEqual, "Follower")
		So(unmarshalAndMarshal("miner"), ShouldEqual, "Miner")
		So(unmarshalAndMarshal("client"), ShouldEqual, "Client")
	})
}

func TestNodeID_ToRawNodeID(t *testing.T) {
	Convey("NodeID to RawNodeID", t, func() {
		k1 := RawNodeID{
			Hash: hash.Hash{0xa},
		}
		k1Node := NodeID(k1.String())
		So(k1Node.ToRawNodeID().IsEqual(&k1.Hash), ShouldBeTrue)

		id := "00000000011a34cb8142780f692a4097d883aa2ac8a534a070a134f11bcca573"
		node := NodeID(id)
		So(node.ToRawNodeID().String(), ShouldEqual, id)
		So(node.ToRawNodeID().ToNodeID(), ShouldEqual, node)
	})
}

func TestNodeID_IsEmpty(t *testing.T) {
	Convey("NodeID is empty", t, func() {
		var nodeID NodeID
		So(nodeID.IsEmpty(), ShouldBeTrue)
		var nodeIDPtr *NodeID
		So(nodeIDPtr.IsEmpty(), ShouldBeTrue)
		id := "00000000011a34cb8142780f692a4097d883aa2ac8a534a070a134f11bcca573"
		node := NodeID(id)
		So(node.IsEmpty(), ShouldBeFalse)

		// test nil values with ToNodeID and IsEmpty
		node = (*RawNodeID)(nil).ToNodeID()
		So(node.IsEmpty(), ShouldBeTrue)
	})
}
