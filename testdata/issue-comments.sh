#!/bin/bash -eux

git checkout -b $1

cat > foo.go <<EOF
package foo
func Foo() {}  // expect golint exported without comment
EOF

git add .
git commit -m "commit"

git push -f -u origin HEAD
