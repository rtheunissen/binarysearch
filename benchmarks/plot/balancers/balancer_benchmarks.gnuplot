array Partition[5] = [ "Median", "Height", "Weight", "Cost", "Log" ]

array All[7] = [ "Median", "Height", "Weight", "Cost", "DSW", "Log", "Constant" ]

array Groups[2] = [ "Partition", "All" ]

array Distributions[1] = [ "Uniform" ]

do for [Group=1:|Groups|] {

    GROUP = Groups[Group]

    ##################################################################
    #
    #           DURATION
    #
    ##################################################################

    MEASUREMENT = "Duration"

    set xlabel "Size × 10^5"
    set ylabel ""

    set title "{/:Bold Duration } / Size"

    x = "column('Size')/(column('Scale')/10)"
    y = "column('Duration')/column('Size')"

    set format y2 "%.0f"

    csv_dir = "benchmarks/csv/balancers/benchmarks"

    smooth = "sbezier"
    load "benchmarks/plot/balancers/balancers_lines.gnuplot"

    smooth = "unique"
    load "benchmarks/plot/balancers/balancers_lines.gnuplot"

    smooth = "cumulative"
    load "benchmarks/plot/balancers/balancers_lines.gnuplot"
}
