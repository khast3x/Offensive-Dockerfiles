# Offensive dockerfiles
![Docker](https://anonimag.es/i/bf7d4fd20803d0bab7be32fa51c4a498.png)

* Adding them as I go. Don't expect production-ready images.  
* Uses either either using python-slim or python-alpine
* Tools will show help dialog if no arguments are passed  

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

## Usage:

```bash
git clone https://github.com/khast3x/Offensive-Dockerfiles.git
cd <tool>
docker build -t <tool> .
docker run -it <tool>:latest
```