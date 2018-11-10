

Offensive Dockerfiles
===

![Docker](https://i.imgur.com/M1KKzwO.jpg)
===


This repository contains a collection of **security-oriented tools** as Dockerfiles.  
This makes it easy to deploy various mission dependent tools using common cloud providers (AWS, Azure, Linode..). The containers are built using Docker. Each container is made to suit required dependencies for each tool.  

* No more buggy environments!  
* Datacenter fiber internet connection!  
* Attack remotly without the burden of running your tools on your machine
* :cloud: Become a real nomad ninja :cloud:
* Mix and match with the [Red Team Infractructure Guide](https://github.com/bluscreenofjeff/Red-Team-Infrastructure-Wiki) and [Red Baron](https://github.com/Coalfire-Research/Red-Baron)!



Efforts have been made to keep Dockerfiles minimal.  

----
:star: **Use responsibly** for your bug bounties or pentesting needs 

---

## Usage example with sqlmap:

```bash
git clone https://github.com/khast3x/Offensive-Dockerfiles.git
cd Offensive-Dockerfiles/sqlmap
docker build -t sqlmap .
docker run -it sqlmap:latest --wizard

```

## Working:

| Name 	| Description 	|
|:----------------------------------------------------------------------------------------------------------------------------------------:	|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------	|
| [tulpar](https://github.com/anilbaranyelken/tulpar) 	| Web Vulnerability Scanner 	|
| [nmap](https://nmap.org) + [Vulscan](https://github.com/scipag/vulscan) + [Vulners](https://github.com/vulnersCom/nmap-vulners) scripts  	| Latest Nmap Scripting Engine (NSE) modules, as well as the Vulscan NSE script and the vulners API to NSE script. 	|
| [sqlmap](https://github.com/sqlmapproject/sqlmap) 	| Automatic SQL injection and database takeover tool 	|
| [dcrawl](https://github.com/kgretzky/dcrawl) 	| Simple, but smart, multi-threaded web crawler for randomly gathering huge lists of unique domain names. 	|
| [V3n0m Scanner](https://github.com/v3n0m-Scanner/V3n0M-Scanner) 	| Popular Pentesting scanner in Python3.6 for SQLi/XSS/LFI/RFI and other Vulns 	|
| [golismero](https://github.com/golismero/golismero) 	| The Web Knife 	|
| [sqliv](https://github.com/Hadesy2k/sqliv) 	| massive SQL injection vulnerability scanner 	|
| [datasploit](https://datasploit.github.io/datasploit) 	| Performs OSINT on a domain / email / username / phone 	|
| [gitminer](http://github.com/danilovazb/gitminer) 	| Tool for advanced mining for content on Github 	|
| [Cr3d0v3r](https://github.com/D4Vinci/Cr3dOv3r) 	| Know the dangers of credential reuse attacks 	|
| [UFONet](https://github.com/epsylon/ufonet) 	| UFONet abuses OSI Layer 7-HTTP to create/manage 'zombies' and to conduct different attacks using; GET/POST, multithreading, proxies, origin spoofing methods, cache evasion techniques, etc. 	|
| [Striker](https://github.com/UltimateHackers/Striker) 	| Striker is an offensive information and vulnerability scanner 	|
| [emailHarvester](https://github.com/maldevel/EmailHarvester) 	| Email addresses harvester 	|
| [BruteX](https://github.com/1N3/BruteX) 	| Automatically brute force all services running on a target 	|
| [BlackWidow](https://github.com/1N3/BlackWidow) 	| A Python based web application scanner to gather OSINT and fuzz for OWASP vulnerabilities on a target website 	|
| [Shiva](https://github.com/UltimateHackers/Shiva) 	|  Improved DOS exploit for wordpress websites (CVE-2018-6389) 	|
| [Memcrashed](https://github.com/649/Memcrashed-DDoS-Exploit) 	| This tool allows you to send forged UDP packets to Memcached servers obtained from Shodan.io 	|
| [ctfr](https://github.com/UnaPibaGeek/ctfr.git) 	| Domain enumeration, it just abuses of Certificate Transparency logs 	|
| [twa](https://github.com/woodruffw/twa) 	| A **t**iny **w**eb **a**uditor with strong opinions 	|
| [Photon](https://github.com/s0md3v/Photon) 	| Incredibly fast crawler designed for OSINT 	|
| [CMSeek](https://github.com/Tuhinshubhra/CMSeeK) 	| CMS Detection and Exploitation suite - Scan WordPress, Joomla, Drupal and 130 other CMSs 	|
| [HashBuster](https://github.com/s0md3v/Hash-Buster) 	| Crack hashes in seconds 	|
|  	|  	|
|  	|  	|


#### To push to repo (currently are sitting as forks)
* CloudScraper
* hershell
* Merlin

## Notes:

* Adding them as I go. Don't expect production-ready images  
* Uses either python-slim or python-alpine
* Tools will show help dialog if no arguments are passed  


