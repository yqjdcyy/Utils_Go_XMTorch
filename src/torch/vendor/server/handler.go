package server

import (
	"bitbucket.org/ansenwork/ilog"
	"config"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
	"utils"
	"io/ioutil"
	"strconv"
)

// Circulate 递归查询
func Circulate(from, to time.Time, url string) {

	// init
	handler := new(utils.Handler)
	handler.SetPath(fmt.Sprintf("%s/%s.log", config.Gateway.Path, uuid.New().String()))
	defer handler.Close()

	// for
	for ; from.Before(to); from = from.AddDate(0, 0, 1) {

		ilog.Infof("%4d%2d%2d", from.Year(), int(from.Month()), from.Day())
		for i := 1; i <= 10; i++ {

			reqURL := fmt.Sprintf(
				"%s%4d%02d%02d%05d",
				url,
				from.Year(),
				int(from.Month()),
				from.Day(),
				i)
			str, err :=request(reqURL)
			if nil != err {
				ilog.Errorf("fail to request for url[%s]", reqURL)
				break
			}
			err = handle(handler, str)
			if nil != err {
				ilog.Errorf("fail to save data: %s", err.Error())
				break
			}
		}
	}
}

func mock(url string)(string, error){

	str:= url[len(url)-5:]
	idx,_:= strconv.ParseInt(str, 10,32)
	if 0== idx%2{
		return "{\"labelList\":[{\"label\":\"申请人\",\"key\":\"applicantName\"},{\"label\":\"手机号码\",\"key\":\"phone\"},{\"label\":\"申请单位\",\"key\":\"entName\"},{\"label\":\"需求房型\",\"key\":\"requirementType\"}],\"serveProcList\":[{\"aliasName\":\"公寓楼预约\",\"introduceId\":\"402820ae5f29cb09015f2a7146920046\",\"displayNo\":1,\"serveCd\":\"FW042\",\"title\":\"提交申请\"},{\"aliasName\":\"公寓楼预约\",\"introduceId\":\"402820ae5f29cb09015f2a7146920047\",\"displayNo\":2,\"serveCd\":\"FW042\",\"title\":\"安排房源通知看房\"}],\"isSuccess\":true}", nil
	}
	return "{\"introduceId\":\"402820ae5f29cb09015f2a7146920047\",\"labelList\":[{\"label\":\"申请人\",\"key\":\"applicantName\"},{\"label\":\"手机号码\",\"key\":\"phone\"},{\"label\":\"申请单位\",\"key\":\"entName\"},{\"label\":\"需求房型\",\"key\":\"requirementType\"}],\"displayNo\":2,\"introduceStatus\":\"1\",\"isEdit\":\"0\",\"serveProcList\":[{\"aliasName\":\"公寓楼预约\",\"introduceId\":\"402820ae5f29cb09015f2a7146920046\",\"displayNo\":1,\"serveCd\":\"FW042\",\"time\":\"2016-04-15 23:26:43\",\"title\":\"提交申请\"},{\"aliasName\":\"公寓楼预约\",\"introduceId\":\"402820ae5f29cb09015f2a7146920047\",\"displayNo\":2,\"serveCd\":\"FW042\",\"time\":\"2018-01-08 11:28:49.542\",\"title\":\"安排房源通知看房\"}],\"orderMap\":{\"introduceType\":\"CHECK_TYPE\",\"introduceId\":\"402820ae5f29cb09015f2a7146920047\",\"introduceStatus\":\"1\",\"entName\":\"美亚柏科股份有限公司\",\"phone\":\"18144055985\",\"applicantName\":\"苏婷婷\",\"requirementType\":\"单身公寓\"},\"isSuccess\":true}" ,nil
}

// request
func request(url string) (string, error) {

	// Sleep
	time.Sleep(time.Duration(config.Gateway.Interval)* time.Second)

	// init
	ilog.Infof("request(%v)", url)

	// POST
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	if nil != err {
		ilog.Errorf("fail to post url[%v]", url)
		return "", err
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("cookie", config.Gateway.Cookie)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	// POST.Response
	respBody, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		ilog.Errorf("fail to get response from post to url[%v]", url)
		return "", err
	}
	return string(respBody), nil
}

// string.handle
func handle(handler *utils.Handler, content string) error {

	res := new(RentDetailRes)
	var emp Order

	// unmashal
	err := json.Unmarshal([]byte(content), res)
	if nil != err {
		ilog.Errorf("fail to unmarshal: %v", err.Error())
		return err
	}
	if res.OrderMap == emp {
		err := fmt.Errorf("detail not found")
		// ilog.Errorf(err.Error())
		return err
	}

	// append
	str:= convert(res)
	handler.Append("%s\n", str)

	return nil
}

func convert(res *RentDetailRes) string {

	// init
	var from, to string
	if len(res.ServerProcs) == 2 {
		to = res.ServerProcs[1].Time
	}
	from = res.ServerProcs[0].Time

	// ilog.Debugf("req= %v\t%v\t%v\t%v\t%v\t%v",
	// 	res.OrderMap.ApplicantName,
	// 	res.OrderMap.Phone,
	// 	res.OrderMap.EntName,
	// 	from,
	// 	to,
	// 	res)

	// fill
	req := RentDetailLocal{
		Name:    res.OrderMap.ApplicantName,
		Phone:   res.OrderMap.Phone,
		Company: res.OrderMap.EntName,
		From:    from,
		To:      to,
	}

	// marshal
	bs, err := json.Marshal(req)
	if nil != err {
		ilog.Errorf("fail to convert RentDetailLocal to string: %s", err.Error())
		return ""
	}

	return string(bs)
}
