package main

import (
	"net/http"
	"regexp"
	"time"
	"log"
	"io/ioutil"
	"fmt"
	"strconv"
	"flag"

	r "crawler/regexps"

	"github.com/PuerkitoBio/gocrawl"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
    "database/sql"
)

// Only enqueue the root and paths beginning with an "a"
//var rxOk = regexp.MustCompile(`http://duckduckgo\.com(/a.*)?$`)
//var rxOk = regexp.MustCompile(`http://(.+\.)*bigmir\.net(/.*)?$`)
//var rxOk = regexp.MustCompile(`http://bigbenrating\.com(/.*)?$`)
var rxOk = regexp.MustCompile(`http://korrespondent\.net(/.*)?&`)

type ExampleExtender struct {
	gocrawl.DefaultExtender // Will use the default implementation of all but Visit and Filter
}

var siteCheckedChannel = make(chan struct{string; bool})


// Override Visit for our need.
func (x *ExampleExtender) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {

	// Return nil and true - let gocrawl find the links
	htmlData, err := ioutil.ReadAll(res.Body)
 	if err != nil {
 		log.Println(err)
 		//os.Exit(1)
 	}
	//log.Println(string(htmlData))
	//var flag = r.CounterRx.MatchString(string(htmlData))
	var flag = r.HasCounter(string(htmlData))
	if (flag == true){
		//log.Print("Found ")
		siteCheckedChannel <- struct {string; bool}{ctx.URL().String(), true }				
	} else {
		siteCheckedChannel <- struct {string; bool}{ctx.URL().String(), false }
	}
	//log.Println(ctx.URL());
	return nil, true
}

// Override Filter for our need.
func (x *ExampleExtender) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	return !isVisited /*&& rxOk.MatchString(ctx.NormalizedURL().String())*/
}
type Site struct{
	Url string
	Id int
	HasCounter bool
	AboveLaw bool
	Hidden bool
	Users int
}
func (s Site) setHavingCounter(flag bool){
	s.HasCounter = flag
}
/*type Struct struct{
	sitesStatChannel chan map[string]Site
}*/
type SomeAnotherStruct struct{
	b int
}
func main() {
	ip := flag.String("ip", "192.168.0.31", "ip")
	port := flag.String("port", "4000", "port")
	flag.Parse()
	sitesStatChannel := make(chan map[string]Site)
	siteSetAboveLawChannel := make(chan int)
	siteIgnoreChannel := make(chan int)
	go manager(sitesStatChannel,siteSetAboveLawChannel,siteIgnoreChannel)
	//http.Handle("/Show", &Struct{sitesStatChannel})
	http.HandleFunc("/Show", func(w http.ResponseWriter, r *http.Request) {
		showSites(w,r,sitesStatChannel)
	})
	http.HandleFunc("/AboveLaw", func(w http.ResponseWriter, r *http.Request) {
		updateAboveTheLaw(w,r, siteSetAboveLawChannel)
	})
	http.HandleFunc("/Ignore", func(w http.ResponseWriter, r *http.Request) {
		ignoreSite(w,r, siteIgnoreChannel)
	})
	addr := fmt.Sprintf("%s:%s", *ip, *port)
	log.Println("Started at",addr)
	err := http.ListenAndServe(addr, nil )
	if err != nil {
		log.Fatal(err)
	}
}
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func getSites() map[string]Site {
	sites := make(map[string]Site)
	db, err := sql.Open("mysql", connectString)
	checkErr(err)
	// insert
	err = db.Ping();
	checkErr(err)
	rows, err := db.Query("SELECT siteid, users_30 as users, url FROM sites where active != 6 and active != 3 and active != 0")

	for rows.Next() {
		var uid int
		var users int
		var url string

		err = rows.Scan(&uid,&users, &url)
		checkErr(err)
		site := Site{Url:url, Id:uid, HasCounter:true, AboveLaw:false, Hidden:false, Users:users}
		sites[url] = site
	}
	log.Print("Finished loading sites")	
	return sites
}
func launch(sites map[string]Site, finish chan bool){
	opts := gocrawl.NewOptions(new(ExampleExtender))
	opts.RobotUserAgent = "Example"
	opts.UserAgent = "Mozilla/5.0 (compatible; Example/1.0; +http://example.com)"
	opts.CrawlDelay = 1 * time.Second
	opts.MaxVisits = 1
	opts.SameHostOnly = false;
	crawler := gocrawl.NewCrawlerWithOptions(opts)
	for url, _ := range sites {
		crawler.Run(url)
	}
	log.Print("Finished crawling")
	finish <- true
}

