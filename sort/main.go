package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title string
	Artist string
	Album string
	Year int
	Length time.Duration
}


type ByYear []*Track
type ByArtist []*Track


type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

var tracks = []*Track{
	{"Go", "Delilah", "From the roots", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go aheads", "Alicia Keys", "As I am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},

}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}

	return d
}

func printTracks(tracks []*Track)  {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, '\t', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "-----", "-----")

	for _, t := range tracks {
		fmt.Fprintln(tw, t.Title, t.Artist, t.Album, t.Year, t.Length)
		tw.Flush()
	}
}

func (p ByArtist) Len() int {
	return len(p)
}

func (p ByArtist) Less(i, j int) bool {
	return p[i].Artist < p[j].Artist
}

func (p ByArtist) Swap(i, j int)  {
	p[i], p[j] = p[j], p[i]
}
func (p ByYear) Len() int {
	return len(p)
}

func (p ByYear) Less(i, j int) bool {
	return p[i].Year < p[j].Year
}

func (p ByYear) Swap(i, j int)  {
	p[i], p[j] = p[j], p[i]
}


func (p customSort) Len() int {
	return len(p.t)
}

func (p customSort) Less(i, j int) bool {
	return p.less(p.t[i], p.t[j])
}

func (p customSort) Swap(i, j int)  {
	p.t[i], p.t[j] = p.t[j], p.t[i]
}

func main() {
	printTracks(tracks)
	fmt.Printf("\n\n\n\n")

	sort.Sort(ByArtist(tracks))
	printTracks(tracks)
	fmt.Printf("\n\n\n\n")

	sort.Sort(sort.Reverse(ByArtist(tracks)))
	printTracks(tracks)
	fmt.Printf("\n\n\n\n")

	sort.Sort(customSort{
		tracks,
		func(x, y *Track) bool {
			if x.Title != y.Title {
				return x.Title < y.Title
			}

			if x.Year != y.Year {
				return x.Year < y.Year
			}
			if x.Length != y.Length {
				return x.Length < y.Length
			}
			return false
		},
	})
	printTracks(tracks)
	fmt.Printf("\n\n\n\n")
}
