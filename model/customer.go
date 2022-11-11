package model

type CreateCustomerRequest struct {
	Cst_id         int    `json:"cst_id"`
	Nationality_id int    `json:"nationality_id"`
	Name           string `json:"name"`
	Dob            string `json:"dob"`
	PhoneNumber    string `json:"phoneNumber"`
	Email          string `json:"email"`
}

type CreateCustomerResponse struct {
	Cst_id         int    `json:"cst_id"`
	Nationality_id int    `json:"nationality_id"`
	Name           string `json:"name"`
	Dob            string `json:"dob"`
	PhoneNumber    string `json:"phoneNumber"`
	Email          string `json:"email"`
}

type GetCustomerResponse struct {
	Cst_id         int    `json:"cst_id"`
	Nationality_id int    `json:"nationality_id"`
	Name           string `json:"name"`
	Dob            string `json:"dob"`
	PhoneNumber    string `json:"phoneNumber"`
	Email          string `json:"email"`
}
