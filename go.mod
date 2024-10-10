// module profilerz

// go 1.23.1

module github.com/exenin/profilerz

go 1.20

require github.com/spf13/cobra v1.8.1

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect

    github.com/exenin/profilerz/util v0.0.0
    github.com/exenin/profilerz/config v0.0.0
    github.com/exenin/profilerz/profile v0.0.0
)


replace github.com/exenin/profilerz/config => ./config
replace github.com/exenin/profilerz/util => ./util
replace github.com/exenin/profilerz/profile => ./profile


