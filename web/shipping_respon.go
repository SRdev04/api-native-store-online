package web

type ShipingRespon struct {
	User_Id   int    `json:"user_id"`
	Kota      string `json:"kota"`
	Kecamatan string `json:"kecamatan"`
	Kelurahan string `json:"kelurahan"`
	Jalan     string `json:"jalan"`
	Wa        int    `json:"wa"`
}
