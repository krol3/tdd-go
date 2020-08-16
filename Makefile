unit-tests:
	go test ./adder

integration-tests:
	ginkgo ./adder

tag:
	git tag v0.1.0

tag-push:
	git push origin v0.1.0