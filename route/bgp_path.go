package route

import (
	"crypto/sha256"
	"fmt"
	"strings"

	"github.com/bio-routing/tflow2/convert"

	bnet "github.com/bio-routing/bio-rd/net"
	"github.com/bio-routing/bio-rd/protocols/bgp/types"
	"github.com/bio-routing/bio-rd/route/api"
)

// BGPPath represents a set of BGP path attributes
type BGPPath struct {
	BGPPathA          *BGPPathA
	ASPath            *types.ASPath
	ClusterList       *[]uint32
	Communities       *[]uint32
	LargeCommunities  *[]types.LargeCommunity
	UnknownAttributes *[]types.UnknownPathAttribute
	PathIdentifier    uint32
	ASPathLen         uint16
}

// BGPPathA represents cachable BGP path attributes
type BGPPathA struct {
	NextHop         *bnet.IP
	Source          *bnet.IP
	LocalPref       uint32
	MED             uint32
	BGPIdentifier   uint32
	OriginatorID    uint32
	Aggregator      *types.Aggregator
	EBGP            bool
	AtomicAggregate bool
	Origin          uint8
}

func (b *BGPPathA) Dedup() *BGPPathA {
	return bgpC.get(b)
}

func (b *BGPPath) Dedup() *BGPPath {
	b.BGPPathA = b.BGPPathA.Dedup()
	return b
}

// ToProto converts BGPPath to proto BGPPath
func (b *BGPPath) ToProto() *api.BGPPath {
	if b == nil {
		return nil
	}

	a := &api.BGPPath{
		PathIdentifier:    b.PathIdentifier,
		NextHop:           b.BGPPathA.NextHop.ToProto(),
		LocalPref:         b.BGPPathA.LocalPref,
		AsPath:            b.ASPath.ToProto(),
		Origin:            uint32(b.BGPPathA.Origin),
		Med:               b.BGPPathA.MED,
		Ebgp:              b.BGPPathA.EBGP,
		BgpIdentifier:     b.BGPPathA.BGPIdentifier,
		Source:            b.BGPPathA.Source.ToProto(),
		Communities:       make([]uint32, len(*b.Communities)),
		LargeCommunities:  make([]*api.LargeCommunity, len(*b.LargeCommunities)),
		UnknownAttributes: make([]*api.UnknownPathAttribute, len(*b.UnknownAttributes)),
		OriginatorId:      b.BGPPathA.OriginatorID,
		ClusterList:       make([]uint32, len(*b.ClusterList)),
	}

	copy(a.Communities, *b.Communities)
	copy(a.ClusterList, *b.ClusterList)

	for i := range *b.LargeCommunities {
		a.LargeCommunities[i] = (*b.LargeCommunities)[i].ToProto()
	}

	for i := range *b.UnknownAttributes {
		a.UnknownAttributes[i] = (*b.UnknownAttributes)[i].ToProto()
	}

	return a
}

// BGPPathFromProtoBGPPath converts a proto BGPPath to BGPPath
func BGPPathFromProtoBGPPath(pb *api.BGPPath) *BGPPath {
	p := &BGPPath{
		BGPPathA: &BGPPathA{
			NextHop:       bnet.IPFromProtoIP(*pb.NextHop),
			LocalPref:     pb.LocalPref,
			OriginatorID:  pb.OriginatorId,
			Origin:        uint8(pb.Origin),
			MED:           pb.Med,
			EBGP:          pb.Ebgp,
			BGPIdentifier: pb.BgpIdentifier,
			Source:        bnet.IPFromProtoIP(*pb.Source),
		},
		PathIdentifier: pb.PathIdentifier,
		ASPath:         types.ASPathFromProtoASPath(pb.AsPath),
	}

	p = p.Dedup()

	communities := make([]uint32, len(pb.Communities))
	p.Communities = &communities

	largeCommunities := make([]types.LargeCommunity, len(pb.LargeCommunities))
	p.LargeCommunities = &largeCommunities

	unknownAttr := make([]types.UnknownPathAttribute, len(pb.UnknownAttributes))
	p.UnknownAttributes = &unknownAttr

	cl := make([]uint32, len(pb.ClusterList))
	p.ClusterList = &cl

	for i := range pb.Communities {
		(*p.Communities)[i] = pb.Communities[i]
	}

	for i := range pb.LargeCommunities {
		(*p.LargeCommunities)[i] = types.LargeCommunityFromProtoCommunity(pb.LargeCommunities[i])
	}

	for i := range pb.UnknownAttributes {
		(*p.UnknownAttributes)[i] = types.UnknownPathAttributeFromProtoUnknownPathAttribute(pb.UnknownAttributes[i])
	}

	for i := range pb.ClusterList {
		(*p.ClusterList)[i] = pb.ClusterList[i]
	}

	return p
}

