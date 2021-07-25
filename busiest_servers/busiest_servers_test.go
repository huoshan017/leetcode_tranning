package busiest_servers

import (
	"testing"
)

func TestBusiestServers(t *testing.T) {
	servers := BusiestServers(3, []*RequestInfo{{1, 5}, {2, 2}, {3, 10}, {4, 3}, {8, 1}, {9, 2}, {10, 2}})
	t.Logf("servers: ")
	for i := 0; i < len(servers); i++ {
		if servers[i] == nil {
			continue
		}
		t.Logf("  %v ", i+1)
		for j := 0; j < len(servers[i].reqList); j++ {
			t.Logf("    %+v", servers[i].reqList[j])
		}
	}
}
