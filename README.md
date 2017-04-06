# Fabio SVG Logo Creator

This creates the Fabio logo with the Microgramma font.
Since this isn't a free font and I only have the TTF
you need to use something like Sketch to convert the
text to an outline path and then export to SVG.

Embedding the necessary glyphs of the font works only
in Safari and not any of the other browsers.

1. `go run main.go > logo-with-font.svg`
1. Open `logo-with-font.svg` in Sketch
1. Right-click the "Fabio" text and select `Convert to Outlines`
1. Select the entire logo and click `Make Exportable` in the bottom-right corner
   and then export to SVG

![Fabio](https://cdn.rawgit.com/fabiolb/logo/74a70c8a/logo.svg)

(Update link in README with new git hash when image changes.)
