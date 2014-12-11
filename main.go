package main

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/libgit2/git2go"
)

// Grab all the commits for a given repository.
func repoCommits(repo *git.Repository) []string {
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

// Find the longest continuous non-digit characters in an array of shas.
func findLongestString(commits []string) string {
	var longestChunk string
	var longestCommit string

	for i := 0; i < len(commits); i++ {
		// find all string-based parts of the sha
		strings := regexp.MustCompile("[a-f]+").FindAllString(commits[i], -1)
		sort.Sort(byLength(strings))

		// prefer the oldest sha to the newest (i.e., the `<=`)
		if len(longestChunk) <= len(strings[0]) {
			longestCommit = commits[i]
			longestChunk = strings[0]
		}
	}
	return longestCommit
}

func main() {
	repo, _ := git.OpenRepository(".")
	commits := repoCommits(repo)

	const dateLayout = "January 1, 2006"

	fmt.Print("Longest string: ")
	longest := findLongestString(commits)
	fmt.Println(longest)
	oid, _ := git.NewOid(longest)
	commit, _ := repo.LookupCommit(oid)
	fmt.Println(commit.Author().Name, commit.Author().When.Format(dateLayout))
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
