package search

import (
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/go-chi/render"
	"io/ioutil"
	"net/http"
	"os/exec"
	_ "reflect"
	"strings"
	"testTruora/data"
	"testTruora/db"
)

func SearchInformation(search Search) string{
	// creo la peticion
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.ssllabs.com/api/v3/analyze?host=%s",search.Url), nil)

	req.Header.Add("User-Agent", `Go`)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	// optendo la informacion de la peticion
 	var data data.Data;
	errorjson := json.Unmarshal([]byte(string(body)), &data)
	if(errorjson != nil){
		fmt.Printf("Error :%v",errorjson)
	}
	sslGrade :=""
	// agrego la información que falta de los servidores
	for i:=0; i< len(data.Servers); i++ {
		out,errT := exec.Command("whois",data.Servers[i].Addres).Output()
		if(errT != nil){
			fmt.Print(errT)
		}
		outString := string(out)
		splittedString := strings.Split(outString, "\n")
		for e:=0; e< len(splittedString); e++ {
			s:=""
			o:=""
			if strings.Contains(splittedString[e], "Country:") {
				s= strings.ReplaceAll(splittedString[e], " ", "")
				s= strings.ReplaceAll(s, "Country:", "")
				data.Servers[i].Country = s
			}
			if strings.Contains(splittedString[e], "OrgName:") {
				o= strings.ReplaceAll(splittedString[e], "OrgName: ", "")
				data.Servers[i].Owner = o
			}
			// verificar el ssl menor
			sslGrade = data.MinSslGrade(data.Servers[i].SslGrade,sslGrade)
		}
	}
	// grado ssl menor
	data.SslGrade = sslGrade

	// datos del logo y el titulo
	infoPage := data.DataWebPage(search.Url)
	data.Logo = infoPage.Logo
	data.Title = infoPage.Title
	data.IsDown = infoPage.Down

	b, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return ""
	}
	// validar el registro de la db
	if(db.CheckData(search.Url)){
		// si no existe lo creo
		db.InsertSearch(string(b),search.Url)
	}else{
		data.ServersChange,data.PreviousSslGrade = db.ChangesData(data,search.Url);
		if(data.ServersChange){
			b, err := json.Marshal(data)
			if err != nil {
				fmt.Printf("Error: %s", err)
				return ""
			}
			db.UpdateServer(string(b),search.Url)
		}
	}
	return string(b)
}

type SearchRequest struct {
	*Search
}

func (s SearchRequest) Bind(r *http.Request) error {
	// url is nil if no url fields are sent in the request. Return an
	// error to avoid a nil pointer dereference.
	if s.Search.Url == "" {
		return errors.New("missing required Url fields.")
	}

	// url is nil if no Userpayload fields are sent in the request. In this app
	// this won't cause a panic, but checks in this Bind method may be required if
	// a.User or futher nested fields like a.User.Name are accessed elsewhere.

	// just a post-process after a decode
	s.Search.Url = strings.ToLower(s.Search.Url) // as an example, we down-case
	return nil
}

type Search struct {
	Url string `json:"url"`
}

func GetData() string  {
	// obj con los datos
	app := db.GetData()

	b, err := json.Marshal(app)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return ""
	}

	return string(b)
}