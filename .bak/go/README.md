# READEM

1. How to generate openssl RSA keys

![GO加密解密之RSA](http://blog.studygolang.com/2013/01/go%E5%8A%A0%E5%AF%86%E8%A7%A3%E5%AF%86%E4%B9%8Brsa/)

```
	openssl genrsa -out private.pem 1024
	openssl rsa -in private.pem -pubout -out public.pem
```

2. How to generate HTML from MD in go

![md](https://github.com/gomarkdown/markdown)

```
	go get -u github.com/gomarkdown/mdtohtml
```