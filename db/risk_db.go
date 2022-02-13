package db

import "time"

var SuccessfulLoginDict = make(map[string]string)
var FailedLoginDict = make(map[string]string)
var UsernameList = []string{}
var IPList = []string{}
var DeviceIDList = []string{}
var FailedLoginCount = 0


func Contains(sli []string, ele string) bool {
    for _, i := range sli {
        if i == ele {
            return true
        }
    }
    return false
}

func Set_username(username string) {
    if !Contains(UsernameList, username){
        UsernameList = append(UsernameList, username)
    }
}

func Set_ip(ip string) {
    if !Contains(IPList, ip){
        IPList = append(IPList, ip)
    }
}

func Set_device(deviceID string) {
    if !Contains(DeviceIDList, deviceID){
        DeviceIDList = append(DeviceIDList, deviceID)
    }
}

func Set_successfulLoginDict(username string, timestamp string) {
    newTime, _ := time.Parse(time.RFC3339, timestamp)
    if _, ok := SuccessfulLoginDict[username]; ok {
        oldTime, _ := time.Parse(time.RFC3339,
                                 SuccessfulLoginDict[username])
        if newTime.After(oldTime) {
            SuccessfulLoginDict[username] = timestamp
        }
    } else {
        SuccessfulLoginDict[username] = timestamp
    }
}

func Set_failedLoginDict(username string, timestamp string) {
    newTime, _ := time.Parse(time.RFC3339, timestamp)
    if _, ok := FailedLoginDict[username]; ok {
        oldTime, _ := time.Parse(time.RFC3339,
                                 FailedLoginDict[username])
        if newTime.After(oldTime) {
            FailedLoginDict[username] = timestamp
        }
    } else {
        FailedLoginDict[username] = timestamp
    }
}


func Set_failedLoginCount(timestamp string) {
    today := time.Now()
    lastMonday, lastSunday := WeekRangeDate(today)
    checkTime, _ := time.Parse(time.RFC3339, timestamp)
    if checkTime.After(lastMonday) && checkTime.Before(lastSunday){
        FailedLoginCount += 1
    }
    FailedLoginCount += 1
}

func WeekRangeDate(date time.Time) (time.Time, time.Time) {
    offset := (int(time.Monday) - int(date.Weekday()) - 14) % 14
    lastMonday := date.Add(time.Duration(offset*24) * time.Hour)
    offset = (int(time.Sunday) - int(date.Weekday()) - 14) % 14
    lastSunday := date.Add(time.Duration(offset*24) * time.Hour)
    return lastMonday, lastSunday
}