// Length get's the length of serialized path
func (b *BGPPath) Length() uint16 {
	asPathLen := uint16(3)
	for _, segment := range *b.ASPath {
		asPathLen++
		asPathLen += uint16(4 * len(segment.ASNs))
	}

	communitiesLen := uint16(0)
	if b.Communities != nil && len(*b.Communities) != 0 {
		communitiesLen += 3 + uint16(len(*b.Communities)*4)
	}

	largeCommunitiesLen := uint16(0)
	if b.LargeCommunities != nil && len(*b.LargeCommunities) != 0 {
		largeCommunitiesLen += 3 + uint16(len(*b.LargeCommunities)*12)
	}

	clusterListLen := uint16(0)
	if b.ClusterList != nil && len(*b.ClusterList) != 0 {
		clusterListLen += 3 + uint16(len(*b.ClusterList)*4)
	}

	unknownAttributesLen := uint16(0)
	if b.UnknownAttributes != nil {
		for _, unknownAttr := range *b.UnknownAttributes {
			unknownAttributesLen += unknownAttr.WireLength()
		}
	}

	originatorID := uint16(0)
	if b.BGPPathA.OriginatorID != 0 {
		originatorID = 4
	}

	return communitiesLen + largeCommunitiesLen + 4*7 + 4 + originatorID + asPathLen + unknownAttributesLen
}

// ECMP determines if routes b and c are euqal in terms of ECMP
func (b *BGPPath) ECMP(c *BGPPath) bool {
	return b.BGPPathA.LocalPref == c.BGPPathA.LocalPref &&
		b.ASPathLen == c.ASPathLen &&
		b.BGPPathA.MED == c.BGPPathA.MED &&
		b.BGPPathA.Origin == c.BGPPathA.Origin
}

// Equal checks if paths are equal
func (b *BGPPath) Equal(c *BGPPath) bool {
	if b.PathIdentifier != c.PathIdentifier {
		return false
	}

	return b.Select(c) == 0
}

// Select returns negative if b < c, 0 if paths are equal, positive if b > c
func (b *BGPPath) Select(c *BGPPath) int8 {
	if c.BGPPathA.LocalPref < b.BGPPathA.LocalPref {
		return 1
	}

	if c.BGPPathA.LocalPref > b.BGPPathA.LocalPref {
		return -1
	}

	// 9.1.2.2.  Breaking Ties (Phase 2)

	// a)
	if c.ASPathLen > b.ASPathLen {
		return 1
	}

	if c.ASPathLen < b.ASPathLen {
		return -1
	}

	// b)
	if c.BGPPathA.Origin > b.BGPPathA.Origin {
		return 1
	}

	if c.BGPPathA.Origin < b.BGPPathA.Origin {
		return -1
	}

	// c)
	if c.BGPPathA.MED > b.BGPPathA.MED {
		return 1
	}

	if c.BGPPathA.MED < b.BGPPathA.MED {
		return -1
	}

	// d)
	if c.BGPPathA.EBGP && !b.BGPPathA.EBGP {
		return -1
	}

	if !c.BGPPathA.EBGP && b.BGPPathA.EBGP {
		return 1
	}

	// e) TODO: interior cost (hello IS-IS and OSPF)

	// f) + RFC4456 9. (Route Reflection)
	bgpIdentifierC := c.BGPPathA.BGPIdentifier
	bgpIdentifierB := b.BGPPathA.BGPIdentifier

	// IF an OriginatorID (set by an RR) is present, use this instead of Originator
	if c.BGPPathA.OriginatorID != 0 {
		bgpIdentifierC = c.BGPPathA.OriginatorID
	}

	if b.BGPPathA.OriginatorID != 0 {
		bgpIdentifierB = b.BGPPathA.OriginatorID
	}

	if bgpIdentifierC < bgpIdentifierB {
		return 1
	}

	if bgpIdentifierC > bgpIdentifierB {
		return -1
	}

	// Additionally check for the shorter ClusterList
	if len(*c.ClusterList) < len(*b.ClusterList) {
		return 1
	}

	if len(*c.ClusterList) > len(*b.ClusterList) {
		return -1
	}

	// g)
	if c.BGPPathA.Source.Compare(b.BGPPathA.Source) == -1 {
		return 1
	}

	if c.BGPPathA.Source.Compare(b.BGPPathA.Source) == 1 {
		return -1
	}

	if c.BGPPathA.NextHop.Compare(b.BGPPathA.NextHop) == -1 {
		return 1
	}

	if c.BGPPathA.NextHop.Compare(b.BGPPathA.NextHop) == 1 {
		return -1
	}

	return 0
}

