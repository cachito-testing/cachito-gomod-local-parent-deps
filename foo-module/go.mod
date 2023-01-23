module github.com/cachito-testing/cachito-gomod-local-parent-deps/foo-module

go 1.18

require github.com/cachito-testing/cachito-gomod-local-parent-deps v0.0.0
require github.com/cachito-testing/cachito-gomod-local-parent-deps/foo-module/bar-module v0.0.0

replace github.com/cachito-testing/cachito-gomod-local-parent-deps => ../
replace github.com/cachito-testing/cachito-gomod-local-parent-deps/foo-module/bar-module => ./foo-package/../bar-module
