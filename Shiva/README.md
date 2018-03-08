# Shiva

> Source https://github.com/UltimateHackers/Shiva

First of all, put Shiva on watch. I will be upgrading it to a full stress testing suite over time.
Shiva is designed to perform Denial Of Service (DOS) attack on wordpress sites by loading all jquery scripts at once through load-scripts.php. So basically its an exploit for [CVE-2018-6389](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-6389).</br>

### Awesomeness
- Shiva uses multithreading to bring down websites as soon as possible,
- You don't need to worry about your IP being exposed because Shiva uses [Proxify](https://github.com/UltimateHackers/proxify) to route all requests through random proxies which keep getting rotated automatically.

### Dependencies & Compatibility
Shiva is compatible with both python2 and python3.</br>
Dependencies:
- requests
- proxify

### Usages
You can attack a target with Shiva as follows:
```
python shiva.py -u target.com -t 50
```
Where target.com is the target website and 50 is the number of threads.</br>
You must keep in mind that Shiva is only effective against wordpress site so make sure your target runs on wordpress.</br>
Number of threads should be selected according to the network speed.

<img src='https://i.imgur.com/dWDfGnr.png' />

### Docker
To deploy using docker, proceed as follows:
```bash
git clone https://github.com/UltimateHackers/Shiva.git
cd Shiva
docker build -t shiva .
docker run -it shiva -u example.com -t 10

```