func manager(sitesStatChannel chan map[string]Site, siteSetAboveLawChannel chan int, siteIgnoreChannel chan int) {
	sites := getSites()
	finish := make(chan bool)
	go launch(sites, finish)
	for {
		select {
		case site := <- siteCheckedChannel:
			_, ok := sites[site.string]
			if ok {
				sites[site.string] = Site{sites[site.string].Url, sites[site.string].Id, site.bool, sites[site.string].AboveLaw, sites[site.string].Hidden, sites[site.string].Users}
			}
		case  sitesStatChannel <- sites:

		case siteId := <- siteSetAboveLawChannel:
			for key, value := range  sites{
				if value.Id == siteId {
					value.AboveLaw = true
					sites[key] = Site{sites[key].Url, sites[key].Id, true, sites[key].AboveLaw, sites[key].Hidden, sites[key].Users}
				}
			}
		case siteId := <- siteIgnoreChannel:
			for key, value := range  sites{
				if value.Id == siteId {
					value.Hidden = true
					sites[key] = Site{sites[key].Url, sites[key].Id, true, sites[key].AboveLaw, sites[key]. Hidden, sites[key].Users}
				}
			}
		case <- finish:
			updateSites:= getSites()
			for url, site := range updateSites {
				if val, ok := sites[url]; ok {
					if val.AboveLaw{
						continue
					}
				}
				sites[url] = site
			}
			go launch(sites, finish)
		}
	

	}
	
}

func showSites(w http.ResponseWriter,
	r *http.Request,
	s chan map[string]Site) {

	r.ParseForm()
	page := 0
	if val, ok := r.Form["p"]; ok {
		page, _ = strconv.Atoi(val[0])
		page = page - 1
	}

	fmt.Fprint(w,"<html><head></head>")
	fmt.Fprint(w,`<body bgcolor="white" leftmargin="5" topmargin="5" marginwidth="5" marginheight="5">
	<table width="100%" border="0" align="center" cellpadding="3" cellspacing="1">
	<tbody><tr>
	<td><table width="760" border="0" align="center" cellpadding="3" cellspacing="1">
	<tbody><tr>
	<td height="20" colspan="6" align="center" bgcolor="#FFECB0" class="small">
		новые сайты</td>
	</tr>
	<tr bgcolor="#DDDDDD">
	<td height="20" align="center" class="small">ID</td>
	<td align="center" class="small">Адресс</td>
	<td align="center" class="small">Пользователи</td>
	<td align="center" class="small">Спрятать</td>
	<td align="center" class="small">Исключения</td>
	</tr>`)
	//stat := <- siteChannel
	stat:= <-s
	fmt.Fprint(w,"<p>")
	counter:=0
	for key, value := range stat {

		if (value.HasCounter != true) && (value.AboveLaw != true) && (value.Hidden != true) && (counter < (page + 1) * 10){
			counter = counter + 1
			//fmt.Fprintf(w,"<a href='%s'>%s</a> users  %d   :  <a href='/AboveLaw?id=%d'>Set as exception</a> <a href='/Ignore?id=%d'>Hide</a> <a href='//admin.bigmir.net/top/edit/%d'>Edit</a> <br/>", key, key,value.Users, value.Id, value.Id, value.Id)
			if (counter > page * 10){
				fmt.Fprintf(w,"<tr><td><a href='//admin.bigmir.net/top/edit/%d'>%d</a></td><td><a href='%s'>%s</a></td><td>%d</td><td><a href='/Ignore?id=%d'>Hide</a></td><td><a href='/AboveLaw?id=%d'>Except</a></td></tr>",value.Id, value.Id, key, key, value.Users, value.Id,value.Id)
			}
		}
	}
	fmt.Fprint(w,`<tr>
	<td height="30" colspan="6">

	<div class="pagingBlock"><span>Страницы:</span> `)
	for i:=0;  i < 11; i++{
		p := page - 5 + i
		if (p > 0) && (p < len(stat) / 10 + 1){
			if (p-1 == page) {
				fmt.Fprintf(w,"<span id=\"current\">%d</span>  ",p)
			} else {
				fmt.Fprintf(w,"<a href=\"Show?p=%d\">%d</a>  ",p,p)
			}
		}
	}
	fmt.Fprint(w,"</td></tr></table></p></body><html>")
}
func updateAboveTheLaw(
	w http.ResponseWriter,
	r *http.Request,
	s chan int) {
	r.ParseForm()
	if val, ok := r.Form["id"]; ok {
		i, err := strconv.Atoi(val[0])
		if err != nil {
			fmt.Fprint(w,"Not a number as param to setAboveLaw")
			return
			}
			s <- i
	}

	fmt.Fprint(w,"<html><head></head><body>\n")
	fmt.Fprint(w,"<p>Successfully updated<p> <br/>")
	fmt.Fprint(w,"<a href='/Show'>Back</a>")
	fmt.Fprint(w,"</body><html>")
}

func ignoreSite(
	w http.ResponseWriter,
	r *http.Request,
	s chan int) {
	r.ParseForm()
	if val, ok := r.Form["id"]; ok {
		i, err := strconv.Atoi(val[0])
		if err != nil {
			fmt.Fprint(w,"Not a number as param to ignore")
			return
		}
		s <- i
	}

	fmt.Fprint(w,"<html><head></head><body>\n")
	fmt.Fprint(w,"<p>Successfully ignored<p> <br/>")
	fmt.Fprint(w,"<a href='/Show'>Back</a>")
	fmt.Fprint(w,"</body><html>")
}