// Auto generated file!

package wargaming

import (
	"context"
	"github.com/wows-tools/wows-stats/wargaming/internal"
	"github.com/wows-tools/wows-stats/wargaming/wotb"
	"strconv"
	"strings"
)

// TournamentsTables returns tournament brackets.
//
// https://developers.wargaming.net/reference/all/wotb/tournaments/tables
//
// realm:
//
//	Valid realms: RealmAsia, RealmEu, RealmNa
//
// stageId:
//
//	Stage ID that can be retrieved from the Tournaments Stages method.
//
// tournamentId:
//
//	Tournament ID that can be retrieved from the Tournaments list method.
func (service *WotbService) TournamentsTables(ctx context.Context, realm Realm, stageId int, tournamentId int, options *wotb.TournamentsTablesOptions) ([]*wotb.TournamentsTables, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"stage_id":      strconv.Itoa(stageId),
		"tournament_id": strconv.Itoa(tournamentId),
	}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.GroupId != nil {
			reqParam["group_id"] = internal.SliceIntToString(options.GroupId, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
	}

	var data []*wotb.TournamentsTables
	err := service.client.getRequest(ctx, sectionWotb, realm, "/tournaments/tables/", reqParam, &data, nil)
	return data, err
}