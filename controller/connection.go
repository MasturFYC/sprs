package controller

type Response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// func EnableCors(w *http.ResponseWriter) {
// 	//	(*w).Header().Set("Access-Control-Allow-Origin", "http://pixel.id:8080")
// 	//(*w).Header().Set("Context-Type", "application/x-www-form-urlencoded")
// 	(*w).Header().Set("Content-Type", "application/json")
// 	//(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
// }

func CombineString(a string, b string) string {
	if len(b) == 0 {
		return a
	}
	return a + ", " + b
}

// var Sql = InitDatabase()
