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
		log.Print("Found ")
		siteCheckedChannel <- struct {string; bool}{ctx.URL().String(), true }				
	} else {
		siteCheckedChannel <- struct {string; bool}{ctx.URL().String(), false }
	}
	log.Println(ctx.URL());
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
	go manager(sitesStatChannel,siteSetAboveLawChannel)
	//http.Handle("/Show", &Struct{sitesStatChannel})
	http.HandleFunc("/Show", func(w http.ResponseWriter, r *http.Request) {
		showSites(w,r,sitesStatChannel)
	})
	http.HandleFunc("/AboveLaw", func(w http.ResponseWriter, r *http.Request) {
		updateAboveTheLaw(w,r, siteSetAboveLawChannel)
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
	db, err := sql.Open("mysql", "rating:LeikyarodTatyig@tcp(update.sputnikmedia.net:63306)/rating?charset=utf8")
	checkErr(err)
	// insert
	err = db.Ping();
	checkErr(err)
	rows, err := db.Query("SELECT siteid, url FROM sites where active != 6 and active != 3 and active != 0")

	for rows.Next() {
	    var uid int
	    var url string
	    err = rows.Scan(&uid, &url)
	    checkErr(err)
	    site := Site{Url:url, Id:uid, HasCounter:true, AboveLaw:false}
	    sites[url] = site	
	}
	log.Print("Finished loading sites")	
	return sites
}
func launch(sites map[string]Site){
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
}

func manager(sitesStatChannel chan map[string]Site, siteSetAboveLawChannel chan int) {
	sites := getSites()
	go launch(sites)
	for {
		select {
		case site := <- siteCheckedChannel:
			_, ok := sites[site.string]
			if ok {
				sites[site.string] = Site{sites[site.string].Url, sites[site.string].Id, site.bool, sites[site.string].AboveLaw}
			}
		case  sitesStatChannel <- sites:

		case siteId := <- siteSetAboveLawChannel:
			for key, value := range  sites{
				if value.Id == siteId {
					value.AboveLaw = true
					sites[key] = Site{sites[key].Url, sites[key].Id, true, sites[key].AboveLaw}
				}
			}
		}
	

	}
	
}

/*func (s *Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w,"<html><head></head><body>\n")
	//stat := <- siteChannel
	stat:= <-s.sitesStatChannel
	fmt.Fprint(w,"<p>")
	for key, value := range stat {
		if (value.HasCounter != true) && (value.AboveLaw != true){
    		fmt.Fprintf(w,"<a href='%s'>%s</a> :  <a href='/AboveLaw?id=%d'>Set as exception</a> <a href='//admin.bigmir.net/top/edit/%d'>Edit</a> <br/>", key, key, value.Id, value.Id)
    		}
	}
	fmt.Fprint(w,"</p></body><html>")
}*/
func showSites(w http.ResponseWriter,
	r *http.Request,
	s chan map[string]Site) {
	fmt.Fprint(w,"<html><head></head><body>\n")
	//stat := <- siteChannel
	stat:= <-s
	fmt.Fprint(w,"<p>")
	for key, value := range stat {
		if (value.HasCounter != true) && (value.AboveLaw != true){
			fmt.Fprintf(w,"<a href='%s'>%s</a> :  <a href='/AboveLaw?id=%d'>Set as exception</a> <a href='//admin.bigmir.net/top/edit/%d'>Edit</a> <br/>", key, key, value.Id, value.Id)
		}
	}
	fmt.Fprint(w,"</p></body><html>")
}
func updateAboveTheLaw(
	w http.ResponseWriter,
	r *http.Request,
	s chan int) {
	r.ParseForm()
	if val, ok := r.Form["id"]; ok {
    	i, err := strconv.Atoi(val[0])
    	if err != nil {
    		fmt.Fprint(w,"TERRIBLE ERROR")
        	return
   		}
   		s <- i
	}

	fmt.Fprint(w,"<html><head></head><body>\n")
	fmt.Fprint(w,"<p>Successfully updated<p> <br/>")
	fmt.Fprint(w,"<a href='/Show'>Back</a>")
	fmt.Fprint(w,"</body><html>")
}


