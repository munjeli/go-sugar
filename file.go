package sugar

import "io/ioutil"

// FileToString takes a path to a file, slurps the file
// into memory and returns it as a string.
func FileToString(p string) (string, error) {
	contents, err := ioutil.ReadFile(p)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}
