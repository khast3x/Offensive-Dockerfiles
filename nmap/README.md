# Nmap, nmap scripts, vulscan
> Famous Nmap scanner built on a very light Alpine image.  
Comes fully equipped with
  the latest Nmap Scripting Engine (NSE) modules, as well as the [Vulscan](https://github.com/scipag/vulscan) NSE script.  
  The databases used by Vulscan are pulled using the original updater script when image is built  
  
### Usage:
```bash
docker run -it nmap -sV --script=vulscan/vulscan.nse www.example.com
```