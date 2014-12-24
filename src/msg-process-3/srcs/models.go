package main

import (
	"fmt"
	"github.com/belogik/goes"
	"github.com/opesun/nested"
)

func searchTopKeyword(conn *goes.Connection) ([]keywordContainer, error) {
	keywords := make([]keywordContainer, 10)

	searchResults, err := conn.Search(QueryTopKeywords, []string{"mrdrive"}, []string{"keywords"}, nil)
	if err != nil {
		return nil, err
	}

	for _, aggregation := range searchResults.Aggregations {
		for key, value := range aggregation.Buckets() {
			var keyword keywordContainer

			if group_by_ips, ok := value["group_by_ips"].(map[string]interface{}); ok {
				keyword.keyword = value["key"].(string)
				if keyword.nb_used, ok = group_by_ips["sum_used"].(map[string]interface{})["value"].(float64); ok {
					keywords[key] = keyword
				} else {
					return nil, fmt.Errorf("Field sum_used is in wrong format")
				}
			} else {
				return nil, fmt.Errorf("Field group_by_ips is in wrong format")
			}
		}
	}
	return keywords, nil
}


func searchTopIps(conn *goes.Connection) ([]IpContainer, error) {
	ips := make([]IpContainer, 10)

	searchResults, err := conn.Search(QueryTopIps, []string{"mrdrive"}, []string{"keywords"}, nil)
	if err != nil {
		return nil, err
	}

	for keyIp, value := range (((searchResults.Aggregations["ips"])["group_by_ips"]).(map[string]interface{})["buckets"]).([]interface{}) {
		var ipContainer IpContainer
		var ok bool

		if ipContainer.ip, ok = nested.GetStr(value, "key_as_string"); ok {
			ipContainer.keywords = make([]keywordContainer, 10)
			if keywords, ok := nested.GetS(value, "ip_to_keyword.top_keyword_per_ip.buckets"); ok {
				for keyKeyword, keyword := range keywords {
					if ipContainer.keywords[keyKeyword].keyword, ok = nested.GetStr(keyword, "key"); ok {
						nb_used, _ := nested.Get(keyword, "ips.sum_used.value")
						ipContainer.keywords[keyKeyword].nb_used, _ = nb_used.(float64)
					}
					ips[keyIp] = ipContainer
				}
			} else {
				return nil, fmt.Errorf("Field top_keyword_per_ip is in wrong format")
			}
		}
	}
	return ips, nil
}