func (b *BGPPath) betterECMP(c *BGPPath) bool {
	if c.BGPPathA.LocalPref < b.BGPPathA.LocalPref {
		return false
	}

	if c.BGPPathA.LocalPref > b.BGPPathA.LocalPref {
		return true
	}

	if c.ASPathLen > b.ASPathLen {
		return false
	}

	if c.ASPathLen < b.ASPathLen {
		return true
	}

	if c.BGPPathA.Origin > b.BGPPathA.Origin {
		return false
	}

	if c.BGPPathA.Origin < b.BGPPathA.Origin {
		return true
	}

	if c.BGPPathA.MED > b.BGPPathA.MED {
		return false
	}

	if c.BGPPathA.MED < b.BGPPathA.MED {
		return true
	}

	return false
}

func (b *BGPPath) better(c *BGPPath) bool {
	if b.betterECMP(c) {
		return true
	}

	if c.BGPPathA.BGPIdentifier < b.BGPPathA.BGPIdentifier {
		return true
	}

	if c.BGPPathA.Source.Compare(b.BGPPathA.Source) == -1 {
		return true
	}

	return false
}

// Print all known information about a route in logfile friendly format
func (b *BGPPath) String() string {
	buf := &strings.Builder{}

	origin := ""
	switch b.BGPPathA.Origin {
	case 0:
		origin = "Incomplete"
	case 1:
		origin = "EGP"
	case 2:
		origin = "IGP"
	}

	bgpType := "internal"
	if b.BGPPathA.EBGP {
		bgpType = "external"
	}

	fmt.Fprintf(buf, "Local Pref: %d, ", b.BGPPathA.LocalPref)
	fmt.Fprintf(buf, "Origin: %s, ", origin)
	fmt.Fprintf(buf, "AS Path: %v, ", b.ASPath)
	fmt.Fprintf(buf, "BGP type: %s, ", bgpType)
	fmt.Fprintf(buf, "NEXT HOP: %s, ", b.BGPPathA.NextHop)
	fmt.Fprintf(buf, "MED: %d, ", b.BGPPathA.MED)
	fmt.Fprintf(buf, "Path ID: %d, ", b.PathIdentifier)
	fmt.Fprintf(buf, "Source: %s, ", b.BGPPathA.Source)
	fmt.Fprintf(buf, "Communities: %v, ", b.Communities)
	fmt.Fprintf(buf, "LargeCommunities: %v", b.LargeCommunities)

	if b.BGPPathA.OriginatorID != 0 {
		oid := convert.Uint32Byte(b.BGPPathA.OriginatorID)
		fmt.Fprintf(buf, ", OriginatorID: %d.%d.%d.%d", oid[0], oid[1], oid[2], oid[3])
	}
	if b.ClusterList != nil {
		fmt.Fprintf(buf, ", ClusterList %s", b.ClusterListString())
	}

	return buf.String()
}

// Print all known information about a route in human readable form
func (b *BGPPath) Print() string {
	buf := &strings.Builder{}

	origin := ""
	switch b.BGPPathA.Origin {
	case 0:
		origin = "Incomplete"
	case 1:
		origin = "EGP"
	case 2:
		origin = "IGP"
	}

	bgpType := "internal"
	if b.BGPPathA.EBGP {
		bgpType = "external"
	}

	fmt.Fprintf(buf, "\t\tLocal Pref: %d\n", b.BGPPathA.LocalPref)
	fmt.Fprintf(buf, "\t\tOrigin: %s\n", origin)
	fmt.Fprintf(buf, "\t\tAS Path: %v\n", b.ASPath)
	fmt.Fprintf(buf, "\t\tBGP type: %s\n", bgpType)
	fmt.Fprintf(buf, "\t\tNEXT HOP: %s\n", b.BGPPathA.NextHop)
	fmt.Fprintf(buf, "\t\tMED: %d\n", b.BGPPathA.MED)
	fmt.Fprintf(buf, "\t\tPath ID: %d\n", b.PathIdentifier)
	fmt.Fprintf(buf, "\t\tSource: %s\n", b.BGPPathA.Source)
	fmt.Fprintf(buf, "\t\tCommunities: %v\n", b.Communities)
	fmt.Fprintf(buf, "\t\tLargeCommunities: %v\n", b.LargeCommunities)

	if b.BGPPathA.OriginatorID != 0 {
		oid := convert.Uint32Byte(b.BGPPathA.OriginatorID)
		fmt.Fprintf(buf, "\t\tOriginatorID: %d.%d.%d.%d\n", oid[0], oid[1], oid[2], oid[3])
	}
	if b.ClusterList != nil {
		fmt.Fprintf(buf, "\t\tClusterList %s\n", b.ClusterListString())
	}

	return buf.String()
}

