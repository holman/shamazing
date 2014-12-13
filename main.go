package main

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/libgit2/git2go"
)

var repo *git.Repository

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
func findResult(commits []string, regex string) string {
	var longestChunk string
	var longestMatch string

	for i := 0; i < len(commits); i++ {
		// find all string-based parts of the sha
		strings := regexp.MustCompile(regex).FindAllString(commits[i], -1)
		sort.Sort(byLength(strings))

		// prefer the oldest sha to the newest (i.e., the `<=`)
		if len(longestChunk) <= len(strings[0]) {
			longestMatch = commits[i]
			longestChunk = strings[0]
		}
	}

	return longestMatch
}

func printResult(commit *git.Commit, label string) {
	const dateLayout = "January 1, 2006"
	name := commit.Author().Name
	date := commit.Author().When.Format(dateLayout)

	fmt.Printf("%s  %s\t%s\t%s\n", label, commit.Id(), name, date)
}

func main() {
	object, _ := git.OpenRepository(".")

	// I can't just assign this to `repo` above- why? If I try to assign as
	// `repo, _` then other objects can't share the `repo` object.
	repo = object

	commits := repoCommits()

	result := findResult(commits, "[a-f]+")
	printResult(commitFromOid(result), "Longest string           ")
	printResult(commitFromOid(result), "Longest integer          ")
	printResult(commitFromOid(result), "Longest repeating segment")
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
