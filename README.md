# Readon [![Build Status](https://travis-ci.org/pascalj/readon.svg?branch=master)](https://travis-ci.org/pascalj/readon)

Readon is a library written in Go. It tries to get only the content from a website and it is roughly inspired by arc90's [Readability](http://lab.arc90.com/2009/03/02/readability/).

## Usage

Just call `readon.NewArticle` and pass it a io.Reader with the whole website. It will return an `readon.Article` containing the title and the content of the website.

```go
file, _ := os.Open("my_article.html")
article, err = readon.NewArticle(file)
if err == nil {
	// output the articles essential html without <html> and <body>
	fmt.Println(article.ArticleHtml)
}
```

Please note that Readon is using heuristics to determine what the content of a website is. It may get something wrong. In that case, please provide an example and file a bug. Thank you!

## Contribution

Contributions of any kind are highly welcome. If you want to contribute code directly, just send a pull request:

1. Fork the project
2. Make your changes
3. Submit a pull request

## BSD License

Copyright (c) 2014, Pascal Jungblut <oss@pascalj.de>
All rights reserved.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
