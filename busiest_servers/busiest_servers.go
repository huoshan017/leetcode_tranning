package busiest_servers

type RequestInfo struct {
	arrival int
	load    int
}

type ServerInfo struct {
	reqList []*RequestInfo
}

// 检测服务器并插入
func CheckServerInsert(req *RequestInfo, serv *ServerInfo) bool {
	if serv.reqList == nil {
		serv.reqList = make([]*RequestInfo, 0)
	}

	if len(serv.reqList) == 0 {
		serv.reqList = append(serv.reqList, req)
		return true
	}

	insert := false
	a := req.arrival + req.load - 1
	for i := 0; i < len(serv.reqList); i++ {
		if a < serv.reqList[i].arrival {
			if i == 0 {
				head := []*RequestInfo{req}
				serv.reqList = append(head, serv.reqList...)
			} else {
				serv.reqList = append(serv.reqList[:i], append([]*RequestInfo{req}, serv.reqList[i:]...)...)
			}
			insert = true
			break
		} else if req.arrival > serv.reqList[i].arrival+serv.reqList[i].load-1 {
			if i == len(serv.reqList)-1 {
				serv.reqList = append(serv.reqList, req)
				insert = true
				break
			}
		}
	}
	return insert
}

// 最忙的服务器
func BusiestServers(k int, reqInfos []*RequestInfo) []*ServerInfo {
	if k <= 0 {
		return nil
	}

	servers := make([]*ServerInfo, k)
	for i := 0; i < len(reqInfos); i++ {
		for j := 0; j < k; j++ {
			if servers[j] == nil {
				servers[j] = &ServerInfo{}
			}
			if CheckServerInsert(reqInfos[i], servers[j]) {
				break
			}
		}
	}

	return servers
}
