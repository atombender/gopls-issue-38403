# Steps

In VSCode:

* Add a dummy edit to `two.go` (e.g. add an empty line) and save.
* In `one.go`, inside the function `One()`, after line 4, start typing: `Three`. In other words, it should look like:

```go
  url := Two()
  Three
```

* `two.go` should turn red in the Explorer, and if you navigate to it, you should see:

![Screenshot](https://i.imgur.com/CvlhAU0.jpg)

# Version

Confirmed with:

* gopls `4b814e061378b7244e65d8b5c81f0100e8ded409` (`golang.org/x/tools/gopls@v0.0.0-20200501005904-d351ea090f9b`)
* gopls `ff647943c706f2183c0ae2c7085bad7aa4d1dedd` (`golang.org/x/tools/gopls@v0.0.0-20200507172948-ff647943c706`)

