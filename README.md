# 🚀 gloc

```text

       _
  __ _| | ___   ___
 / _` | |/ _ \ / __|
| (_| | | (_) | (__
 \__, |_|\___/ \___|
 |___/
```

🚀 gloc – A blazing-fast LOC (Lines of Code) counter in Go, inspired by tokei. Simple & efficient!

## ✨ Features

### Features

- 🚀 **Fast and Efficient** – Quickly counts lines of code with minimal overhead.
- 🏗️ **Multi-Language Support** – Supports a wide range of programming languages.
- 🛠️ **Simple CLI Interface** – Easy-to-use command-line tool with straightforward commands.
- 📊 **Detailed Statistics** – Displays blank, comment, and code lines per language.
- 🐳 **Docker Support** – Run via Docker without local installation.
- ⚡ **Optimized Performance** – Faster than traditional `cloc` in most scenarios.
- 🔄 **CI/CD Integration** – Works seamlessly with Jenkins, GitHub Actions, and other CI/CD tools.
- 🎯 **Custom Exclusions** – Exclude specific files and extensions for precise analysis.

## 🚀 Installation

Download from [latest releases ](https://github.com/trinhminhtriet/gloc/releases)

## 💡 Usage

### Basic Usage

```sh
make build
rm -rf /usr/local/bin/gloc
ln -s $(pwd)/bin/gloc /usr/local/bin/gloc
gloc --version
gloc .
```

```
$ gloc .
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Markdown                         3              8              0             18
Go                               1             29              1            323
-------------------------------------------------------------------------------
TOTAL                            4             37              1            341
-------------------------------------------------------------------------------
```

### Via Docker

with [dockerhub](https://hub.docker.com/repository/docker/trinhminhtriet/gloc)

```
docker pull trinhminhtriet/gloc
docker run --rm -v "${PWD}":/workdir trinhminhtriet/gloc .
```

### Integration Jenkins CI

use [SLOCCount Plugin](https://wiki.jenkins-ci.org/display/JENKINS/SLOCCount+Plugin).

```
$ cloc --by-file --output-type=sloccount . > sloccount.scc
```

```
$ cat sloccount.scc
398 Go      ./main.go
190 Go      ./language.go
132 Markdown        ./README.md
24  Go      ./xml.go
18  Go      ./file.go
15  Go      ./option.go
```

## Support Languages

use `--show-lang` option

```
$ gloc --show-lang
```

## Performance

- CPU 3.8GHz 8core Intel Core i7 / 32GB 2667MHz DDR4 / MacOSX 13.3.1
- cloc 1.96
- tokei 12.1.2 compiled with serialization support: json
- gloc v0.1.1 [1dfa360](https://github.com/trinhminhtriet/gloc/commit/1dfa360c412270a7d9e12346eec7052c0d904ff4)
- target repository is [golang/go commit:f742ddc](https://github.com/golang/go/tree/f742ddc349723667fc9af5d0f16233f7762aeaa0)

### cloc

```
$ time cloc .
   12003 text files.
   11150 unique files.
    1192 files ignored.

8 errors:
Line count, exceeded timeout:  ./src/cmd/dist/build.go
Line count, exceeded timeout:  ./src/cmd/trace/static/webcomponents.min.js
Line count, exceeded timeout:  ./src/net/http/requestwrite_test.go
Line count, exceeded timeout:  ./src/vendor/golang.org/x/net/idna/tables10.0.0.go
Line count, exceeded timeout:  ./src/vendor/golang.org/x/net/idna/tables11.0.0.go
Line count, exceeded timeout:  ./src/vendor/golang.org/x/net/idna/tables12.0.0.go
Line count, exceeded timeout:  ./src/vendor/golang.org/x/net/idna/tables13.0.0.go
Line count, exceeded timeout:  ./src/vendor/golang.org/x/net/idna/tables9.0.0.go

github.com/AlDanial/cloc v 1.96  T=35.07 s (317.9 files/s, 78679.3 lines/s)
-----------------------------------------------------------------------------------
Language                         files          blank        comment           code
-----------------------------------------------------------------------------------
Go                                9081         205135         337681        1779107
Text                              1194          11530              0         210849
Assembly                           563          15549          21625         122329
HTML                                17           3197             78          24983
C                                  139           1324            982           6895
JSON                                20              0              0           3122
CSV                                  1              0              0           2119
Markdown                            27            674            106           1949
Bourne Shell                        16            253            868           1664
JavaScript                          10            234            221           1517
Perl                                10            173            171           1111
C/C++ Header                        26            145            346            724
Bourne Again Shell                  16            120            263            535
Python                               1            133            104            375
CSS                                  3              4             13            337
DOS Batch                            5             56             66            207
Windows Resource File                4             23              0            146
Logos                                2             16              0            101
Dockerfile                           2             13             15             47
C++                                  2             11             14             24
make                                 5              9             10             21
Objective-C                          1              2              3             11
Fortran 90                           2              1              3              8
awk                                  1              1              6              7
YAML                                 1              0              0              5
MATLAB                               1              1              0              4
-----------------------------------------------------------------------------------
SUM:                             11150         238604         362575        2158197
-----------------------------------------------------------------------------------
cloc .  33.70s user 1.48s system 99% cpu 35.237 total
```

### tokei

```
$ time tokei --sort code  --exclude "**/*.txt" .
===============================================================================
 Language            Files        Lines         Code     Comments       Blanks
===============================================================================
 Go                   9242      2330107      1812147       318036       199924
 GNU Style Assembly    565       159534       127093        16888        15553
 C                     143         9272         6949         1000         1323
 JSON                   21         3122         3122            0            0
 Shell                  16         2785         2267          342          176
 JavaScript             10         1972         1520          218          234
 Perl                    9         1360         1032          170          158
 C Header               27         1222          727          349          146
 BASH                   16          918          521          279          118
 Python                  1          612          421           70          121
 CSS                     3          354          337           13            4
 Autoconf                9          283          274            0            9
 Batch                   5          329          207           66           56
 Alex                    2          117          101            0           16
 Dockerfile              2           75           47           15           13
 C++                     2           49           24           14           11
 Makefile                5           40           20           10           10
 Objective-C             2           21           15            3            3
 FORTRAN Modern          2           12            8            3            1
 Markdown               18         2402            0         1853          549
-------------------------------------------------------------------------------
 HTML                   17        19060        18584           49          427
 |- CSS                  4         2071         1852           10          209
 |- HTML                 1          219          212            0            7
 |- JavaScript           8         6920         6876           16           28
 (Total)                          28270        27524           75          671
===============================================================================
 Total               10117      2533646      1975416       339378       218852
===============================================================================
tokei --sort code --exclude "**/*.txt" .  0.76s user 0.50s system 562% cpu 0.224 total
```

### gloc

```
time gloc --exclude-ext=txt .
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                            9096         205242         352844        1764503
Assembly                       563          15555          21624         122324
HTML                            17           3197            212          24849
C                              139           1324            983           6894
JSON                            20              0              0           3122
BASH                            27            345           1106           2122
Markdown                        18            549             28           1825
JavaScript                      10            234            218           1520
C Header                        26            145            346            724
Perl                            10            173            584            698
Python                           1            133            104            375
CSS                              3              4             13            337
Batch                            5             56              0            273
Plan9 Shell                      4             23             50             96
Bourne Shell                     5             28             24             78
C++                              2             11             14             24
Makefile                         5             10             10             20
Objective-C                      2              3              3             15
FORTRAN Modern                   2              1              3              8
Awk                              1              1              6              7
-------------------------------------------------------------------------------
TOTAL                         9956         227034         378172        1929814
-------------------------------------------------------------------------------
gloc --exclude-ext=txt .  0.65s user 0.51s system 119% cpu 0.970 total
```

## 🤝 How to contribute

We welcome contributions!

- Fork this repository;
- Create a branch with your feature: `git checkout -b my-feature`;
- Commit your changes: `git commit -m "feat: my new feature"`;
- Push to your branch: `git push origin my-feature`.

Once your pull request has been merged, you can delete your branch.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
