package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/libgit2/git2go"
)

var repo *git.Repository
var commits []string

// Grab all the commits for a given repository.
func repoCommits() []string {
	head, _ := repo.Head()
	walk, _ := repo.Walk()
	oid := head.Target()

	walk.Push(oid)

	oids := []string{}

	walk.Iterate(func(commit *git.Commit) bool {
		oids = append(oids, commit.Id().String())
		return true
	})

	return oids
}

func commitFromOid(sha string) *git.Commit {
	oid, _ := git.NewOid(sha)
	commit, _ := repo.LookupCommit(oid)
	return commit
}

// Search through all the commits for a regex.
func findResult(regex string) (string, string) {
	var longestChunk string
	var longestMatch string

	for i := 0; i < len(commits); i++ {
		// find all string-based parts of the sha
		strings := regexp.MustCompile(regex).FindAllString(commits[i], -1)
		sort.Sort(byLength(strings))

		// prefer the oldest sha to the newest (i.e., the `<=`)
		if len(strings) != 0 && len(longestChunk) <= len(strings[0]) {
			longestMatch = commits[i]
			longestChunk = strings[0]
		}
	}

	return longestMatch, longestChunk
}

func printResult(sha string, chunk string, label string) {
	if sha == "" || chunk == "" {
		fmt.Printf("%s  not found\n", label)
	} else {
		commit := commitFromOid(sha)
		name := commit.Author().Name
		sha = strings.Replace(sha, chunk, fmt.Sprintf("\x1b[33;1m%s\x1b[0m", chunk), -1)

		fmt.Printf("%s  %s\t%s\n", label, sha, name)
	}
}

func main() {
	object, _ := git.OpenRepository(".")

	// I can't just assign this to `repo` above- why? If I try to assign as
	// `repo, _` then other objects can't share the `repo` object.
	repo = object

	commits = repoCommits()

	result, chunk := findResult("[a-f]+")
	printResult(result, chunk, "Longest string  ")

	result, chunk = findResult("[0-9]+")
	printResult(result, chunk, "Longest integer ")

	result, chunk = findResult(`(\\w)(\\1+)`)
	printResult(result, chunk, "Longest repeat  ")
}

type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}
