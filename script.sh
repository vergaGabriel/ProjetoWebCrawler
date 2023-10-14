#!/bin/bash
rm -rf ProjetoWebCrawler
sudo passwd jenkins
git clone https://github.com/vergaGabriel/ProjetoWebCrawler.git
cd ProjetoWebCrawler

sudo -S apt-get install go
go run main.go