<h1 align="center">
  <br>
  <img src="https://i.imgur.com/7MQg7EH.jpg" alt="Digart pi" width="850">
  <br>
</h1>

<h4 align="center">circle visualization of a number's digits in Go</h4>

<p align="center">
  <a href="https://goreportcard.com/report/github.com/hugolgst/digart"><img src="https://goreportcard.com/badge/github.com/hugolgst/digart"></a>
  <a href="https://travis-ci.org/hugolgst/digart"><img src="https://travis-ci.org/hugolgst/digart.svg?branch=master"></a>
</p>

<p align="center">
  <a href="#introduction">Introduction</a> —
  <a href="#getting-started">Getting started</a> —
  <a href="#license">License</a>
</p>

## Introduction
In the image below, each dot represents a decimal which is in a segment that represents the previous decimal, its position is its position in π.

## Getting started
Download the [latest release](https://github.com/hugolgst/digart/releases):

And then run the tool

```bash
$ ./digart
```

To change number you can just write the file path
```
resources/e.txt
resources/phi.txt
resources/pi.txt
```

And the number of digits (not lower than 1500)

```bash
$ ./digart resources/e.txt 3000
```

## License
[MIT](https://github.com/hugolgst/digart/blob/master/LICENSE)
