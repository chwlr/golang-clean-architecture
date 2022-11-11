package model

type CreateFamilyRequest struct {
	Fl_id       int    `json:"fl_id"`
	Cst_id      int    `json:"cst_id"`
	Fl_relation string `json:"fl_relation"`
	Fl_name     string `json:"fl_name"`
	Fl_dob      string `json:"fl_dob"`
}

type CreateFamilyResponse struct {
	Fl_id       int    `json:"fl_id"`
	Cst_id      int    `json:"cst_id"`
	Fl_relation string `json:"fl_relation"`
	Fl_name     string `json:"fl_name"`
	Fl_dob      string `json:"fl_dob"`
}

type GetFamilyResponse struct {
	Fl_id       int    `json:"fl_id"`
	Cst_id      int    `json:"cst_id"`
	Fl_relation string `json:"fl_relation"`
	Fl_name     string `json:"fl_name"`
	Fl_dob      string `json:"fl_dob"`
}
