# revealgo

**revealgo** is a small web application for giving Markdown-driven presentations implemented in **Go**! The `revealgo` command starts a local web server to serve the your markdown presentation file with `reveal.js`. The presentation can be viewed in a web browser. The reveal.js library offers comprehensive presenting features such as slide transitions, speaker notes and more.

## Installation

### Binary

Download the binary from [**GitHub Releases**](https://github.com/yusukebe/revealgo/releases) and install it somewhere in your `$PATH`. `revealgo` currently provides pre-built binaries for Linux, macOS and Windows.

### Homebrew

You can also install via [hombrew](https://brew.sh) on macOS:

```
$ brew install yusukebe/tap/revealgo
```

### Source

To install from the source, use `go install` after `git clone`:

```
$ git clone git@github.com:yusukebe/revealgo.git
$ cd revealgo
$ git submodule update --init --recursive
$ go install ./cmd/revealgo
```

## Usage

The usage:

```
$ revealgo [options] MARKDOWN.md
```

Then access the local web server such as `http://localhost:3000` with Chrome, Firefox, or Safari.

Available options:

```text
-p, --port            TCP port number of this server (default: 3000)
--theme               Slide theme or original css file name. default themes:
                      beige, black, blood, league, moon, night, serif, simple, sky, solarized, and white (default: black.css)
--template            Custom HTML template file name. default template: slide
--disable-auto-open   Disable automatic opening of the browser
--transition          Transition effect for slides: default, cube, page, concave, zoom, linear, fade, none (default: default)
--separator           Horizontal slide separator characters (default: ^---)
--vertical-separator  Vertical slide separator characters (default: ^___)
--multiplex           Enable slide multiplexing
-v, --version         Show the version
```

### Screenshots

Run `revealgo` command:

![Command Line](https://cloud.githubusercontent.com/assets/10682/12741641/b5afb504-c9c1-11e5-94d6-c364912cfcc2.png)

Open the server address with your web browser:

![Slides](https://cloud.githubusercontent.com/assets/10682/12741672/f9cda548-c9c1-11e5-9c21-fcaf1af3cdf4.png)

### Sample Markdown

```text
## This is an H2 Title

Description...

The horizontal slide separator characters are '---'

---

# This is second title

The vertical slide separator characters are '^___'

___

## This is a third title

---

## This is a forth title

<!-- .slide: data-background="#f70000" data-transition="page" -->

You can add slide attributes like above.
```

### Customize Theme

While `revealgo` is running, open another terminal and get the theme file `black.css`:

```
$ curl http://localhost:3000/revealjs/css/theme/black.css > original.css
```

Edit `original.css`, And then run `revealgo` with `--theme` option:

```
$ revealgo --theme original.css slide.md
```

### Customize Slide Configurations

Get the default slide HTML file:

```
$ curl http://localhost:3000/ > slide.html
```

Edit `slide.html`, and then open `http://localhost:3000/slide.html` with your browser. A slide with the modified configurations will come up.

### Customize Slide template

A custom `slide.html` can also be provided using the `--template` option, or putting a `slide.html` next to your custom theme file.
This allows you to use templating within your `slide.html`

You could use the `slide.html` output from your localhost, or grab a copy of the orignal asset to use as the base.
This can be obtained from: <https://github.com/yusukebe/revealgo/blob/master/assets/templates/slide.html>

### Using slide multiplex

> The multiplex plugin allows your audience to follow the slides of the
> presentation you are controlling on their own phone, tablet or laptop
>
> --- [reveal.js site](https://revealjs.com/multiplex/)

When `--multiplex` is enabled, the client slides can be found on the `/` path and
the master ones under `/master/`. The master presentation will push its changes
to all the client ones for every transition, is like having a remote control!

For example, your laptop's IP address in the local network is `192.168.100.10`
and you are using the port `3000`, so your audience should see the slides on
`http://192.168.100.10:3000/`, and you should be able to control their slides
through `http://192.168.100.10:3000/master/`.

**NOTE**: Bear in mind multiplex feature will not work as expected when 1) the
presenter computer firewall denies incomig traffic or 2) the local network
does not allow traffic between devices on the port you picked

## Related projects

- reveal.js <https://github.com/hakimel/reveal.js/>
- App::revealup <https://github.com/yusukebe/App-revealup>
- reveal-md <https://github.com/webpro/reveal-md>

## Contributing

See [docs/CONTRIBUTING.md](./docs/CONTRIBUTING.md).

## Contributors

Thanks to all [contributors](https://github.com/yusukebe/revealgo/graphs/contributors)!

## Author

Yusuke Wada <http://github.com/yusukebe>

## License

Distributed under the MIT License. See [LICENSE](LICENSE) for more information.
