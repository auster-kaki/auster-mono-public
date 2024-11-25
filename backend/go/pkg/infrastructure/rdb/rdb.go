package rdb

import (
	"cmp"
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/auster-kaki/auster-mono/pkg/app/repository"
	"github.com/auster-kaki/auster-mono/pkg/infrastructure/rdb/table"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/extra/bunslog"
)

type rdb struct {
	user                    *table.User
	hobby                   *table.Hobby
	userHobby               *table.UserHobby
	vendor                  *table.Vendor
	travelSpot              *table.TravelSpot
	travelSpotHobby         *table.TravelSpotHobby
	travelSpotItinerary     *table.TravelSpotItinerary
	travelSpotItineraryItem *table.TravelSpotItineraryItem
	travelSpotDiary         *table.TravelSpotDiary
	reservation             *table.Reservation
	encounter               *table.Encounter
}

func NewDB() (*rdb, error) {
	var (
		host           = cmp.Or(os.Getenv("MYSQL_HOST"), "127.0.0.1")
		dataSourceName = fmt.Sprintf("root:@tcp(%s:3306)/auster?charset=utf8mb4&parseTime=true", host)
	)
	sqlDB, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqlDB, mysqldialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	db.AddQueryHook(bunslog.NewQueryHook(bunslog.WithLogger(slog.Default())))

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &rdb{
		user:                    table.NewUser(db),
		hobby:                   table.NewHobby(db),
		userHobby:               table.NewUserHobby(db),
		vendor:                  table.NewVendor(db),
		travelSpot:              table.NewTravelSpot(db),
		travelSpotHobby:         table.NewTravelSpotHobby(db),
		travelSpotItinerary:     table.NewTravelSpotItinerary(db),
		travelSpotItineraryItem: table.NewTravelSpotItineraryItem(db),
		travelSpotDiary:         table.NewTravelSpotDiary(db),
		reservation:             table.NewReservation(db),
		encounter:               table.NewEncounter(db),
	}, nil
}

func (r *rdb) User() repository.UserRepository {
	return r.user
}

func (r *rdb) Hobby() repository.HobbyRepository { return r.hobby }

func (r *rdb) UserHobby() repository.UserHobbyRepository { return r.userHobby }

func (r *rdb) TravelSpot() repository.TravelSpotRepository {
	return r.travelSpot
}

func (r *rdb) Vendor() repository.VendorRepository {
	return r.vendor
}

func (r *rdb) Encounter() repository.EncounterRepository {
	return r.encounter
}

func (r *rdb) TravelSpotHobby() repository.TravelSpotHobbyRepository {
	return r.travelSpotHobby
}

func (r *rdb) TravelSpotItinerary() repository.TravelSpotItineraryRepository {
	return r.travelSpotItinerary
}

func (r *rdb) TravelSpotItineraryItem() repository.TravelSpotItineraryItemRepository {
	return r.travelSpotItineraryItem
}

func (r *rdb) TravelSpotDiary() repository.TravelSpotDiaryRepository {
	return r.travelSpotDiary
}

func (r *rdb) Reservation() repository.ReservationRepository {
	return r.reservation
}
