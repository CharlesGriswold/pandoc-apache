# pandoc-apache
An apache handler for markdown (and other) files that uses pandoc for a backend. Written in Go.

Pandoc-apache can currently handle files with the following extensions: .docbook, .docx, .epub, .haddock, .md, .odt, .opml, .org, .rst, .t2t, and .textile.

Output is HTML5 by default. Other output types can be specified by append a ?TYPE to the URL. Currently supported output types include raw, plain, HTML4, and HTML5. For instance, `http://www.yoursite.com/index.md?raw` would dump the raw text of the file
