
<h1 align="center">
  <br>
  <a href="https://github.com/s0md3v/Photon"><img src="https://image.ibb.co/h5OZAK/photonsmall.png" alt="Photon"></a>
  <br>
  Photon
  <br>
</h1>

<h4 align="center">Incredibly fast crawler designed for OSINT.</h4>

<p align="center">
  <a href="https://github.com/s0md3v/Photon/releases">
    <img src="https://img.shields.io/github/release/s0md3v/Photon.svg">
  </a>
  <a href="https://pypi.org/project/photon/">
    <img src="https://img.shields.io/badge/pypi-@photon-red.svg?style=style=flat-square"
         alt="pypi">
  </a>
  <a href="https://github.com/s0md3v/Photon/issues?q=is%3Aissue+is%3Aclosed">
      <img src="https://img.shields.io/github/issues-closed-raw/s0md3v/Photon.svg">
  </a>
  <a href="https://travis-ci.com/s0md3v/Photon">
    <img src="https://img.shields.io/travis/com/s0md3v/Photon.svg">
  </a>
</p>

This is an Alpine build Docker image.  
For the rest of the documentation, please visit [s0md3v](ttps://github.com/s0md3v/Photon)'s repo.

#### Docker

Photon can be launched using a lightweight Python-Alpine (103MB) Docker image.  
```bash
$ git clone https://github.com/s0md3v/Photon.git
$ cd Photon
$ docker build -t photon .
$ docker run -it --name photon photon:latest -u google.com
```
To view results, you can either head over to the local docker volume, which you can find by running `docker inspect photon` or by mounting the target loot folder:

```bash
$ docker run -it --name photon -v "$PWD:/Photon/google.com" -u google.com
```
