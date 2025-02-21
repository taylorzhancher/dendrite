package caching

import "github.com/matrix-org/gomatrixserverlib/fclient"

type SpaceSummaryRoomsCache interface {
	GetSpaceSummary(roomID string) (r fclient.MSC2946SpacesResponse, ok bool)
	StoreSpaceSummary(roomID string, r fclient.MSC2946SpacesResponse)
}

func (c Caches) GetSpaceSummary(roomID string) (r fclient.MSC2946SpacesResponse, ok bool) {
	return c.SpaceSummaryRooms.Get(roomID)
}

func (c Caches) StoreSpaceSummary(roomID string, r fclient.MSC2946SpacesResponse) {
	c.SpaceSummaryRooms.Set(roomID, r)
}
