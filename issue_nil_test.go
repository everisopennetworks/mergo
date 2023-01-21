package mergo

import (
	"reflect"
	"testing"
)

func TestDeleteNilElement(t *testing.T) {
	tests := []struct {
		dst     map[string]interface{}
		src     map[string]interface{}
		options Config
		want    map[string]interface{}
		name    string
	}{
		{
			map[string]interface{}{
				"/ManagedElement=DU-01/GNBDUFunction=0/NRCellDU=3": map[string]interface{}{
					"body": map[string]interface{}{
						"NRCellDU": []interface{}{
							nil,
							nil,
							nil,
							map[string]interface{}{
								"id": 3,
								"attributes": map[string]interface{}{
									"SystemInformationBlocks": map[string]interface{}{
										"sib1": map[string]interface{}{
											"uac-BarringInfo": map[string]interface{}{
												"uac-BarringForCommon": []interface{}{
													nil,
													nil,
													nil,
													nil,
													nil,
													map[string]interface{}{
														"accessCategory": 5,
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
			map[string]interface{}{
				"/ManagedElement=DU-01/GNBDUFunction=0/NRCellDU=3": map[string]interface{}{
					"body": map[string]interface{}{
						"NRCellDU": []interface{}{
							nil,
							nil,
							nil,
							map[string]interface{}{
								"id": 3,
								"attributes": map[string]interface{}{
									"SystemInformationBlocks": map[string]interface{}{
										"sib1": map[string]interface{}{
											"uac-BarringInfo": map[string]interface{}{
												"uac-BarringForCommon": []interface{}{
													nil,
													nil,
													nil,
													nil,
													nil,
													map[string]interface{}{
														"uac-barringInfoSetIndex": nil,
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
			Config{
				sliceDeepMerge: true,
				Overwrite:      false,
			},
			map[string]interface{}{
				"/ManagedElement=DU-01/GNBDUFunction=0/NRCellDU=3": map[string]interface{}{
					"body": map[string]interface{}{
						"NRCellDU": []interface{}{
							nil,
							nil,
							nil,
							map[string]interface{}{
								"id": 3,
								"attributes": map[string]interface{}{
									"SystemInformationBlocks": map[string]interface{}{
										"sib1": map[string]interface{}{
											"uac-BarringInfo": map[string]interface{}{
												"uac-BarringForCommon": []interface{}{
													nil,
													nil,
													nil,
													nil,
													nil,
													map[string]interface{}{
														"accessCategory":          5,
														"uac-barringInfoSetIndex": nil,
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
			"Sending second the null value",
		},
		{
			map[string]interface{}{
				"/ManagedElement=DU-01/GNBDUFunction=0/NRCellDU=3": map[string]interface{}{
					"body": map[string]interface{}{
						"NRCellDU": []interface{}{
							nil,
							nil,
							nil,
							map[string]interface{}{
								"id": 3,
								"attributes": map[string]interface{}{
									"SystemInformationBlocks": map[string]interface{}{
										"sib1": map[string]interface{}{
											"uac-BarringInfo": map[string]interface{}{
												"uac-BarringForCommon": []interface{}{
													nil,
													nil,
													nil,
													nil,
													nil,
													map[string]interface{}{
														"uac-barringInfoSetIndex": nil,
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
			map[string]interface{}{
				"/ManagedElement=DU-01/GNBDUFunction=0/NRCellDU=3": map[string]interface{}{
					"body": map[string]interface{}{
						"NRCellDU": []interface{}{
							nil,
							nil,
							nil,
							map[string]interface{}{
								"id": 3,
								"attributes": map[string]interface{}{
									"SystemInformationBlocks": map[string]interface{}{
										"sib1": map[string]interface{}{
											"uac-BarringInfo": map[string]interface{}{
												"uac-BarringForCommon": []interface{}{
													nil,
													nil,
													nil,
													nil,
													nil,
													map[string]interface{}{
														"accessCategory": 5,
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
			Config{
				sliceDeepMerge: true,
				Overwrite:      false,
			},
			map[string]interface{}{
				"/ManagedElement=DU-01/GNBDUFunction=0/NRCellDU=3": map[string]interface{}{
					"body": map[string]interface{}{
						"NRCellDU": []interface{}{
							nil,
							nil,
							nil,
							map[string]interface{}{
								"id": 3,
								"attributes": map[string]interface{}{
									"SystemInformationBlocks": map[string]interface{}{
										"sib1": map[string]interface{}{
											"uac-BarringInfo": map[string]interface{}{
												"uac-BarringForCommon": []interface{}{
													nil,
													nil,
													nil,
													nil,
													nil,
													map[string]interface{}{
														"accessCategory":          5,
														"uac-barringInfoSetIndex": nil,
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
			"Sending second any value",
		},
		{
			map[string]interface{}{
				"/ManagedElement=DU-01/GNBDUFunction=0/NRCellDU=3": map[string]interface{}{
					"body": map[string]interface{}{
						"NRCellDU": []interface{}{
							nil,
							nil,
							nil,
							map[string]interface{}{
								"id": 3,
								"attributes": map[string]interface{}{
									"SystemInformationBlocks": map[string]interface{}{
										"sib1": map[string]interface{}{
											"uac-BarringInfo": map[string]interface{}{
												"uac-BarringForCommon": []interface{}{
													nil,
													nil,
													nil,
													nil,
													nil,
													map[string]interface{}{
														"accessCategory":          5,
														"uac-barringInfoSetIndex": 8,
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
			map[string]interface{}{
				"/ManagedElement=DU-01/GNBDUFunction=0/NRCellDU=3": map[string]interface{}{
					"body": map[string]interface{}{
						"NRCellDU": []interface{}{
							nil,
							nil,
							nil,
							map[string]interface{}{
								"id": 3,
								"attributes": map[string]interface{}{
									"SystemInformationBlocks": map[string]interface{}{
										"sib1": map[string]interface{}{
											"uac-BarringInfo": map[string]interface{}{
												"uac-BarringForCommon": []interface{}{
													nil,
													nil,
													nil,
													nil,
													nil,
													map[string]interface{}{
														"uac-barringInfoSetIndex": nil,
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
			Config{
				sliceDeepMerge: true,
				Overwrite:      false,
			},
			map[string]interface{}{
				"/ManagedElement=DU-01/GNBDUFunction=0/NRCellDU=3": map[string]interface{}{
					"body": map[string]interface{}{
						"NRCellDU": []interface{}{
							nil,
							nil,
							nil,
							map[string]interface{}{
								"id": 3,
								"attributes": map[string]interface{}{
									"SystemInformationBlocks": map[string]interface{}{
										"sib1": map[string]interface{}{
											"uac-BarringInfo": map[string]interface{}{
												"uac-BarringForCommon": []interface{}{
													nil,
													nil,
													nil,
													nil,
													nil,
													map[string]interface{}{
														"accessCategory":          5,
														"uac-barringInfoSetIndex": nil,
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
			"Sending to overwrite with null value",
		},
	}
	for i, test := range tests {
		err := Merge(&test.dst, test.src, func(c2 *Config) {
			*c2 = test.options
		})
		if err != nil {
			t.Errorf("%d: failed to Merge: %v", i+1, err)
		}
		if !reflect.DeepEqual(test.dst, test.want) {
			t.Errorf("The test %d called %s mismatch, dst: %v want: %v", i+1, test.name, test.dst, test.want)
		}
	}
}
