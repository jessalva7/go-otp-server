package repository

import (
	"github.com/dgraph-io/ristretto"
	"log"
	"time"
)

type ristrettoRepository struct {
	cache *ristretto.Cache
	ttlCache time.Duration
}

func NewRistrettoRepository( config *ristretto.Config, ttlCache time.Duration ) Authenticate {

	cache, err := ristretto.NewCache(config)
	if err != nil{
		log.Fatal( err.Error() )
	}

	return &ristrettoRepository{
		cache: cache,
		ttlCache: ttlCache,
	}

}

func (r *ristrettoRepository) Authenticate( phoneNumber string, otp int )  bool{

	if foundOtp, isFound := r.cache.Get( phoneNumber ); isFound && foundOtp.(int) == otp {

		r.RemoveIfPresent( phoneNumber )
		return true
	}

	return false

}

func (r *ristrettoRepository) RemoveIfPresent( phoneNumber string )  bool{

	if _, isFound := r.cache.Get( phoneNumber ); isFound {
		r.cache.Del( phoneNumber )
		return true
	}

	return false

}

func (r *ristrettoRepository) SaveOTP( phoneNumber string, otp int ){

	r.RemoveIfPresent( phoneNumber )
	// setting cost to 0, so that Coster will be ran to find true cost
	r.cache.SetWithTTL( phoneNumber, otp, 0, r.ttlCache )

}