#!/bin/bash
rm -rf ProjetoWebCrawler

git clone https://github.com/vergaGabriel/ProjetoWebCrawler.git
cd ProjetoWebCrawler

apt-get install go
go run main.go