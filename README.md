[![Go Report Card](https://goreportcard.com/badge/github.com/mingrammer/commonregex)](https://goreportcard.com/report/github.com/mingrammer/commonregex) ![build](https://travis-ci.org/mingrammer/commonregex.svg?branch=master)

# commonregex

> Inspired by [CommonRegex](https://github.com/madisonmay/CommonRegex)

This is a collection of often used regular expressions. It provides these as simple functions for getting the matched strings corresponding to specific patterns.

## Installation
```shell
go get github.com/mingrammer/commonregex
```

## Usage

```go
import (
	cregex "github.com/mingrammer/commonregex"
)

func main() {
  text := `John, please get that article on www.linkedin.com to me by 5:00PM on Jan 9th 2012. 4:00 would be ideal, actually. If you have any questions, You can reach me at (519)-236-2723x341 or get in touch with my associate at harold.smith@gmail.com`
  
  date_list := cregex.Date(text)
  // ['Jan 9th 2012']
  time_list := cregex.Time(text)
  // ['5:00PM', '4:00']
  link_list := cregex.Links(text)
  // ['www.linkedin.com', 'harold.smith@gmail.com']
  phone_list := cregex.PhonesWithExts(text)  
  // ['(519)-236-2723x341']
  email_list := cregex.Emails(text)
  // ['harold.smith@gmail.com']
}
```

## Features

* date
* time
* phone
* phones with exts
* link
* email
* ip
* ipv6
* price
* hex color
* credit card
* VISA credit card
* MC credit card
* ISBN 10/13
* btc address
* street address
* zip code
* po box
* SSN
* MD5
* SHA1
* SHA256
* GUID
* MAC address