// Prepend the given BGPPath with the given ASN given times
func (b *BGPPath) Prepend(asn uint32, times uint16) {
	if times == 0 {
		return
	}

	if len(*b.ASPath) == 0 {
		b.insertNewASSequence()
	}

	first := (*b.ASPath)[0]
	if first.Type == types.ASSet {
		b.insertNewASSequence()
	}

	for i := 0; i < int(times); i++ {
		if len(*b.ASPath) == types.MaxASNsSegment {
			b.insertNewASSequence()
		}

		old := (*b.ASPath)[0].ASNs
		asns := make([]uint32, len(old)+1)
		copy(asns[1:], old)
		asns[0] = asn
		(*b.ASPath)[0].ASNs = asns
	}

	b.ASPathLen = b.ASPath.Length()
}

func (b *BGPPath) insertNewASSequence() {
	pa := make(types.ASPath, len(*b.ASPath)+1)
	copy(pa[1:], (*b.ASPath))
	pa[0] = types.ASPathSegment{
		ASNs: make([]uint32, 0),
		Type: types.ASSequence,
	}

	b.ASPath = &pa
}

// Copy creates a deep copy of a BGPPath
func (b *BGPPath) Copy() *BGPPath {
	if b == nil {
		return nil
	}

	cp := *b

	asPath := make(types.ASPath, len(*cp.ASPath))
	cp.ASPath = &asPath
	copy(*cp.ASPath, *b.ASPath)

	if cp.Communities != nil {
		communities := make([]uint32, len(*cp.Communities))
		cp.Communities = &communities
		copy(*cp.Communities, *b.Communities)
	}

	if cp.LargeCommunities != nil {
		largeCommunities := make([]types.LargeCommunity, len(*cp.LargeCommunities))
		cp.LargeCommunities = &largeCommunities
		copy(*cp.LargeCommunities, *b.LargeCommunities)
	}

	if b.ClusterList != nil {
		clusterList := make([]uint32, len(*cp.ClusterList))
		cp.ClusterList = &clusterList
		copy(*cp.ClusterList, *b.ClusterList)
	}

	return &cp
}

// ComputeHash computes an hash over all attributes of the path
func (b *BGPPath) ComputeHash() string {
	s := fmt.Sprintf("%s\t%d\t%v\t%d\t%d\t%v\t%d\t%s\t%v\t%v\t%d\t%d\t%v",
		b.BGPPathA.NextHop,
		b.BGPPathA.LocalPref,
		*b.ASPath,
		b.BGPPathA.Origin,
		b.BGPPathA.MED,
		b.BGPPathA.EBGP,
		b.BGPPathA.BGPIdentifier,
		b.BGPPathA.Source,
		b.Communities,
		b.LargeCommunities,
		b.PathIdentifier,
		b.BGPPathA.OriginatorID,
		b.ClusterList)

	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}

// CommunitiesString returns the formated communities
func (b *BGPPath) CommunitiesString() string {
	str := &strings.Builder{}

	for i, com := range *b.Communities {
		if i > 0 {
			str.WriteByte(' ')
		}
		str.WriteString(types.CommunityStringForUint32(com))
	}

	return str.String()
}

// ClusterListString returns the formated ClusterList
func (b *BGPPath) ClusterListString() string {
	str := &strings.Builder{}

	for i, cid := range *b.ClusterList {
		if i > 0 {
			str.WriteByte(' ')
		}
		octes := convert.Uint32Byte(cid)

		fmt.Fprintf(str, "%d.%d.%d.%d", octes[0], octes[1], octes[2], octes[3])
	}

	return str.String()
}

// LargeCommunitiesString returns the formated communities
func (b *BGPPath) LargeCommunitiesString() string {
	str := &strings.Builder{}

	for i, com := range *b.LargeCommunities {
		if i > 0 {
			str.WriteByte(' ')
		}
		str.WriteString(com.String())
	}

	return str.String()
}
