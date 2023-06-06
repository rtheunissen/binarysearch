load "benchmarks/plot/balancers.gnuplot"

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
