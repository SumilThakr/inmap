package epi

import (
    "sort"
    "fmt"
    "math"
    "strings"
    "strconv"
    "runtime"
//    "sync"
    "github.com/ctessum/geom"
    "github.com/ctessum/geom/encoding/shp"
    "github.com/ctessum/geom/index/rtree"
    "github.com/ctessum/geom/proj"
)

// inmapOutput holds either population or concentration data read from
// InMAP results
type inmapOutput struct {
    geom.Polygonal

    // Data holds the number of people in each population category, or
    // the TotalPM2.5 concentration in ug/m3.
    Data []float64
}

type mortality struct {
    geom.Polygonal

    // MortData holds the mortality rate for each population category
    MortData []float64 // Deaths per 100,000 people per year

    // Io holds the underlying incidence rate for each population
    // category
    Io []float64
}

func loadPopulation(f string) (*rtree.Rtree, map[string]int, error) {
    projection, _ := proj.Parse("+proj=longlat +units=degrees")
    CensusPopColumns := []string{"TotalPop"}

    var err error
    popshp, err := shp.NewDecoder(f)
    if err != nil {
        return nil, nil, err
    }
    popsr, err := popshp.SR()
    if err != nil {
        return nil, nil, err
    }
