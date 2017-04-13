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

<img src="https://cdn.rawgit.com/fabiolb/logo/f966a7d/logo.svg" height="48"/>

(Update link in README with new git hash when image changes.)

To create the logo without the fabio text run

`go run main.go -text=false > logo-no-text.svg`

This creates a 272x145 canvas with a transparent background.
I need to update the code to create a rectangular logo. The
`logo-no-text*` PNGs have been created by hand.
