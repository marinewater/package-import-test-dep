package package_import_test_dep

import "github.com/awnumar/memguard/core"

func SetInterval(interval uint64) {
	core.Interval = interval
}

func GetInterval() uint64 {
	return core.Interval
}
