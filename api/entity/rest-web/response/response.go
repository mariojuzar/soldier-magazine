package response

type VerifiedMagazineStatusResponse struct {
	SoldierID			uint	`json:"soldier_id"`
	IsVerifiedMagazine 	bool	`json:"is_verified_magazine"`
}