package main

var (
	QueryTopKeywords = map[string]interface{} {
		"size":0,
		"aggs" : map[string]interface{} {
			"group_by_keyword": map[string]interface{} {
				"terms":  map[string]interface{} {
					"field": "keyword",
					"order":  map[string]interface{} {
						"group_by_ips>sum_used": "desc",
					},
				},
				"aggs" :  map[string]interface{} {
					"group_by_ips" :  map[string]interface{} {
						"nested" :  map[string]interface{} {
							"path" : "ips",
						},
						"aggs" :  map[string]interface{} {
							"sum_used" :  map[string]interface{} {
								"sum" :  map[string]interface{} {
									"field" : "ips.used",
								},
							},
						},
					},
				},
			},
		},
	}

	QueryTopIps = map[string]interface{} {
		"size": 0,
		"aggs" : map[string]interface{} {
			"ips" : map[string]interface{} {
				"nested" : map[string]interface{} {
					"path" : "ips",
				},
				"aggs" : map[string]interface{} {
					"group_by_ips" : map[string]interface{} {
						"terms" : map[string]interface{} {
							"field" : "ips.ip",
							"order" : map[string]interface{} {
								"total_used" : "DESC",
							},
						},
						"aggs" : map[string]interface{} {
							"total_used" : map[string]interface{} {
								"sum" : map[string]interface{} {
									"field" : "ips.used",
								},
							},
							"ip_to_keyword": map[string]interface{} {
								"reverse_nested": map[string]interface{} {},
								"aggs": map[string]interface{} {
									"top_keyword_per_ip": map[string]interface{} {
										"terms": map[string]interface{} {
											"field": "keyword",
											"order": map[string]interface{} {
												"ips>sum_used": "DESC",
											},
										},
										"aggs" : map[string]interface{} {
											"ips" : map[string]interface{} {
												"nested" : map[string]interface{} {
													"path" : "ips",
												},
												"aggs" : map[string]interface{} {
													"sum_used" : map[string]interface{} {
														"sum" : map[string]interface{} {
															"field":"ips.used",
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	QueryKeyword = func(keyword string) map[string]interface{} {
		var query map[string]interface{}

		query = map[string]interface{} {
			"query": map[string]interface{} {
				"match": map[string]interface{} {
					"keyword": keyword,
				},
			},
		}
		return query
	}
)
