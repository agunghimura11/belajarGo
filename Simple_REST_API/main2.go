package main

import (
    "github.com/gorilla/mux"
)

type Sisi struct {
    JenisBangun string `json:"jenis bangun"`
    Panjang     int    `json:"panjang"`
    Lebar       int    `json:"lebar"`
}

type Hasil struct {
    JenisBangun string `json:"jenis bangun"`
    Luas        int    `json:"luas"`
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/api/hitung-luas",Luas)
    log.Fatal(http.ListenAndServe( ":8080", router ))
}

func Luas (w http.ResponseWriter, r *http.Request){
    var sisi Sisi
    var hasilHitung Hasil

    if r.Method != "POST" {
        WrapAPIError(w,r,http.StatusText(http.StatusMethodNotAllowed), http.StatusBadRequest)
        return

    }

    body, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        WrapAPIError(w,r, "cant", http.StatusBadRequest)
        return
    }

    err = json.Unmarshal(body, &sisi); if err != nil {
        WrapAPIError(w,r,"error unmarshal : "+err.Error(), http.StatusInternalServerError)
        return 
    }

    WrapAPIData(w,r,Hasil{
        JenisBangun:sisi.JenisBangun,
        Luas:sisi.RumusLuas(),
    },"success", http.StatusOK)
}

func WrapAPIError(w http.ResponseWriter, r *http.Request, message string)
{
    w.Header().Set( key:"Content-Type", value:"application/json")
    w.WriteHeader(code)
    result,err := json.Marshal(map[string]interface{}{
        "code":code,
        "error_type":http.StatusText(code),
        "error_details":message,
    })

    if err == nil {
        w.Write(result)
    }else{
        log.Println(fmt.Sprintf("error wrap Api error sv", err))
    }
}

func WrapAPIData(w http.ResponseWriter, r *http.Request, data interface(), message string, code int)
{
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    result,err := json.Marshal(map[string]interface{}{
        "code":code,
        "status":message,
        "data":data,
    })

    if err == nil {
        w.Write(result)
    }else{
        log.Println(fmt.Sprintf(format: "error wrap Api error sv", err))
    }
}
