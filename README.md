redirect-tracker
================

Redirect tracker prints all redirect urls of given url.

### How to install

```
$ go get -u github.com/li-go/redirect-tracker
```

### Usage

```
$ $GOBIN/redirect-tracker https://google.com
requesting https://google.com (301 Moved Permanently)
requesting https://www.google.com/ (200 OK)
```
