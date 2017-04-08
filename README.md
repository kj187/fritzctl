# fritzctl - console AVM FRITZ!Box client

![fritzctl](/images/fritzctl.png?raw=true "fritzctl")

## About

fritzctl is a command line client for the AVM FRITZ!Box primarily focused on the
[AVM Home Automation HTTP Interface](https://avm.de/fileadmin/user_upload/Global/Service/Schnittstellen/AHA-HTTP-Interface.pdf).

Software is tested with

*   FRITZ!Box Fon WLAN 7390 running FRITZ!OS 06.51, 06.80, 06.83
*   FRITZ!Box 6490 Cable running FRITZ!OS 06.63
*   FRITZ!Box 7490 running FRITZ!OS 06.83

## CI farm

*   [![Build Satus TravisCI](https://travis-ci.org/bpicode/fritzctl.svg)](https://travis-ci.org/bpicode/fritzctl) Travis CI
*   [![Build Status CircleCI](https://circleci.com/gh/bpicode/fritzctl/tree/master.svg?style=shield)](https://circleci.com/gh/bpicode/fritzctl) Circle CI
*   [![Build Status GitlabCI](https://gitlab.com/bpicode/fritzctl/badges/master/build.svg)](https://gitlab.com/bpicode/fritzctl/commits/master) Gitlab CI
*   [![Build Status SemaphoreCI](https://semaphoreci.com/api/v1/bpicode/fritzctl/branches/master/shields_badge.svg)](https://semaphoreci.com/bpicode/fritzctl)  Semaphore CI
*   [![Build status AppVeyor](https://ci.appveyor.com/api/projects/status/k7qqx91w6mja3u7h?svg=true)](https://ci.appveyor.com/project/bpicode/fritzctl) AppVeyor
    
## Meta

*    [![License](https://img.shields.io/github/license/bpicode/fritzctl.svg)](https://opensource.org/licenses/MIT)
*    [![GoDoc](https://godoc.org/github.com/bpicode/fritzctl?status.svg)](https://godoc.org/github.com/bpicode/fritzctl)
*    [![Repository size](https://reposs.herokuapp.com/?path=bpicode/fritzctl)](https://github.com/bpicode/fritzctl)

## Code metrics

*    [![Go Report Card](https://goreportcard.com/badge/github.com/bpicode/fritzctl)](https://goreportcard.com/report/github.com/bpicode/fritzctl) Static analysis
*    [![codecov](https://codecov.io/gh/bpicode/fritzctl/branch/master/graph/badge.svg)](https://codecov.io/gh/bpicode/fritzctl) Test coverage
*    [![Issue Count](https://codeclimate.com/github/bpicode/fritzctl/badges/issue_count.svg)](https://codeclimate.com/github/bpicode/fritzctl) Code issues
*    [![Code Climate](https://codeclimate.com/github/bpicode/fritzctl/badges/gpa.svg)](https://codeclimate.com/github/bpicode/fritzctl) Code GPA
*    [![codebeat badge](https://codebeat.co/badges/605cf539-21dd-4a60-a892-e0d6da3021fe)](https://codebeat.co/projects/github-com-bpicode-fritzctl) Code rating

## Latest binaries

*   [![Download .deb](https://api.bintray.com/packages/bpicode/fritzctl_deb/fritzctl/images/download.svg)](https://bintray.com/bpicode/fritzctl_deb/fritzctl/_latestVersion)
    .deb packages
*   [![Download .rpm](https://api.bintray.com/packages/bpicode/fritzctl_rpm/fritzctl/images/download.svg)](https://bintray.com/bpicode/fritzctl_rpm/fritzctl/_latestVersion)
    .rpm packages 
*   [![Download](https://api.bintray.com/packages/bpicode/fritzctl_win/fritzctl/images/download.svg)](https://bintray.com/bpicode/fritzctl_win/fritzctl/_latestVersion)
    .zip windows

[![Get automatic notifications about new fritzctl versions (deb)](https://www.bintray.com/docs/images/bintray_badge_color.png)](https://bintray.com/bpicode/fritzctl_deb/fritzctl?source=watch)
[![Get automatic notifications about new fritzctl versions (rpm)](https://www.bintray.com/docs/images/bintray_badge_bw.png)](https://bintray.com/bpicode/fritzctl_rpm/fritzctl?source=watch)
[![Get automatic notifications about new fritzctl versions (rpm)](https://www.bintray.com/docs/images/bintray_badge_greyscale.png)](https://bintray.com/bpicode/fritzctl_win/fritzctl?source=watch)


## Install (debian/ubuntu/...)

Add the repository

```bash
echo "deb https://dl.bintray.com/bpicode/fritzctl_deb jessie main" | sudo tee -a /etc/apt/sources.list
```

and its signing key

```bash
wget -qO - https://api.bintray.com/users/bpicode/keys/gpg/public.key | sudo apt-key add -
```

The fingerprint of the repository key `3072D/35E71039` is
`93AC 2A3D 418B 9C93 2986  6463 15FC CFC9 35E7 1039`.
Update your local repository data and install

```bash
sudo apt update
sudo apt install fritzctl
```

## Install (opensuse)

Add the repository

```bash
wget https://bintray.com/bpicode/fritzctl_rpm/rpm -O bintray-bpicode-fritzctl_rpm.repo && sudo zypper ar -f bintray-bpicode-fritzctl_rpm.repo && rm bintray-bpicode-fritzctl_rpm.repo
```

Update your local repository data and install

```bash
sudo zypper refresh
sudo zypper in fritzctl
```

## Install (windows)

Windows binaries can found in the [windows directory](https://dl.bintray.com/bpicode/fritzctl_win/).

## Direct downloads

There are several locations from where one can download the packages, e.g.

*   directly from the [debian repository](https://bintray.com/bpicode/fritzctl_deb/fritzctl)
    or the [directory index](https://dl.bintray.com/bpicode/fritzctl_deb/)
*   directly from the [rpm repository](https://bintray.com/bpicode/fritzctl_rpm/fritzctl)
    or the [directory index](https://dl.bintray.com/bpicode/fritzctl_rpm/),
*   directly from the [windows repository](https://bintray.com/bpicode/fritzctl_win/fritzctl)
    or the [directory index](https://dl.bintray.com/bpicode/fritzctl_win/).

## Usage

![Demo usage](/images/fritzctl_demo.gif?raw=true "Demo usage")

## Credits

The fritzctl image is licensed under the Creative Commons 3.0 Attributions license. It is build upon the following work:
*   The Go gopher was designed by [Renee French](http://reneefrench.blogspot.com/), licensed under the Creative Commons 3.0 Attributions license.
*   The Go gopher image was created by [Takuya Ueda](https://twitter.com/tenntenn), licensed under the Creative Commons 3.0 Attributions license. At the time of this writing it was available at [golang-samples/gopher-vector](https://github.com/golang-samples/gopher-vector/blob/master/gopher.svg).
*   The router image was created by [Sascha Doerdelmann](https://pixabay.com/en/users/saschadoerdelmann-4359717/), licensed under the Creative Commons CC0 Public Domain Dedication. At the time of this writing it was available at [pixabay](https://pixabay.com/en/wlan-telecommunications-router-2007682/).

