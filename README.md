
![Docker](https://i.imgur.com/6fi4Vgs.png)
===

# Offensive Dockerfiles

This repository contains a collection of **security-oriented tools** as Dockerfiles.  
This makes it easy to deploy various mission dependent tools using common cloud providers (AWS, Azure, Linode..). The containers are built using Docker. Each container is made to suit required dependencies for each tool.  

* No more buggy environments!  
* Datacenter fiber internet connection!  
* Attack remotly without the burden of running your tools on your machine
* Become a real nomad ninja
* Mix and match with the [Red Team Infractructure Guide](https://github.com/bluscreenofjeff/Red-Team-Infrastructure-Wiki) and [Red Baron](https://github.com/Coalfire-Research/Red-Baron)!



Efforts have been made to keep Dockerfiles minimal.  
**Use responsibly** for your bug bounties or pentesting needs :rainbow:

## Usage example with sqlmap:

```bash
git clone https://github.com/khast3x/Offensive-Dockerfiles.git
cd Offensive-Dockerfiles/sqlmap
docker build -t sqlmap .
docker run -it sqlmap:latest --wizard

```

## Working:

* tulpar
* nmap / NSE / Vulscan / Vulners
* sqlmap
* dcrawl
* v3n0m
* golismero
* sqliv
* datasploit
* gitminer
* Cr3dOv3r
* UFONet
* Striker
* EmailHarvester
* BruteX
* BlackWidow
* Shiva
* MemcrashedDDOS
* ctfr
* twa
* Photon
* CMSeeK
* Hash-Buster

-- To push to repo (currently are sitting as forks)
* CloudScraper
* hershell
* Merlin

## Notes:

* Adding them as I go. Don't expect production-ready images  
* Uses either python-slim or python-alpine
* Tools will show help dialog if no arguments are passed  


[:house_with_garden:](http://toto.com)

