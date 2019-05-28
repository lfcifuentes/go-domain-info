package data

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

type App struct {
	Items[] DataSearch
}

type DataSearch struct{
	Url string
	Servers[] Server `json:"endpoints"`
	ServersChange bool
	SslGrade string
	PreviousSslGrade string
	Logo string
	Title string
	IsDown bool
}

type Data struct{
	Servers[] Server `json:"endpoints"`
	ServersChange bool
	SslGrade string
	PreviousSslGrade string
	Logo string
	Title string
	IsDown bool
}

type Server struct {
	Addres string `json:"ipAddress"`
	SslGrade string `json:"grade"`
	Country string `json:"country"`
	Owner string `json:"owner"`
	StatusMessage string `json:"statusMessage"`
}

type Page struct {
	Title string
	Logo string
	Down bool
}

func (data Data) MinSslGrade(s string,a string) string  {
	if(s==a||a==""){
		return s;
	}
	if(s=="A"){
		return a
	}else if(s=="B"){
		if(a!="A"){
			return a
		}
	}else if(s=="C"){
		if(a=="D" || a=="E" || a=="F"){
			return a
		}
	}else if(s=="D"){
		if(a=="E" || a=="F"){
			return a
		}
	}else if(s=="E"){
		if(a=="F"){
			return a
		}
	}
	return s
}

func (data Data) DataWebPage(str string) Page{
	if(str[0:4] != "http"){
		str = fmt.Sprintf("http://%s/",str)
	}
	page:= Page{Title:"",Logo:"", Down:true}

	resp, err := http.Get(str)
	if err != nil {
	}else{
		page.Down = false;
		doc, err := goquery.NewDocument(str)
		if err != nil {
			log.Fatal(err)
		}
		// start scraping page title
		doc.Find("head").Each(func(i int, s *goquery.Selection) {
			s.Find("link").Each(func(i int, selection *goquery.Selection) {
				if selection.AttrOr("rel","none") == "shortcut icon" ||
					selection.AttrOr("rel","none") == "icon"	{
					page.Logo = selection.AttrOr("href","none")
					if(page.Logo[0:2] != "//" && page.Logo[0:4] != "http"){
						page.Logo = fmt.Sprintf("%s/%s",str,page.Logo)
					}
				}
			})
			page.Title = s.Find("title").Text()
		})
	}
	defer resp.Body.Close()

	return page
}

func (data Data) Compare(d Data) (bool,string){

	if data.SslGrade != d.SslGrade {
		return true, d.SslGrade
	}

	if data.IsDown != d.IsDown {
		return true, d.SslGrade
	}

	if data.Logo != d.Logo {
		return true,d.SslGrade
	}

	if len(data.Servers) != len(d.Servers){
		return true,d.SslGrade
	}
	for i:=0; i< len(data.Servers); i++ {
		if(data.Servers[i].SslGrade != d.Servers[i].SslGrade){
			return true,d.SslGrade
		}
		if(data.Servers[i].Owner != d.Servers[i].Owner){
			return true,d.SslGrade
		}
		if(data.Servers[i].Addres != d.Servers[i].Addres){
			return true,d.SslGrade
		}
		if data.Servers[i].Country != d.Servers[i].Country {
			return true,d.SslGrade
		}
	}

	return true,d.SslGrade

}