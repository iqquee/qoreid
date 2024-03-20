package qoreid

type (
	// VerifyNinWithNinReq request object
	VerifyNinWithNinReq struct {
		IdNumber   string `json:"IdNumber" validate:"required"`
		Firstname  string `json:"firstname" validate:"required"`
		Lastname   string `json:"lastname" validate:"required"`
		Middlename string `json:"middlename"`
		DOB        string `json:"dob"` // Format YYYY-MM-DD
		Phone      string `json:"phone"`
		Email      string `json:"email"`
		Gender     string `json:"gender"`
	}

	// VerifyNinWithNinRes response object
	VerifyNinWithNinRes struct {
		ID        int `json:"id"`
		Applicant struct {
			Firstname string `json:"firstname"`
			Lastname  string `json:"lastname"`
		} `json:"applicant"`
		Summary struct {
			NINCheck struct {
				Status       string `json:"status"`
				FieldMatches struct {
					Firstname bool `json:"firstname"`
					Lastname  bool `json:"lastname"`
				} `json:"fieldMatches"`
			} `json:"nin_check"`
		} `json:"summary"`
		Status struct {
			State  string `json:"state"`
			Status string `json:"status"`
		} `json:"status"`
		NIN struct {
			NIN        string `json:"nin"`
			Firstname  string `json:"firstname"`
			Lastname   string `json:"lastname"`
			Middlename string `json:"middlename"`
			Phone      string `json:"phone"`
			Gender     string `json:"gender"`
			Birthdate  string `json:"birthdate"`
			Photo      string `json:"photo"`
			Address    string `json:"address"`
		} `json:"nin"`
	}
)
