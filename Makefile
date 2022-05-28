run:
	mdbook serve -o
d:
	git status
	git add .
	git commit -m "fuck day"
	git push
	git status
	git log

