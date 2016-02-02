# revealgo

**revealgo** is a small web application for giving Markdown-driven presentations implemented in **Go**! The `revealgo` command starts a local web server to serve the your markdown presentation file with `reveal.js`. The presentation can be viewed in a web browser. The reveal.js library offers comprehensive presenting features such as slide transitions, speaker notes and more.

## Install

To install, use `go get`:

```
$ go get github.com/yusukebe/revealgo/cmd/revealgo
```

## Usage

The usage:

```
$ revealgo [options] MARKDOWN_FILE
```

Available options:

```
-p, --port            tcp port number of this server. default is 3000.
--theme               slide theme or original css file name. default themes:
                      beige, black, blood, league, moon, night, serif, simple, sky, solarized, and white
--transition          transition effect for slides: default, cube, page, concave, zoom, linear, fade, none
```

### Sample Makrdown

```markdown
## This is an H2 Title

Description...

The horizontal slide separator characters are '---'

---

# This is second title

The vertical slide separator characters are '___'

___

## This is a third title

---

## This is a forth title
<!-- .slide: data-background="#f70000" data-transition="page" -->

You can add slide attributes like above.
```

## See Also

* reveal.js <https://github.com/hakimel/reveal.js/>
* App::revealup <https://github.com/yusukebe/App-revealup>

## Author

Yusuke Wada <http://github.com/yusukebe>

