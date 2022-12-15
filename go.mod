module github.com/Lord-Leonard/go_pdf_example

go 1.19

require github.com/johnfercher/maroto v0.39.0

replace (
	github.com/johnfercher/maroto => ../maroto
)

require (
	github.com/boombuler/barcode v1.0.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/jung-kurt/gofpdf v1.16.2 // indirect
	github.com/ruudk/golang-pdf417 v0.0.0-20201230142125-a7e3863a1245 // indirect
)
