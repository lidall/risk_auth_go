package handler

import (
     "io/ioutil"
     "log"
     "net/http"
     "risk_auth/db"
     "risk_auth/parser"
     "strconv"
     "strings"
)
 
func Handle_user_check(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        username := r.URL.Query().Get("username")
        answer := db.Contains(db.UsernameList, username)
        log.Printf("Check user known: %s", username)
        w.Write([]byte(strconv.FormatBool(answer) + "\n"))
    default:
        w.WriteHeader(http.StatusNotImplemented)
        w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
    }

}

func Handle_ip_check(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        ip := r.URL.Query().Get("ip")
        log.Printf("Check ip known: %s", ip)
        answer := db.Contains(db.IPList, ip)
        w.Write([]byte(strconv.FormatBool(answer) + "\n"))
    default:
        w.WriteHeader(http.StatusNotImplemented)
        w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
    }

}

func Handle_device_check(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        device := r.URL.Query().Get("device")
        log.Printf("Check device known: %s", device)
        answer := db.Contains(db.DeviceIDList, device)
        w.Write([]byte(strconv.FormatBool(answer) + "\n"))
    default:
        w.WriteHeader(http.StatusNotImplemented)
        w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
    }

}

func Handle_internal_check(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        ip := r.URL.Query().Get("ip")
        log.Printf("Check ip internal: %s", ip)
        if strings.Contains(ip, "10.97.2."){
            w.Write([]byte("true\n"))
        } else {
            w.Write([]byte("false\n"))
        }
    default:
        w.WriteHeader(http.StatusNotImplemented)
        w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
    }

}

func Handle_success_check(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        username := r.URL.Query().Get("username") 
        log.Printf("Check user last successful login date: %s", username)
        if _, ok := db.SuccessfulLoginDict[username]; ok {
            w.Write([]byte(db.SuccessfulLoginDict[username] + "\n"))
        } else {
            w.Write([]byte("No record for this user\n"))
        }
    default:
        w.WriteHeader(http.StatusNotImplemented)
        w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
    }

}

func Handle_fail_check(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        username := r.URL.Query().Get("username") 
        log.Printf("Check user last failed login date: %s", username)
        if _, ok := db.FailedLoginDict[username]; ok {
            w.Write([]byte(db.FailedLoginDict[username] + "\n"))
        } else {
            w.Write([]byte("No record for this user\n"))
        }
    default:
        w.WriteHeader(http.StatusNotImplemented)
        w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
    }

}

func Handle_count_check(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        log.Printf("Check failed count")
        w.Write([]byte(strconv.Itoa(db.FailedLoginCount) + "\n"))
    default:
        w.WriteHeader(http.StatusNotImplemented)
        w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
    }

}

func Handle_log(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "POST":
        reqBody, err := ioutil.ReadAll(r.Body)
        log.Println("Processing log...")
        parser.FileParse(string(reqBody))
        if err != nil {
                log.Fatal(err)
        }
        log.Println("Finshed!")
        w.Write([]byte("Log processed!\n"))
    default:
        w.WriteHeader(http.StatusNotImplemented)
        w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
    }

}
