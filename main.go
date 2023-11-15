package main

import(
        "fmt" //formatted I/O
        "github.com/gocolly/colly" //scraping framework
        "strconv"
        "time"
        "log"

        "database/sql"
        _ "github.com/go-sql-driver/mysql"
)
 
 
func main(){
        var cont int = 0

        db, err := sql.Open("mysql", "adminwebcrawler:Q51|3X{~Xy2$k6|+=Qs@tcp(database-webcrawler-univem.cwbkeok0cciy.us-east-2.rds.amazonaws.com:3306)/db_webcrawler")
        if err != nil {
                panic(err.Error())
            }
        

        c := colly.NewCollector(colly.AllowedDomains("www.amazon.com.br"))

        c.OnRequest(func(r *colly.Request){
                fmt.Println("Link of the page:", r.URL)
        })
 
        c.OnHTML("div.s-result-list.s-search-results.sg-row", func(h*colly.HTMLElement){
                h.ForEach("div.a-section.a-spacing-base", func(_ int, h*colly.HTMLElement){
                        time.Sleep(30 * time.Second) 
                        cont++
                        data := time.Now()
                        data_format := data.Format(time.RFC3339)
                        var name string
                        name = h.ChildText("span.a-size-base-plus.a-color-base.a-text-normal")
                        var stars string
                        stars = h.ChildText("span.a-icon-alt")
                        var price string
                        price = h.ChildText("span.a-price-whole")
 
                        _, err2 := db.Query("INSERT INTO db_webcrawler.TB_RAW_WEBCRAWLER (PRODUTO, VALOR, NOTA, DATA_EXTRACAO) VALUES ('"+name+"','"+price+"','"+stars+"','"+data_format+"')")

                        if err2 != nil {
                                log.Fatal(err2)
                        }
                })
        })
 
        for i:=1; i<2; i++{
                num := strconv.Itoa(i)
                fmt.Println(num)
                fmt.Println("https://www.amazon.com.br/s?k=whey&page="+num)
                c.Visit("https://www.amazon.com.br/s?k=whey&page="+num)
        }

        defer db.Close()
        fmt.Println("Foram inseridas um total de", cont, "linhas!")
}

