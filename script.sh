#!/bin/bash
rm -rf ProjetoWebCrawler

git clone https://github.com/vergaGabriel/ProjetoWebCrawler.git
cd ProjetoWebCrawler

go get github.com/gocolly/colly
go get github.com/go-sql-driver/mysql

go run main.go