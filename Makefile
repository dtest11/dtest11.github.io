run:
	mdbook serve -o
d:
	git status
	git add .
	git commit -m "pub"
	git push

