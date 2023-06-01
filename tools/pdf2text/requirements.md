## Image Magick
`sudo apt install libmagickwand-dev`

Remove section:
```xml
<!-- disable ghostscript format types -->
<policy domain="coder" rights="none" pattern="PS" />
<policy domain="coder" rights="none" pattern="PS2" />
<policy domain="coder" rights="none" pattern="PS3" />
<policy domain="coder" rights="none" pattern="EPS" />
<policy domain="coder" rights="none" pattern="PDF" />
<policy domain="coder" rights="none" pattern="XPS" />
```
from 
`/etc/ImageMagick-6/policy.xml`

## Go Tesseract
`sudo apt install libtesseract-dev`

`-tags ocr`