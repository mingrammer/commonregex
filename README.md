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
  parser := cregex.New()
  
  text := `John, please get that article on www.linkedin.com to me by 5:00PM on Jan 9th 2012. 4:00 would be ideal, actually. If you have any questions, You can reach me at (519)-236-2723x341 or get in touch with my associate at harold.smith@gmail.com`
  
  date_list := parser.Date(text)
  // ['Jan 9th 2012']
  time_list := parser.Time(text)
  // ['5:00PM', '4:00']
  link_list := parser.Links(text)
  // ['www.linkedin.com']
  phone_list := parser.PhonesWithExts(text)  
  // ['(519)-236-2723x341']
  email_list := parser.Emails(text)
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
* btc address
* street address
* zip code
* po box