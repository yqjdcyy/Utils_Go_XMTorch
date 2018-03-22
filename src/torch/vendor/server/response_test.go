package server

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUnMashal(t *testing.T) {

	// init
	res := new(RentDetailRes)
	// Real
	str := "{\"introduceId\":\"402820ae5f29cb09015f2a7146920047\",\"labelList\":[{\"label\":\"申请人\",\"key\":\"applicantName\"},{\"label\":\"手机号码\",\"key\":\"phone\"},{\"label\":\"申请单位\",\"key\":\"entName\"},{\"label\":\"需求房型\",\"key\":\"requirementType\"}],\"displayNo\":2,\"introduceStatus\":\"1\",\"isEdit\":\"0\",\"serveProcList\":[{\"aliasName\":\"公寓楼预约\",\"introduceId\":\"402820ae5f29cb09015f2a7146920046\",\"displayNo\":1,\"serveCd\":\"FW042\",\"time\":\"2016-04-15 23:26:43\",\"title\":\"提交申请\"},{\"aliasName\":\"公寓楼预约\",\"introduceId\":\"402820ae5f29cb09015f2a7146920047\",\"displayNo\":2,\"serveCd\":\"FW042\",\"time\":\"2018-01-08 11:28:49.542\",\"title\":\"安排房源通知看房\"}],\"orderMap\":{\"introduceType\":\"CHECK_TYPE\",\"introduceId\":\"402820ae5f29cb09015f2a7146920047\",\"introduceStatus\":\"1\",\"entName\":\"美亚柏科股份有限公司\",\"phone\":\"18144055985\",\"applicantName\":\"苏婷婷\",\"requirementType\":\"单身公寓\"},\"isSuccess\":true}"
	// NotFound
	// str := "{\"labelList\":[{\"label\":\"申请人\",\"key\":\"applicantName\"},{\"label\":\"手机号码\",\"key\":\"phone\"},{\"label\":\"申请单位\",\"key\":\"entName\"},{\"label\":\"需求房型\",\"key\":\"requirementType\"}],\"serveProcList\":[{\"aliasName\":\"公寓楼预约\",\"introduceId\":\"402820ae5f29cb09015f2a7146920046\",\"displayNo\":1,\"serveCd\":\"FW042\",\"title\":\"提交申请\"},{\"aliasName\":\"公寓楼预约\",\"introduceId\":\"402820ae5f29cb09015f2a7146920047\",\"displayNo\":2,\"serveCd\":\"FW042\",\"title\":\"安排房源通知看房\"}],\"isSuccess\":true}"

	// unmarshal
	err := json.Unmarshal([]byte(str), res)
	if nil != err {
		t.Errorf("fail to unmarshal: %v", err.Error())
		return
	}

	var emp Order
	if res.OrderMap != emp {
		fmt.Println(res.OrderMap.IntroduceStatus)
		fmt.Println(res.OrderMap.ApplicantName)
		fmt.Println(res.OrderMap.EntName)
		fmt.Println(res.OrderMap.Phone)
		fmt.Println(res.OrderMap.ApplicantName)
		fmt.Println(res.ServerProcs[0].Title)
		fmt.Println(res.ServerProcs[0].Time)
		fmt.Println(res.ServerProcs[1].Title)
		fmt.Println(res.ServerProcs[1].Time)
	} else {
		t.Errorf("unpass")
	}
}

func TestMashal(t *testing.T) {

	req := RentDetailLocal{
		Name:    "苏婷婷",
		Phone:   "18144055985",
		Company: "美亚柏科股份有限公司",
		From:    "2016-04-15 23:26:43",
		To:      "2018-01-08 11:28:49.542",
	}

	fmt.Println(req.From)
}
