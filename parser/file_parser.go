package parser

import (
	"strings"
	"risk_auth/db"
	"encoding/json"
	"fmt"
)

func FileParse(content string) {
	contentList := strings.Split(content, "\n")
	for _, ele := range contentList {
    	if strings.Contains(ele, "authentication_type"){
    		stringJson := strings.Split(ele, " [AUDIT] ")[1]
    		var json_object map[string]interface{}
    		err := json.Unmarshal([]byte(stringJson), &json_object)
		    if err != nil {
		        panic(err)
		    }
		    var interfaceUser interface{} = json_object["distinguished_name_user"]
		    username := fmt.Sprintf("%v", interfaceUser)
            var interfaceIP interface{} = json_object["client_ip"]
            ip := fmt.Sprintf("%v", interfaceIP)
            var interfaceDevice interface{} = json_object["distinguished_name_device_id"]
            deviceID := fmt.Sprintf("%v", interfaceDevice)
            var interfaceTimestamp interface{} = json_object["timestamp"]
            timestamp := fmt.Sprintf("%v", interfaceTimestamp)
            var interfaceEvent interface{} = json_object["event_type"]
			eventType := fmt.Sprintf("%v", interfaceEvent)
			if eventType == "authentication_succeeded" {
				db.Set_username(username)
				db.Set_ip(ip)
				db.Set_device(deviceID)
				db.Set_successfulLoginDict(username, timestamp)
			}else if eventType == "authentication_failed" {
				db.Set_failedLoginDict(username, timestamp)
				db.Set_failedLoginCount(timestamp)
            }

    	}
    }

}