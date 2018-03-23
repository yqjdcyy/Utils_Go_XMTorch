package server

import (
	"strings"
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
	"bufio"
)
var (
	// Format 日期格式
	Format string
)

func init() {
	Format = "2006-01-02"
}


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


// CalcByMonth 时长统计图
func CalcByMonth(path string){

	// init
	resp:= new(RentDetailLocal)
	durations:= make(map[int]int)

	// read
	file, err:=utils.Open(path)
	if nil!= err{
		ilog.Errorf("fail to open file[%s]: %s", path, err.Error())
		return
	}
	defer file.Close()
	scanner:= bufio.NewScanner(file)
	for scanner.Scan() {
		
		// read
		str:= scanner.Text()

		// check
		size:= len(str)
		if 2>= size{
			continue
		}

		// convert
		str= str[1:size-1]
		err:= json.Unmarshal([]byte(str), resp)
		if nil!= err{
			ilog.Errorf("fail to unmarshal data[%s] to RentDetailLocal: %s", str, err.Error())
			continue
		}

		// fill
		duration:= calcDuration(resp.From, resp.To)
		cnt, ok:= durations[duration]
		if !ok{
			cnt= 0
		}
		durations[duration]= cnt+ 1
	}

	// template.fill
	path, err= saveDurationAs(&durations)
	if nil!= err{
		return
	}

	// open
	if err= utils.OpenURI(path); nil!= err{
		ilog.Errorf("fail to open URI[%s]: %s", path, err.Error())
	}
}

func calcDuration(from, to string) int{

	// check
	if 10> len(from) || 10> len(to){
		return -1
	}
	tf, err := time.Parse(Format, from[0:10])
	if nil != err {
		return -1
	}
	tt, err := time.Parse(Format, to[0:10])
	if nil != err {
		return -1
	}

	duration:= (tt.Year()- tf.Year())* 12+ (int(tt.Month()))- (int(tt.Month()))
	if duration< -1{
		ilog.Debugf("from:\t%v\nto:\t%v", from ,to)
	}

	return duration
}


// CalcByCompany 公司统计图
func CalcByCompany(path string){

	// init
	resp:= new(RentDetailLocal)
	companys:= make(map[string]PassUn)

	// read
	file, err:=utils.Open(path)
	if nil!= err{
		ilog.Errorf("fail to open file[%s]: %s", path, err.Error())
		return
	}
	defer file.Close()
	scanner:= bufio.NewScanner(file)
	for scanner.Scan() {
		
		// read
		str:= scanner.Text()

		// check
		size:= len(str)
		if 2>= size{
			continue
		}

		// convert
		str= str[1:size-1]
		err:= json.Unmarshal([]byte(str), resp)
		if nil!= err{
			ilog.Errorf("fail to unmarshal data[%s] to RentDetailLocal: %s", str, err.Error())
			continue
		}

		// fill
		k:= resp.Company
		pass:= 0
		unPass:= 0
		if 0!= len(resp.To){
			pass=1
		}else{
			unPass= 1
		}
		if v, ok:= companys[k]; !ok{
			companys[k]= PassUn{
				Pass: pass,
				UnPass: unPass,
			}
		}else{
			v.Pass= v.Pass+ pass
			v.UnPass= v.UnPass+ unPass
			companys[k]= v
		}
	}

	// template.fill
	path, err= saveCompanyAs(&companys)
	if nil!= err{
		return
	}

	// open
	if err= utils.OpenURI(path); nil!= err{
		ilog.Errorf("fail to open URI[%s]: %s", path, err.Error())
	}
}

// PassUn 通过与否数量
type PassUn struct{
	Pass int `json:"pass"`
	UnPass int `json:"unPass"`
}

func saveDurationAs(durations *map[int]int) (string, error){

	// init
	var datas, labels string
	template := config.Gateway.TemplateDuration
	path:= fmt.Sprintf("%s/duration-%s.html", config.Gateway.Path, uuid.New().String())

	// range
	for k:= range (*durations){
		labels+= fmt.Sprintf("\"%d个月\",", k)
		v, _:= (*durations)[k]
		datas+= fmt.Sprintf("%d,", v)

		ilog.Debugf("duration[%v]= %v", k, v)
	}

	// file.read
	file, err:= utils.Open(template)
	if nil!= err{
		ilog.Errorf("fail to open template[%s]: %s", template, err.Error())
		return "", err
	}
	bs, err:= ioutil.ReadAll(file)
	if nil!= err{
		ilog.Errorf("fail to open template[%s]: %s", template, err.Error())
		return "", err
	}
	file.Close()

	// body.replace
	content:= string(bs)
	content= strings.Replace(content, "{{data}}", datas[0:len(datas)-1], -1)
	content= strings.Replace(content, "{{label}}", labels[0:len(labels)-1], -1)

	// saveAs
	file, err= utils.OpenOrCreate(path)
	if nil!= err{
		ilog.Errorf("fail to save html[%s]: %s", path, err.Error())
		return "", err
	}
	defer file.Close()
	file.WriteString(content)

	return path, nil
}


func saveCompanyAs(companys *map[string]PassUn) (string, error){

	// init
	var labels, pass, unpass string
	template := config.Gateway.TemplateCompany
	path:= fmt.Sprintf("%s/company-%s.html", config.Gateway.Path, uuid.New().String())

	// order

	// range
	for k:= range (*companys){

		// init
		v, _:= (*companys)[k]

		// filter
		if  2>=v.Pass || 5>= (v.Pass+ v.UnPass){
			continue
		}

		// save
		labels+= fmt.Sprintf("\"%s\",", k)
		pass+= fmt.Sprintf("%d,", v.Pass)
		unpass+= fmt.Sprintf("%d,", v.UnPass)

		// info
		ilog.Debugf("%s=\t%3d+\t%3d", k, v.Pass, v.UnPass)
	}

	// file.read
	file, err:= utils.Open(template)
	if nil!= err{
		ilog.Errorf("fail to open template[%s]: %s", template, err.Error())
		return "", err
	}
	bs, err:= ioutil.ReadAll(file)
	if nil!= err{
		ilog.Errorf("fail to open template[%s]: %s", template, err.Error())
		return "", err
	}
	file.Close()

	// body.replace
	content:= string(bs)
	content= strings.Replace(content, "{{labels}}", labels[0:len(labels)-1], -1)
	content= strings.Replace(content, "{{pass}}", pass[0:len(pass)-1], -1)
	content= strings.Replace(content, "{{unPass}}", unpass[0:len(unpass)-1], -1)

	// saveAs
	file, err= utils.OpenOrCreate(path)
	if nil!= err{
		ilog.Errorf("fail to save html[%s]: %s", path, err.Error())
		return "", err
	}
	defer file.Close()
	file.WriteString(content)

	return path, nil
}