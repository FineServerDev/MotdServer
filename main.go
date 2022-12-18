package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	//创建post服务器
	http.HandleFunc("/api/v1/motd/be", handlemotdbe)
	http.HandleFunc("/api/v1/motd/je", handlemotdje)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}

type errers struct {
	Status string `json:"status"`
	Err    string `json:"err"`
}

func handlemotdje(w http.ResponseWriter, r *http.Request) {
	// 调用json包的解析，解析请求body
	r.ParseForm()
	username, errr := r.Form["ip"]
	if errr == false {
		w.Header().Set("Content-Type", "application/json")
		var errer errers
		errer.Status = "fail"
		errer.Err = "error"
		json.NewEncoder(w).Encode(errer)
		return
	}
	pwd, errr := r.Form["port"]
	if errr == false {
		w.Header().Set("Content-Type", "application/json")
		var errer errers
		errer.Status = "fail"
		errer.Err = "error"
		json.NewEncoder(w).Encode(errer)
		return
	}
	je, err := MotdJava(username[0] + ":" + pwd[0])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		var errer errers
		errer.Status = "fail"
		errer.Err = err.Error()
		json.NewEncoder(w).Encode(errer)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(je)
}

func handlemotdbe(w http.ResponseWriter, r *http.Request) {
	// 调用json包的解析，解析请求body
	r.ParseForm()
	username, errr := r.Form["ip"]
	if errr == false {
		w.Header().Set("Content-Type", "application/json")
		var errer errers
		errer.Status = "fail"
		errer.Err = "error"
		json.NewEncoder(w).Encode(errer)
		return
	}
	pwd, errr := r.Form["port"]
	if errr == false {
		w.Header().Set("Content-Type", "application/json")
		var errer errers
		errer.Status = "fail"
		errer.Err = "error"
		json.NewEncoder(w).Encode(errer)
		return
	}
	be, err := MotdBE(username[0] + ":" + pwd[0])
	if err != nil {
		var errer errers
		errer.Status = "fail"
		errer.Err = err.Error()
		json.NewEncoder(w).Encode(errer)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(*be)
}
