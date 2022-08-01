// Alipay.com Inc. Copyright (c) 2004-2019 All Rights Reserved.
// @author bofeng.lt
// @version : Json2Model, v 0.1 2022/07/31 10:42 PM bofeng.lt Exp $$
// @Description:

package _731_json_to_model

type ReqBody struct {
	Data []struct {
		Sence       string `json:"sence"`
		CurMetaData struct {
			Replicates int     `json:"replicates"`
			CpuUtil    float64 `json:"cpuUtil"`
		} `json:"curMetaData"`
		BaseLine struct {
			CoreUtilThreshold float64 `json:"coreUtilThreshold"`
			MaxReplicates     int     `json:"maxReplicates"`
			MinReplicates     int     `json:"minReplicates"`
			MinSupportQps     int     `json:"minSupportQps"`
			Transfer          int     `json:"transfer"`
			Interval          float64 `json:"interval"`
			XOrigin           int     `json:"x_origin"`
		} `json:"baseLine"`
		PredictSeries []int `json:"predictSeries"`
		Tab3          struct {
			Sence       string `json:"sence"`
			CurMetaData struct {
				Replicates int     `json:"replicates"`
				CpuUtil    float64 `json:"cpuUtil"`
			} `json:"curMetaData"`
			BaseLine struct {
				CoreUtilThreshold float64 `json:"coreUtilThreshold"`
				MaxReplicates     int     `json:"maxReplicates"`
				MinReplicates     int     `json:"minReplicates"`
				MinSupportQps     int     `json:"minSupportQps"`
				Transfer          int     `json:"transfer"`
				Interval          float64 `json:"interval"`
				XOrigin           int     `json:"x_origin"`
			} `json:"baseLine"`
			PredictSeries []int `json:"predictSeries"`
		} `json:"tab3"`
	} `json:"data"`
}
