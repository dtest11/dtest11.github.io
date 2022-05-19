gitbook build
rm -rf docs
mkdir docs

mv _book/* docs/
git add .
git commit -m "push docs"
git push -f