package repository

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/entity"
)

type Repository interface {
	User() UserRepository
	Hobby() HobbyRepository
	UserHobby() UserHobbyRepository
	TravelSpot() TravelSpotRepository
	TravelSpotHobby() TravelSpotHobbyRepository
	TravelSpotItinerary() TravelSpotItineraryRepository
	TravelSpotItineraryItem() TravelSpotItineraryItemRepository
	TravelSpotDiary() TravelSpotDiaryRepository
	Reservation() ReservationRepository
	Vendor() VendorRepository
	Encounter() EncounterRepository
}

type UserRepository interface {
	GetAll(ctx context.Context) (entity.Users, error)
	FindByID(ctx context.Context, id entity.UserID) (*entity.User, error)
	Create(ctx context.Context, users ...*entity.User) error
	Update(ctx context.Context, user *entity.User) error
}

type HobbyRepository interface {
	Create(ctx context.Context, hobbies ...*entity.Hobby) error
	GetAll(ctx context.Context) (entity.Hobbies, error)
	GetByIDs(ctx context.Context, ids []entity.HobbyID) (entity.Hobbies, error)
	GetByID(ctx context.Context, id entity.HobbyID) (*entity.Hobby, error)
}

type UserHobbyRepository interface {
	Create(ctx context.Context, userHobbies ...*entity.UserHobby) error
	DeleteByUserID(ctx context.Context, userID entity.UserID) error
	GetByUserID(ctx context.Context, userID entity.UserID) (entity.UserHobbies, error)
}

type TravelSpotRepository interface {
	Create(ctx context.Context, travelSpots ...entity.TravelSpot) error
	GetByVendorID(ctx context.Context, vendorID entity.VendorID) (entity.TravelSpots, error)
	FindByID(ctx context.Context, id entity.TravelSpotID) (*entity.TravelSpot, error)
	GetByIDs(ctx context.Context, ids []entity.TravelSpotID) (entity.TravelSpots, error)
}

type TravelSpotHobbyRepository interface {
	GetByHobbyID(ctx context.Context, hobbyID entity.HobbyID) (entity.TravelSpotHobbies, error)
}

type TravelSpotItineraryRepository interface {
	Create(ctx context.Context, travelSpotItineraries ...entity.TravelSpotItinerary) error
	GetByTravelSpotID(ctx context.Context, travelSpotID entity.TravelSpotID) (entity.TravelSpotItineraries, error)
	GetByTravelSpotIDs(ctx context.Context, travelSpotIDs []entity.TravelSpotID) (entity.TravelSpotItineraries, error)
	GetByKinds(ctx context.Context, kinds []string) (entity.TravelSpotItineraries, error)
}

type TravelSpotItineraryItemRepository interface {
	Create(ctx context.Context, travelSpotItineraryItems ...entity.TravelSpotItineraryItem) error
	GetByTravelSpotItineraryIDs(ctx context.Context, travelSpotItineraryIDs []entity.TravelSpotItineraryID) (entity.TravelSpotItineraryItems, error)
}

type TravelSpotDiaryRepository interface {
	Create(ctx context.Context, travelSpotDiaries ...*entity.TravelSpotDiary) error
	Update(ctx context.Context, travelSpotDiary *entity.TravelSpotDiary) error
	FindByID(ctx context.Context, id entity.TravelSpotDiaryID) (*entity.TravelSpotDiary, error)
	GetByIDs(ctx context.Context, ids []entity.TravelSpotDiaryID) (entity.TravelSpotDiaries, error)
	FindByUserIDAndTravelSpotID(ctx context.Context, userID entity.UserID, travelSpotID entity.TravelSpotID) (entity.TravelSpotDiaries, error)
}

type ReservationRepository interface {
	Create(ctx context.Context, reservations ...*entity.Reservation) error
	FindByID(ctx context.Context, id entity.ReservationID) (*entity.Reservation, error)
	FindByUserIDAndTravelSpotID(ctx context.Context, userID entity.UserID, travelSpotID entity.TravelSpotID) (entity.Reservations, error)
	FindSpecialOfferByUserID(ctx context.Context, userID entity.UserID) (*entity.Reservation, error)
	GetEndedReservations(ctx context.Context, userID entity.UserID) (entity.Reservations, error)
	GetUpcomingReservations(ctx context.Context, userID entity.UserID) (entity.Reservations, error)
}

type VendorRepository interface {
	Create(ctx context.Context, vendors ...entity.Vendor) error
	GetAll(ctx context.Context) (entity.Vendors, error)
	FindByID(ctx context.Context, id entity.VendorID) (*entity.Vendor, error)
}

type EncounterRepository interface {
	Create(ctx context.Context, encounters ...entity.Encounter) error
	Update(ctx context.Context, encounter *entity.Encounter) error
	FindByID(ctx context.Context, id entity.EncounterID) (*entity.Encounter, error)
	GetByUserID(ctx context.Context, userID entity.UserID) (entity.Encounters, error)
}
