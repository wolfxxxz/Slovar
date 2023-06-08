build:
	go build

run:
	./SlovarNV
autosafe:
	git commit -a -m "autosafe"
pushsafe:
	git push origin master "autosafe"