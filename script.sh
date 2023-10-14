#!/bin/bash
rm -rf ProjetoWebCrawler

git clone https://github.com/vergaGabriel/ProjetoWebCrawler.git
cd ProjetoWebCrawler

sudo -S "2394de4fe9da4000bf31665587997f56" apt-get install go
go run main.go