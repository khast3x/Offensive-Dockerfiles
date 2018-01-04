# Offensive Dockerfiles
![Docker](https://koalalorenzo2014.files.wordpress.com/2015/04/docker-icon-bw.png?w=300&h=212)

# Abstract
Alpine or slim Dockerfiles of various information security offensive tools, making them easy to deploy in hosted cloud environments.  
It becomes much easier to deploy and use these tools with mission specific assets such as:

* Hosting environment
* SOCKS proxy using docker built-in networking
	* Transparent Tor overlay proxy
	* Example given in ``docker-compose`` format with attached minimal Tor SOCKS proxy
	* Can also be deployed using docker ``bridge networking`` for complete transparent proxy of services
	* Using mission specific domain names for implant reverse callbacks

#### Working:
* tulpar
* nmap / NSE / Vulscan
* sqlmap
* dcrawl
* v3n0m
* golismero
* sqliv
* datasploit
* gitminer
* Cr3dOv3r
* UfoNET

#### Notes:

* Adding them as I go. Don't expect production-ready images  
* Uses either either using python-slim or python-alpine
* Tools will show help dialog if no arguments are passed  

## Usage:

```bash
# Build a specific tool for deployment
git clone https://github.com/khast3x/Offensive-Dockerfiles.git
cd Offensive-Dockerfiles
cd <tool>
docker build -t <tool> .
docker run -it <tool>:latest

# Use docker-compose to build all in one
git clone https://github.com/khast3x/Offensive-Dockerfiles.git
cd Offensive-Dockerfiles
docker-compose build
```