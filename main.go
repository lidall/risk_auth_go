package main

import (
    "log"
    "net/http"
    "risk_auth/handler"
)
 

func main() {
    http.HandleFunc("/risk/isuserknown", handler.Handle_user_check)
    http.HandleFunc("/risk/isipknown", handler.Handle_ip_check)
    http.HandleFunc("/risk/isdeviceknown", handler.Handle_device_check)
    http.HandleFunc("/risk/isipinternal", handler.Handle_internal_check)
    http.HandleFunc("/risk/lastsuccessfullogindate", handler.Handle_success_check)
    http.HandleFunc("/risk/lastfailedlogindate", handler.Handle_fail_check)
    http.HandleFunc("/risk/failedlogincountlastweek", handler.Handle_count_check)
    http.HandleFunc("/log", handler.Handle_log)
    log.Println("Listening...")
    http.ListenAndServe(":8080", nil)
}