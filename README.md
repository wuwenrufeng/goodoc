# goodoc
=======

用go实现网页源码转换为pdf,markdown及office文档。html2pdf，html2word, html2markdown, markdown2html, html2excel...

goodoc ≈ good go doc


![](https://gitee.com/wuwenrufeng/assets/raw/master/goodoc/goodoc.gif)

## How to setup goodoc

#### 0. Install pandoc and pdfTex

#### 1. Install [pandoc](https://pandoc.org/installing.html), and [LaTeX](https://tug.org/mactex/) should be installed  too.

```
➜  ~  pandoc --version
pandoc 2.17.1.1
➜  ~  latex --version
pdfTeX 3.141592653-2.6-1.40.22 (TeX Live 2021)
kpathsea version 6.3.3
```

#### 2. Clone the repo 

```
git clone git@github.com:wuwenrufeng/goodoc.git
```

#### 3. Compile the source and run examples

```
cd examples

go run *.go
```
#### 4. run http server
```
cd server

go run main.go
```

## TODO

- logs

- more tests

- more examples
