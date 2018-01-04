# Nmap, nmap scripts, vulscan and vulners

## Abstract
Alpine based.  
Comes fully equipped with
the latest Nmap Scripting Engine (NSE) modules, as well as the [Vulscan](https://github.com/scipag/vulscan) NSE script and the [vulners](https://github.com/vulnersCom/nmap-vulners) API to NSE script.  
The databases used by Vulscan are pulled using the original updater script when image is built  

> You can run both **vulscan** and **vulners** scripts with nmap at the same time, the results will complement each other as shown below  

## Source

https://github.com/scipag/vulscan  
https://github.com/vulnersCom/nmap-vulners

## Usage

```bash
cd nmap/
docker build -t nmap .
docker run -it nmap:latest
docker run -it nmap -sV --script=vulscan,vulscan.nse www.example.com

```

## Help:

```bash
docker run -it nmap # will show nmap help
```

```bash

docker run -it nmap -sV --script vulners,vulscan example.com

Starting Nmap 7.60 ( https://nmap.org ) at 2018-01-04 11:05 UTC
Nmap scan report for example.com (93.184.216.34)
Host is up (0.083s latency).
Other addresses for example.com (not scanned): 2606:2800:220:1:248:1893:25c8:1946
Not shown: 993 filtered ports
PORT     STATE  SERVICE  VERSION
53/tcp   closed domain
80/tcp   open   http     Edgecast CDN httpd (dca/5328)
| http-server-header: 
|   ECS (dca/246E)
|_  ECS (dca/24C1)
| vulscan: scip VulDB - http://www.scip.ch/en/?vuldb:
| [106716] Fastly CDN Module up to 1.2.25 on Magento2 information disclosure
| [74135] Cdnvote up to 0.4.0 cdnvote-post.php sql injection
| [61925] Wimleers CDN 6.x-2.2/7.x-2.2 settings.php unknown vulnerability
| [61647] Kyle Browning CDN2 Video 6.x-1.x cross site request forgery
| [61646] Kyle Browning CDN2 Video 6.x-1.x cross site scripting
| [42063] CDNetworks Download Client 1.0.5 ActiveX Control neffylauncher.dll unknown vulnerability
| [42062] CDNetworks Download Client 1.0.5 ActiveX Control neffylauncher.dll directory traversal
| 
| MITRE CVE - http://cve.mitre.org:
| [CVE-2012-5159] phpMyAdmin 3.5.2.2, as distributed by the cdnetworks-kr-1 mirror during an unspecified time frame in 2012, contains an externally introduced modification (Trojan Horse) in server_sync.php, which allows remote attackers to execute arbitrary PHP code via an eval injection attack.
| [CVE-2012-2917] Cross-site scripting (XSS) vulnerability in the Share and Follow plugin 1.80.3 for WordPress allows remote attackers to inject arbitrary web script or HTML via the CDN API Key (cnd-key) in a share-and-follow-menu page to wp-admin/admin.php.
| [CVE-2012-2155] Cross-site request forgery (CSRF) vulnerability in the CDN2 Video module 6.x for Drupal allows remote attackers to hijack the authentication of unspecified victims via unknown vectors.
| [CVE-2012-2154] Cross-site scripting (XSS) vulnerability in the CDN2 Video module 6.x for Drupal allows remote attackers to inject arbitrary web script or HTML via unspecified vectors.
| [CVE-2012-1645] The CDN module 6.x-2.2 and 7.x-2.2 for Drupal, when running in Origin Pull mode with the "Far Future expiration" option enabled, allows remote attackers to read arbitrary PHP files via unspecified vectors, as demonstrated by reading settings.php.
| [CVE-2008-1886] The NeffyLauncher 1.0.5 ActiveX control (NeffyLauncher.dll) in CDNetworks Nefficient Download uses weak cryptography for a KeyCode that blocks unauthorized use of the control, which allows remote attackers to bypass this protection mechanism by calculating the required KeyCode.  NOTE: this can be used by arbitrary web sites to host exploit code that targets this control.
| [CVE-2008-1885] Directory traversal vulnerability in the NeffyLauncher 1.0.5 ActiveX control (NeffyLauncher.dll) in CDNetworks Nefficient Download allows remote attackers to download arbitrary code onto a client system via a .. (dot dot) in the SkinPath parameter and a .zip URL in the HttpSkin parameter.  NOTE: this can be leveraged for code execution by writing to a Startup folder.
| 
| OSVDB - http://www.osvdb.org:
| [80686] CDN2 Video Module for Drupal Form API Unspecified CSRF
| [80685] CDN2 Video Module for Drupal Unspecified XSS
| [79317] CDN Module for Drupal PHP File Source Code Disclosure
| [71039] cdnvote Plugin for WordPress cdnvote-post.php Multiple Parameter SQL Injection
| [44460] CDNetworks Nefficient Download NeffyLauncher ActiveX (NeffyLauncher.dll) KeyCode Cryptography Weakness
| [44247] CDNetworks Nefficient Download NeffyLauncher ActiveX (NeffyLauncher.dll) SkinPath Property Traversal Arbitrary File Download
| 
| SecurityFocus - http://www.securityfocus.com/bid/:
| [83663] CDNsun and OnApp Remote Denial of Service Vulnerability
| [52812] Drupal CDN2 Video Module Cross Site Request Forgery and Cross Site Scripting Vulnerabilities
| [52041] Drupal CDN Module Information Disclosure Vulnerability
| [46483] WordPress cdnvote 'cdnvote-post.php' Multiple SQL Injection Vulnerabilities
| [28666] CDNetworks Nefficient Download 'NeffyLauncher.dll' ActiveX Control Multiple Vulnerabilities
| 
| SecurityTracker - http://www.securitytracker.com:
| [1002416] Cisco Internet Content Distribution Network (iCDN) Products May Let Remote Users Masquerade as Valid Authenticated Clients Due to Security Flaw in RSA BSAFE Library
| 
| IBM X-Force - http://xforce.iss.net:
| [74522] CDN2 Video for Drupal unspecified cross-site request forgery
| [74520] CDN2 Video module for Drupal unspecified cross-site scripting
| [73250] CDN module for Drupal unspecified information disclosure
| [65630] cdnvote plugin for WordPress cdnvote-post.php SQL injection
| 
| Exploit-DB - http://www.exploit-db.com:
| [5397] CDNetworks Nefficient Download (NeffyLauncher.dll) Code Execution Vuln
| 
| OpenVAS (Nessus) - http://www.openvas.org:
| No findings
|_
443/tcp  open   ssl/http Edgecast CDN httpd (dca/24F9)
| vulscan: scip VulDB - http://www.scip.ch/en/?vuldb:
| [106716] Fastly CDN Module up to 1.2.25 on Magento2 information disclosure
| [74135] Cdnvote up to 0.4.0 cdnvote-post.php sql injection
| [61925] Wimleers CDN 6.x-2.2/7.x-2.2 settings.php unknown vulnerability
| [61647] Kyle Browning CDN2 Video 6.x-1.x cross site request forgery
| [61646] Kyle Browning CDN2 Video 6.x-1.x cross site scripting
| [42063] CDNetworks Download Client 1.0.5 ActiveX Control neffylauncher.dll unknown vulnerability
| [42062] CDNetworks Download Client 1.0.5 ActiveX Control neffylauncher.dll directory traversal
| 
| MITRE CVE - http://cve.mitre.org:
| [CVE-2012-5159] phpMyAdmin 3.5.2.2, as distributed by the cdnetworks-kr-1 mirror during an unspecified time frame in 2012, contains an externally introduced modification (Trojan Horse) in server_sync.php, which allows remote attackers to execute arbitrary PHP code via an eval injection attack.
| [CVE-2012-2917] Cross-site scripting (XSS) vulnerability in the Share and Follow plugin 1.80.3 for WordPress allows remote attackers to inject arbitrary web script or HTML via the CDN API Key (cnd-key) in a share-and-follow-menu page to wp-admin/admin.php.
| [CVE-2012-2155] Cross-site request forgery (CSRF) vulnerability in the CDN2 Video module 6.x for Drupal allows remote attackers to hijack the authentication of unspecified victims via unknown vectors.
| [CVE-2012-2154] Cross-site scripting (XSS) vulnerability in the CDN2 Video module 6.x for Drupal allows remote attackers to inject arbitrary web script or HTML via unspecified vectors.
| [CVE-2012-1645] The CDN module 6.x-2.2 and 7.x-2.2 for Drupal, when running in Origin Pull mode with the "Far Future expiration" option enabled, allows remote attackers to read arbitrary PHP files via unspecified vectors, as demonstrated by reading settings.php.
| [CVE-2008-1886] The NeffyLauncher 1.0.5 ActiveX control (NeffyLauncher.dll) in CDNetworks Nefficient Download uses weak cryptography for a KeyCode that blocks unauthorized use of the control, which allows remote attackers to bypass this protection mechanism by calculating the required KeyCode.  NOTE: this can be used by arbitrary web sites to host exploit code that targets this control.
| [CVE-2008-1885] Directory traversal vulnerability in the NeffyLauncher 1.0.5 ActiveX control (NeffyLauncher.dll) in CDNetworks Nefficient Download allows remote attackers to download arbitrary code onto a client system via a .. (dot dot) in the SkinPath parameter and a .zip URL in the HttpSkin parameter.  NOTE: this can be leveraged for code execution by writing to a Startup folder.
| 
| OSVDB - http://www.osvdb.org:
| [80686] CDN2 Video Module for Drupal Form API Unspecified CSRF
| [80685] CDN2 Video Module for Drupal Unspecified XSS
| [79317] CDN Module for Drupal PHP File Source Code Disclosure
| [71039] cdnvote Plugin for WordPress cdnvote-post.php Multiple Parameter SQL Injection
| [44460] CDNetworks Nefficient Download NeffyLauncher ActiveX (NeffyLauncher.dll) KeyCode Cryptography Weakness
| [44247] CDNetworks Nefficient Download NeffyLauncher ActiveX (NeffyLauncher.dll) SkinPath Property Traversal Arbitrary File Download
| 
| SecurityFocus - http://www.securityfocus.com/bid/:
| [83663] CDNsun and OnApp Remote Denial of Service Vulnerability
| [52812] Drupal CDN2 Video Module Cross Site Request Forgery and Cross Site Scripting Vulnerabilities
| [52041] Drupal CDN Module Information Disclosure Vulnerability
| [46483] WordPress cdnvote 'cdnvote-post.php' Multiple SQL Injection Vulnerabilities
| [28666] CDNetworks Nefficient Download 'NeffyLauncher.dll' ActiveX Control Multiple Vulnerabilities
| 
| SecurityTracker - http://www.securitytracker.com:
| [1002416] Cisco Internet Content Distribution Network (iCDN) Products May Let Remote Users Masquerade as Valid Authenticated Clients Due to Security Flaw in RSA BSAFE Library
| 
| IBM X-Force - http://xforce.iss.net:
| [74522] CDN2 Video for Drupal unspecified cross-site request forgery
| [74520] CDN2 Video module for Drupal unspecified cross-site scripting
| [73250] CDN module for Drupal unspecified information disclosure
| [65630] cdnvote plugin for WordPress cdnvote-post.php SQL injection
| 
| Exploit-DB - http://www.exploit-db.com:
| [5397] CDNetworks Nefficient Download (NeffyLauncher.dll) Code Execution Vuln
| 
| OpenVAS (Nessus) - http://www.openvas.org:
| No findings
|_
554/tcp  closed rtsp
1119/tcp closed bnetgame
1755/tcp closed wms
1935/tcp closed rtmp


```

## Demo  
[![asciicast](https://asciinema.org/a/141837.png)](https://asciinema.org/a/141837)

